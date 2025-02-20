import { PUBLIC_API_URL } from '$env/static/public';
import { redirect } from '@sveltejs/kit';

const PUBLIC_ROUTES = ['/', '/login', '/register'];
const AUTH_ROUTES = ['/app'];

function generateRequestId() {
    return Math.random().toString(36).substring(2, 15);
}

function logEvent(requestId: string, message: string, data?: any) {
    console.log(`[${new Date().toISOString()}] [${requestId}] ${message}`, data || '');
}

function isTokenExpired(token: string): boolean {
    try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        return payload.exp * 1000 < Date.now();
    } catch {
        return true;
    }
}

export const handle = async ({ event, resolve }) => {
    const requestId = generateRequestId();
    const accessToken = event.cookies.get('access_token');
    const refreshToken = event.cookies.get('refresh_token');
    const path = event.url.pathname;

    logEvent(requestId, '=== Request Start ===');
    logEvent(requestId, 'Request path:', path);
    logEvent(requestId, 'Auth status:', {
        hasAccessToken: !!accessToken,
        hasRefreshToken: !!refreshToken
    });

    // First, try to refresh token if we have refresh token but no access token
    if (refreshToken && (!accessToken || isTokenExpired(accessToken))) {
        logEvent(requestId, 'Attempting token refresh');
        try {
            const response = await fetch(`${PUBLIC_API_URL}/auth/refresh`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Accept': 'application/json',
                    'X-Requested-With': 'XMLHttpRequest',
                    'Cookie': `refresh_token=${refreshToken}`
                }
            });

            if (response.ok) {
                const setCookieHeader = response.headers.get('set-cookie');
                if (setCookieHeader) {
                    const cookies = setCookieHeader.split(',').map(cookie => cookie.trim());
                    
                    for (const cookie of cookies) {
                        const [nameValue] = cookie.split(';');
                        const [name, value] = nameValue.split('=');
                        
                        if (name === 'access_token') {
                            event.cookies.set('access_token', value, {
                                path: '/',
                                secure: true,
                                httpOnly: true,
                                sameSite: 'strict',
                                maxAge: 3600
                            });
                        }
                    }
                }
                
                // Update accessToken after refresh
                const newAccessToken = event.cookies.get('access_token');
                if (newAccessToken) {
                    event.locals.authenticated = true;
                }
            }
        } catch (error) {
            logEvent(requestId, 'Token refresh failed:', error);
            // Don't delete tokens here, let the main logic handle it
        }
    }

    // Update authentication status after potential refresh
    event.locals.authenticated = !!event.cookies.get('access_token');

    // Handle routing logic
    if (path === '/' && event.locals.authenticated) {
        logEvent(requestId, 'Authenticated user at root, redirecting to /app');
        throw redirect(302, '/app');
    }

    if (AUTH_ROUTES.includes(path) && !event.locals.authenticated) {
        logEvent(requestId, 'Unauthenticated user accessing protected route');
        event.cookies.delete('access_token', { path: '/' });
        event.cookies.delete('refresh_token', { path: '/' });
        throw redirect(302, '/login');
    }

    const response = await resolve(event);
    logEvent(requestId, '=== Request End ===');
    return response;
};

import { PUBLIC_API_URL } from '$env/static/public';
import { redirect } from '@sveltejs/kit';

const PUBLIC_ROUTES = ['/', '/login', '/register'];
const AUTH_ROUTES = ['/app'];

export const handle = async ({ event, resolve }) => {
    const accessToken = event.cookies.get('access_token');
    const refreshToken = event.cookies.get('refresh_token');
    const path = event.url.pathname;

    // Set authentication status in locals
    event.locals.authenticated = !!accessToken;

    // Handle public routes
    if (PUBLIC_ROUTES.includes(path)) {
        if (accessToken) {
            throw redirect(302, '/app');
        }
    }
    // Handle protected routes
    else if (AUTH_ROUTES.includes(path)) {
        if (!accessToken) {
            if (refreshToken) {
                // Try to refresh the token
                try {
                    const response = await fetch(`${PUBLIC_API_URL}/auth/refresh`, {
                        method: 'POST',
                        headers: {
                            Cookie: `refresh_token=${refreshToken}`
                        }
                    });
                    
                    if (!response.ok) {
                        throw redirect(302, '/login');
                    }
                } catch {
                    throw redirect(302, '/login');
                }
            } else {
                throw redirect(302, '/login');
            }
        }
    }

    return await resolve(event);
};

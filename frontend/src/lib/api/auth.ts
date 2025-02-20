// api/auth.ts
import { PUBLIC_API_URL } from '$env/static/public';
import type { LoginRequest, LoginResponse } from '$lib/types/auth';

export async function register(email: string, password: string) {
    console.log(PUBLIC_API_URL)
    const response = await fetch(`${PUBLIC_API_URL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.message || 'Registration failed');
    }

    return response.json();
}

export async function login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await fetch(`${PUBLIC_API_URL}/auth/login`, {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials)
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Login failed');
    }

    return response.json();
}

export async function refresh(): Promise<void> {
    const response = await fetch(`${PUBLIC_API_URL}/auth/refresh`, {
        method: 'POST',
        credentials: 'include'
    });

    if (!response.ok) {
        throw new Error('Failed to refresh token');
    }
}

export async function logout(): Promise<void> {
    // Clear cookies by setting expiry to past
    const response = await fetch(`/logout`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        credentials: 'include'
    });

    if (!response.ok) {
        throw new Error('Failed to logout');
    }
    
    // Clear cookies by setting expiry to past
    document.cookie = 'access_token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    document.cookie = 'refresh_token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    
    // Redirect to login page
    window.location.href = '/login';
    return;
}
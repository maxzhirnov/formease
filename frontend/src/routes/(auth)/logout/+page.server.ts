import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
    default: async ({ cookies }) => {
        cookies.delete('access_token', { path: '/' });
        cookies.delete('refresh_token', { path: '/refresh' });
        throw redirect(302, '/login');
    }
};

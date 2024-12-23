import type { PageLoad } from './$types';
import type { FormData, ThemeName, FloatingShapesTheme } from "$lib/types";
import { PUBLIC_API_URL } from '$env/static/public';

export const load: PageLoad = async ({ params, fetch }) => {
    try {
        // Use the provided fetch instance
        const response = await fetch(`${PUBLIC_API_URL}/forms/${params.id}`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch form');
        }

        const form: FormData = await response.json();
        
        return {
            form: {
                questions: form.questions,
                thankYouMessage: form.thankYouMessage,
                theme: form.theme as ThemeName,
                floatingShapesTheme: form.floatingShapesTheme as FloatingShapesTheme
            }
        };
    } catch (error) {
        console.error('Error loading form:', error);
        return {
            form: null,
            error: 'Failed to load form'
        };
    }
};

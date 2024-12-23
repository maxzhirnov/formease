import type { PageLoad } from './$types';
import type { ThemeName, FloatingShapesTheme } from "$lib/types";
import { formService } from '$lib/services/formService';

export const load: PageLoad = async ({ params, fetch: customFetch }) => {
    try {
        const form = await formService.fetchPublic(params.id, customFetch);
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

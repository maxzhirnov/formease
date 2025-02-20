import type { PageLoad } from './$types';
import { formService } from '$lib/services/formService';
import type { ThemeName } from '$lib/types/theme';
import type { FloatingShapesTheme } from '$lib/types/shapes';

export const load: PageLoad = async ({ params, fetch: customFetch }) => {
    try {
        const form = await formService.api.fetchPublic(params.id, customFetch);
        return {
            form: {
                id: form.id,
                questions: form.questions,
                thankYouMessage: form.thankYouMessage,
                theme: form.theme as ThemeName,
                floatingShapesTheme: form.floatingShapesTheme as FloatingShapesTheme,
                isDraft: form.isDraft
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

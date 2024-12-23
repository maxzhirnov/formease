import type { FormData } from '$lib/types';
import { PUBLIC_API_URL } from '$env/static/public';

const fetchWithCreds = (url: string, options: RequestInit = {}, customFetch: typeof fetch = fetch) => {
    return customFetch(url, {
        ...options,
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
    });
};

export async function createForm(formData: FormData): Promise<FormData> {
    const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms`, {
        method: 'POST',
        body: JSON.stringify(formData)
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to create form');
    }

    return await response.json();
}

export async function listForms() {
    console.log('[listForms] Starting to fetch forms...');

    try {
        const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms`);
        console.log(`[listForms] Fetch request completed with status: ${response.status}`);

        if (!response.ok) {
            const error = await response.json();
            console.error('[listForms] Error fetching forms:', error);
            throw new Error(error.error || 'Failed to fetch forms');
        }

        const forms = await response.json();
        console.log(`[listForms] Received forms: ${forms}`);
        return forms;
    } catch (error) {
        if (error instanceof Error) {
            console.error('[listForms] An error occurred:', error.message);
            throw error;
        } else {
            console.error('[listForms] An unexpected error occurred:', error);
            throw new Error('An unknown error occurred');
        }
    }
}

export async function getForm(id: string, {}, customFetch: typeof fetch = fetch): Promise<FormData> {
    const response = await fetchWithCreds(
        `${PUBLIC_API_URL}/my-forms/${id}`, 
        {},
        customFetch
    );
    
    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to fetch form');
    }

    return await response.json();
}

export async function getFormPublic(id: string, customFetch: typeof fetch = fetch): Promise<FormData> {
    const response = await customFetch(`${PUBLIC_API_URL}/forms/${id}`);
    
    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to fetch form');
    }

    return await response.json();
}

export async function updateForm(id: string, formData: FormData) {
    console.log('Sending form data to server:', formData);
    const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms/${id}`, {
        method: 'PUT',
        body: JSON.stringify(formData)
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to update form');
    }

    return response.json();
}

export async function deleteForm(id: string) {
    const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms/${id}`, {
        method: 'DELETE'
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to delete form');
    }

    return response.json();
}

export async function generateFormWithAI(description: string): Promise<FormData> {
    try {
        const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms/generate-ai`, {
            method: 'POST',
            body: JSON.stringify({ description })
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || 'Failed to generate form');
        }

        return await response.json();
    } catch (error) {
        console.error('AI Form Generation Error:', error);
        throw error;
    }
}

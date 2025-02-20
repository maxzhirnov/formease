import type { FormData } from '$lib/types/form';
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
    try {
        const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms`);

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Failed to fetch forms');
        }

        const forms = await response.json();
        return forms;
    } catch (error) {
        if (error instanceof Error) {
            throw error;
        } else {
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

export interface AIFormRequest {
    topic: string;
    formType: string;
    numQuestions: number;
    preferences?: string[];
}

export async function generateFormWithAI(aiRequest: AIFormRequest): Promise<FormData> {
        const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms/generate-form`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(aiRequest)
        });
    
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Failed to generate AI form');
        }
    
        return await response.json();
}

export async function toggleDraft(id: string, formData: FormData) {
    const response = await fetchWithCreds(`${PUBLIC_API_URL}/my-forms/${id}/toggle-draft`, {
        method: 'GET',
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to update form');
    }

    return response.json();
}

export async function uploadImage(file: File): Promise<string> {
    const formData = new FormData();
        formData.append('image', file);
        
        const response = await fetch(`${PUBLIC_API_URL}/image-upload`, {
            method: 'POST',
            body: formData,
            credentials: 'include'
        });

        
        if (!response.ok) {
            throw new Error('Upload failed');
        }

        const  resp  = await response.json();

        return resp.url;
}

export interface ImageResponse {
    images: ImageData[];
    total: number;
    page: number;
    limit: number;
}

export async function getImages(page: number, limit: number): Promise<ImageResponse> {
    const response = await fetch(`${PUBLIC_API_URL}/images?page=${page}&limit=${limit}`, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        }
    });
    
    if (!response.ok) {
        throw new Error('Failed to fetch images');
    }

    return await response.json();
}

export async function deleteImage(imageId: string): Promise<void> {
    const response = await fetch(`${PUBLIC_API_URL}/images/${imageId}`, {
        method: 'DELETE',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        }
    });

    if (!response.ok) {
        throw new Error('Failed to delete image');
    }
}
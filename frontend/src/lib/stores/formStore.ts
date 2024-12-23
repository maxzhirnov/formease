// src/lib/stores/formStore.ts

import { writable } from 'svelte/store';
import type { FormData, Question } from '$lib/types';
import { QuestionType } from '$lib/types';
import { get } from 'svelte/store';
import { createForm } from '$lib/api/forms';
import { createQuestion } from '$lib/factories/QuestionFactory';


// Initial form data structure with default values
const initialFormData: FormData = {
  id: '',
  isDraft: true,
  name: '',
  userId: '',
  theme: 'light',
  floatingShapesTheme: 'default',
  questions: [],
  thankYouMessage: { title: '', subtitle: '', icon: '', button: { text: '', url: '', newTab: false } }
};

// Create a writable store with initial form data
export const formStore = writable<FormData>(initialFormData);

// Initialize a counter for question IDs
let currentQuestionId = 0; 

/**
 * Resets the form data to its initial state.
 * This function can be called to clear all fields and questions.
 */
export const resetForm = () => {
    formStore.set(initialFormData);
};

/**
 * Updates specific properties of the form data.
 * @param newData - An object containing the properties to update.
 * Example usage:
 * updateForm({ name: 'New Form Name' });
 */
export const updateForm = (newData: Partial<FormData>) => {
    formStore.update(currentData => ({ ...currentData, ...newData }));
};

/**
 * Adds a new question to the form.
 * @param type - The type of question to add (e.g., QuestionType.SINGLE_CHOICE).
 * This function increments the question ID counter and creates a new question object
 * based on the specified type.
 */
export const addQuestion = (type: QuestionType) => {
    formStore.update(formData => {

        const id = formData.questions?.length ?? 0 > 0 
            ? Math.max(...formData.questions?.map(q => q.id) ?? [0]) + 1 
            : 1;

        const newQuestion = createQuestion(id,type);
        const currentQuestions = formData.questions || [];
        
        return {
            ...formData,
            questions: [...currentQuestions, newQuestion]
        };
    });
};

/**
 * Removes a question from the form by its ID.
 * @param id - The ID of the question to remove.
 * Example usage:
 * removeQuestion(1); // Removes the question with ID 1.
 */
export const removeQuestion = (id: number) => {
    formStore.update(formData => {
        if (!formData.questions) {
            return formData;
        }
        
        return {
            ...formData,
            questions: formData.questions.filter(q => q.id !== id)
        };
    });
};

/**
 * Updates a specific question in the form by its ID.
 * @param id - The ID of the question to update.
 * @param updatedQuestion - The updated question object that replaces the existing one.
 * Example usage:
 * updateQuestion(1, updatedQuestion); // Updates the question with ID 1.
 */
export const updateQuestion = (updatedQuestion: Question) => {
    formStore.update(currentData => {
        if (!currentData.questions) {
            return {
                ...currentData,
                questions: [updatedQuestion]
            };
        }

        return {
            ...currentData,
            questions: currentData.questions.map(q => 
                q.id === updatedQuestion.id ? updatedQuestion : q
            )
        };
    });
};
/**
 * Saves the current form data to the server.
 * This function uses the createForm API function to send data.
 */
export const saveFormToServer = async () => {
    const currentFormData = get(formStore); // Get current form data from the store
    try {
        const savedForm = await createForm(currentFormData); // Call the createForm API function
        console.log('Form saved successfully:', savedForm);
        resetForm(); // Optionally reset the store after saving
    } catch (error) {
        console.error('Error saving form:', error);
        // Handle error feedback for user (e.g., show a message)
    }
};

/**
 * Updates the form name in the store.
 * @param name - The new name for the form.
 * Example usage:
 * updateFormName('My New Form');
 */
export const updateFormName = (name: string) => {
    formStore.update(currentData => ({
        ...currentData,
        name
    }));
};

/**
 * Updates an existing form on the server
 */
export const updateFormOnServer = async (id: string) => {
    const currentFormData = get(formStore);
    try {
        const updatedForm = await updateForm(currentFormData);
        console.log('Form updated successfully:', updatedForm);
        return updatedForm;
    } catch (error) {
        console.error('Error updating form:', error);
        throw error;
    }
};
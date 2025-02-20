import { createForm, updateForm, getForm, getFormPublic, deleteForm, listForms, generateFormWithAI, toggleDraft, type AIFormRequest } from '$lib/api/forms';
import {type FormData} from '$lib/types/form';
import { FormStateService, formStateService } from './formStateService';

export class FormApiService {
    constructor(private stateService: FormStateService) {}
    
    // Server Interactions
  async create(): Promise<FormData> {
    const currentFormData = this.stateService.getCurrentForm();
    const formData: FormData = {
      id: currentFormData.id || '',
      isDraft: currentFormData.isDraft ?? true,
      name: currentFormData.name || 'Untitled Form',
      theme: currentFormData.theme || 'light',
      userId: currentFormData.userId || '', // Ensure this is set
      floatingShapesTheme: currentFormData.floatingShapesTheme || 'default',
      questions: currentFormData.questions || [],
      thankYouMessage: currentFormData.thankYouMessage || {
        title: '',
        subtitle: '',
        icon: '',
        button: { text: '', url: '', newTab: false }
      }
    };
    
    return createForm(formData);
  }

  async generateAIForm(description: AIFormRequest): Promise<FormData> {
      try {
          // Reset the form before generating a new one
          this.stateService.resetForm();

          // Call API to generate form
          const generatedForm = await generateFormWithAI(description);

          // Update the store with the generated form
          this.stateService.setForm ({
            ...generatedForm,
            isDraft: true, // Ensure it's a draft
            id: '', // Clear the ID to force a new creation
          })

          // Create the form in the backend
          const createdForm = await this.create();

          return createdForm;
      } catch (error) {
          console.error('AI Form Generation Failed:', error);
          throw error;
      }
  }
  

  async update(id: string): Promise<FormData> {
    const currentFormData = this.stateService.getCurrentForm();
    return updateForm(id, currentFormData);
  }

  async toggleDraft(id: string): Promise<void> {
    await toggleDraft(id, this.stateService.getCurrentForm());
  }

  async fetch(id: string, customFetch: typeof fetch = fetch): Promise<FormData> {
    const form = await getForm(id, {}, customFetch);
    this.stateService.setForm(form)
    return form;
}



  async fetchPublic(id: string, customFetch: typeof fetch = fetch): Promise<FormData> {
    this.stateService.setForm(await getFormPublic(id, customFetch));
    return this.stateService.getCurrentForm();
  }

  async delete(id: string): Promise<void> {
    await deleteForm(id);
  }

  async list(): Promise<FormData[]> {
    return listForms();
  }
}

export const formApiService = new FormApiService(formStateService);
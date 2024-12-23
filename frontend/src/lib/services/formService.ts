import { get } from 'svelte/store';
import { formStore } from '$lib/stores/formStore';
import { createForm, updateForm, getForm, getFormPublic, deleteForm, listForms, generateFormWithAI } from '$lib/api/forms';
import  { QuestionType, type BaseQuestion, type FormData, type Question } from '$lib/types';
import { createQuestion } from '$lib/factories/QuestionFactory';

class FormService {
  // Store access
  private store = formStore;

  // Getters
  getCurrentForm(): FormData {
    return get(this.store);
  }

  // Form Management
  resetForm() {
    this.store.set({
      id: '',
      isDraft: true,
      name: '',
      userId: '',
      theme: 'light',
      floatingShapesTheme: 'default',
      questions: [],
      thankYouMessage: { 
        title: '', 
        subtitle: '', 
        icon: '', 
        button: { text: '', url: '', newTab: false } 
      }
    });
  }

  setUserId(userId: string) {
    this.store.update(formData => ({
      ...formData,
      userId
    }));
  }

  // Server Interactions
  async create(): Promise<FormData> {
    const currentFormData = this.getCurrentForm();
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

  async generateAIForm(description: string): Promise<FormData> {
      try {
          // Reset the form before generating a new one
          this.resetForm();

          // Call API to generate form
          const generatedForm = await generateFormWithAI(description);

          // Update the store with the generated form
          this.store.set({
              ...generatedForm,
              isDraft: true, // Ensure it's a draft
              id: '', // Clear the ID to force a new creation
          });

          // Create the form in the backend
          const createdForm = await this.create();

          return createdForm;
      } catch (error) {
          console.error('AI Form Generation Failed:', error);
          throw error;
      }
  }
  

  async update(id: string): Promise<FormData> {
    const currentFormData = this.getCurrentForm();
    return updateForm(id, currentFormData);
  }

  async fetch(id: string, customFetch: typeof fetch = fetch): Promise<FormData> {
    const form = await getForm(id, customFetch);
    this.store.set(form);
    return form;
  }

  async fetchPublic(id: string, customFetch: typeof fetch = fetch): Promise<FormData> {
    return getFormPublic(id, customFetch);
  }

  async delete(id: string): Promise<void> {
    await deleteForm(id);
  }

  async list(): Promise<FormData[]> {
    return listForms();
  }

  // Utility Methods
  updateFormName(name: string) {
    this.store.update(formData => ({ ...formData, name }));
  }

  updateTheme(theme: string) {
    this.store.update(formData => ({ ...formData, theme }));
  }

  // Advanced Form Logic
  findQuestionById(id: number): Question | undefined {
    return this.getCurrentForm().questions?.find(q => q.id === id);
  }

  getNextQuestionId(currentQuestion: Question): number | undefined {
    if (!currentQuestion.nextQuestion) return undefined;

    // First, check condition-based routing
    const conditionMatch = currentQuestion.nextQuestion.conditions.find(
      condition => condition.answer // You might want to pass the actual answer here
    );

    // If condition match exists, return its nextId
    if (conditionMatch) return conditionMatch.nextId;

    // Otherwise, return default
    return currentQuestion.nextQuestion.default;
  }

  // Question Management Methods
  // Method to create a question with correct type
  addQuestion(type: QuestionType): void {
    this.store.update(form => {
      const id = this.generateQuestionId();
      const newQuestion = createQuestion(id, type);
      return {
        ...form,
        questions: [...(form.questions || []), newQuestion]
      };
    });
  }

  // Typed method for adding options
  addQuestionOption(questionId: number, optionText: string, optionIcon?: string): void {
    this.store.update(form => ({
      ...form,
      questions: (form.questions || []).map(question => {
        if (
          question.id === questionId && 
          (question.type === QuestionType.SINGLE_CHOICE || 
           question.type === QuestionType.MULTIPLE_CHOICE)
        ) {
          return {
            ...question,
            options: [
              ...question.options, 
              { 
                text: optionText, 
                ...(optionIcon ? { icon: optionIcon } : {}) 
              }
            ]
          };
        }
        return question;
      })
    }));
  }

  // Typed method for removing options
  removeQuestionOption(questionId: number, optionIndex: number): void {
    this.store.update(form => ({
      ...form,
      questions: (form.questions || []).map(question => {
        if (
          question.id === questionId && 
          (question.type === QuestionType.SINGLE_CHOICE || 
           question.type === QuestionType.MULTIPLE_CHOICE)
        ) {
          return {
            ...question,
            options: question.options.filter((_, index) => index !== optionIndex)
          };
        }
        return question;
      })
    }));
  }

  // Typed method for updating questions
  updateQuestion(questionId: number, updates: Partial<Question>): void {
    this.store.update(form => ({
      ...form,
      questions: (form.questions || []).map(question => {
        // Type guard to ensure we're updating the correct question type
        if (question.id === questionId) {
          // Ensure type consistency
          if (
            updates.type && 
            updates.type !== question.type
          ) {
            throw new Error('Cannot change question type');
          }

          return {
            ...question,
            ...updates
          } as Question;
        }
        return question;
      })
    }));
  }

  // Typed method for getting a question
  getQuestionById(questionId: number): Question | undefined {
    const currentForm = get(this.store);
    return (currentForm.questions || []).find(q => q.id === questionId);
  }

  // Typed method for updating next question routing
  updateNextQuestionRouting(
    questionId: number, 
    routing: BaseQuestion['nextQuestion']
  ): void {
    this.store.update(form => ({
      ...form,
      questions: (form.questions || []).map(question => 
        question.id === questionId 
          ? { ...question, nextQuestion: routing } 
          : question
      )
    }));
  }

  // Additional type-specific methods could be added
  setInputQuestionType(
    questionId: number, 
    inputType: 'text' | 'email' | 'number'
  ): void {
    this.store.update(form => ({
      ...form,
      questions: (form.questions || []).map(question => {
        if (
          question.id === questionId && 
          question.type === QuestionType.INPUT
        ) {
          return {
            ...question,
            inputType
          };
        }
        return question;
      })
    }));
  }

  // Method to set max selections for multiple choice
  setMaxSelections(questionId: number, maxSelections: number): void {
    this.store.update(form => ({
      ...form,
      questions: (form.questions || []).map(question => {
        if (
          question.id === questionId && 
          question.type === QuestionType.MULTIPLE_CHOICE
        ) {
          return {
            ...question,
            maxSelections
          };
        }
        return question;
      })
    }));
  }

  private generateQuestionId(): number {
    const currentForm = get(this.store);
    const questions = currentForm.questions || [];
    return questions.length > 0 
      ? Math.max(...questions.map(q => q.id)) + 1 
      : 1;
  }
}

// Singleton export
export const formService = new FormService();

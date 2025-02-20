import { formState } from '$lib/state/formState.svelte';
import  { type FormData } from '$lib/types/form';
import { createQuestion } from '$lib/factories/QuestionFactory';
import { QuestionType, type BaseQuestion, type MultipleChoiceQuestion, type Question, type QuestionOption, type SingleChoiceQuestion } from '$lib/types/questions';
import type { ThemeName } from '$lib/types/theme';
import type { FloatingShapesTheme } from '$lib/types/shapes';

export class FormStateService {
  constructor(private store = formState) {}
  

  getCurrentForm(): FormData {
    return this.store.state;
  }
  

  setForm(newForm: FormData) {
    this.store.update(newForm);
  }

  updateFormName(name: string) {
    this.store.update({ name });
  }

  resetForm() {
    this.store.reset();
  }

  // utilities
  hasQuestions(): boolean {
    try {
      const questions = this.store.state.questions;
      if (!questions) return false;
      return questions.length > 0;
    } catch {
      return false;
    }
  }

  private isChoiceQuestion(question: Question): question is SingleChoiceQuestion | MultipleChoiceQuestion {
    return (
      question.type === QuestionType.SINGLE_CHOICE || 
      question.type === QuestionType.MULTIPLE_CHOICE
    );
  }


  setUserId(userId: string): void {
    if (!userId.trim()) {
      throw new Error('User ID cannot be empty');
    }
    this.store.update({ userId });
  }

  updateTheme(themeName: ThemeName): void {
    this.store.update({ theme: themeName });
  }

  updateFloatingShapesTheme(floatingShapesTheme: FloatingShapesTheme): void {
    this.store.update({ floatingShapesTheme });
  }

  findQuestionById(id: number): Question | undefined {
    try {
      const questions = this.store.state.questions;
      if (!questions) return undefined;
      return questions.find(question => question.id === id);
    } catch (error) {
      console.error(`Error finding question with id ${id}:`, error);
      return undefined;
    }
  }

  getNextQuestionId(currentQuestion: Question): number | undefined {
    if (!currentQuestion?.nextQuestion) {
      return undefined;
    }

    const { conditions, default: defaultNext } = currentQuestion.nextQuestion;

    // Check for condition-based routing
    const matchingCondition = conditions?.find(condition => condition.answer);
    
    return matchingCondition?.nextId ?? defaultNext;
  }


  addQuestion(type: QuestionType): void {
    const newQuestion = createQuestion(this.generateQuestionId(), type);
    const questions = [...(this.store.state.questions || []), newQuestion];
    
    this.store.update({ questions });
  }

 /**
   * Adds a new option to a single or multiple choice question
   * @param questionId - The ID of the question to add the option to
   * @param optionText - The text for the new option
   * @param optionIcon - Optional icon for the new option
   * @throws {Error} If question not found or type is invalid
   */
 addQuestionOption(questionId: number, optionText: string, optionIcon?: string): void {
    try {
      // Get current questions state
      const questions = this.store.state.questions;
      
      // Validate questions array exists
      if (!Array.isArray(questions)) {
        throw new Error('Questions array is not initialized');
      }

      // Find and validate question
      const question = questions.find(q => q.id === questionId);
      if (!question) {
        throw new Error(`Question with ID ${questionId} not found`);
      }

      // Type guard for choice questions
      if (!this.isChoiceQuestion(question)) {
        throw new Error(`Cannot add options to question type: ${question.type}`);
      }

      // Create new option
      const newOption = {
        id: this.generateOptionId(question),
        text: optionText.trim(),
        ...(optionIcon ? { icon: optionIcon } : {}),
        color: ''
      };

      // Update questions array with new option
      const updatedQuestions = questions.map(q => {
        if (q.id === questionId && this.isChoiceQuestion(q)) {
          return {
            ...q,
            options: [...(q.options || []), newOption]
          };
        }
        return q;
      });

      // Update store
      this.store.update({ questions: updatedQuestions });

    } catch (error) {
      console.error('Failed to add question option:', error);
      throw error;
    }
  }

  private generateOptionId(question: SingleChoiceQuestion | MultipleChoiceQuestion): number {
      const existingOptions = question.options || [];
      return existingOptions.length > 0 
          ? Math.max(...existingOptions.map(o => o.id)) + 1 
          : 1;
  }

  /**
 * Updates a specific option within a question
 * @param questionId - The ID of the question containing the option
 * @param optionId - The ID of the option to update
 * @param updates - Partial option object containing the updates
 * @throws {Error} If question or option not found
 */
updateQuestionOption(questionId: number, optionId: number, updates: Partial<QuestionOption>): void {
    try {
        // Get and validate questions array
        const questions = this.store.state.questions;
        if (!Array.isArray(questions)) {
            throw new Error('Questions array is not initialized');
        }

        // Find and validate question
        const question = questions.find(q => q.id === questionId);
        if (!question) {
            throw new Error(`Question with ID ${questionId} not found`);
        }

        // Type guard for choice questions
        if (!this.isChoiceQuestion(question)) {
            throw new Error(`Cannot update options for question type: ${question.type}`);
        }

        // Find the option to update
        const optionToUpdate = question.options.find(opt => opt.id === optionId);
        if (!optionToUpdate) {
            throw new Error(`Option with ID ${optionId} not found in question ${questionId}`);
        }

        // Update questions array
        const updatedQuestions = questions.map(q => {
            if (q.id === questionId && this.isChoiceQuestion(q)) {
                return {
                    ...q,
                    options: q.options.map(opt => 
                        opt.id === optionId 
                            ? { ...opt, ...updates }
                            : opt
                    )
                };
            }
            return q;
        });

        // Update store
        this.store.update({ questions: updatedQuestions });

    } catch (error) {
        console.error('Failed to update question option:', error);
        throw error;
    }
  }

 /**
   * Removes an option from a single or multiple choice question
   * @param questionId - The ID of the question to remove the option from
   * @param optionIndex - The index of the option to remove
   * @throws {Error} If validation fails or option cannot be removed
   */
 removeQuestionOption(questionId: number, optionIndex: number): void {
    try {
      // Get and validate questions array
      const questions = this.store.state.questions;
      if (!Array.isArray(questions)) {
        throw new Error('Questions array is not initialized');
      }

      // Find and validate question
      const question = questions.find(q => q.id === questionId);
      if (!question) {
        throw new Error(`Question with ID ${questionId} not found`);
      }

      // Validate question type
      if (!this.isChoiceQuestion(question)) {
        throw new Error(`Cannot remove options from question type: ${question.type}`);
      }

      // Validate option index
      if (!question.options || optionIndex < 0 || optionIndex >= question.options.length) {
        throw new Error(`Invalid option index: ${optionIndex}`);
      }

      // Validate minimum options (if needed)
      // if (question.options.length <= 1) {
      //   throw new Error('Cannot remove the last option from a choice question');
      // }

      // Update questions array
      const updatedQuestions = questions.map(q => {
        if (q.id === questionId && this.isChoiceQuestion(q)) {
          return {
            ...q,
            options: [
              ...q.options.slice(0, optionIndex),
              ...q.options.slice(optionIndex + 1)
            ]
          };
        }
        return q;
      });

      // Update store
      this.store.update({ questions: updatedQuestions });

    } catch (error) {
      console.error('Failed to remove question option:', error);
      throw error;
    }
  }


  /**
   * Updates a specific question with provided updates
   * @param questionId - The ID of the question to update
   * @param updates - Partial question object containing the updates
   * @throws {Error} If question not found or update is invalid
   */
  updateQuestion(questionId: number, updates: Partial<Question>): void {
      try {
        // Get and validate questions array
        const questions = this.store.state.questions;
        if (!Array.isArray(questions)) {
          throw new Error('Questions array is not initialized');
        }

        // Find the question to update
        const questionToUpdate = questions.find(q => q.id === questionId);
        if (!questionToUpdate) {
          throw new Error(`Question with ID ${questionId} not found`);
        }

        // Validate updates
        this.validateQuestionUpdates(questionToUpdate, updates);

        // Create updated questions array
        const updatedQuestions = questions.map(question => {
          if (question.id === questionId) {
            return {
              ...question,
              ...updates
            } as Question;
          }
          return question;
        });

        // Update store
        this.store.update({ questions: updatedQuestions });

      } catch (error) {
        console.error('Failed to update question:', error);
        throw error;
      }
  }

  private validateQuestionUpdates(
    currentQuestion: Question, 
    updates: Partial<Question>
  ): void {
    // Prevent type changes
    if (updates.type && updates.type !== currentQuestion.type) {
      throw new Error('Cannot change question type');
    }

    // Add any additional validation rules here
    this.validateQuestionFields(updates);
  }
  
  private validateQuestionFields(updates: Partial<Question>): void {
    // Add field-specific validation
    if (updates.question && updates.question.trim().length === 0) {
      throw new Error('Question title cannot be empty');
    }

    // Add more field validations as needed
  }



  /**
   * Retrieves a question by its ID
   * @param questionId - The ID of the question to retrieve
   * @returns The found question or undefined
   */
  getQuestionById(questionId: number): Question | undefined {
    return this.store.state.questions?.find(q => q.id === questionId);
  }

  // Typed method for updating next question routing
  /**
   * Updates the routing configuration for a specific question
   * @param questionId - The ID of the question to update
   * @param routing - The new routing configuration
   * @throws {Error} If question not found or routing is invalid
   */
  updateNextQuestionRouting(
      questionId: number,
      routing: BaseQuestion['nextQuestion']
    ): void {
      try {
        // Get and validate questions array
        const questions = this.store.state.questions;
        if (!Array.isArray(questions)) {
          throw new Error('Questions array is not initialized');
        }

        // Validate question exists
        const questionToUpdate = questions.find(q => q.id === questionId);
        if (!questionToUpdate) {
          throw new Error(`Question with ID ${questionId} not found`);
        }

        // Update questions array
        const updatedQuestions = questions.map(question =>
          question.id === questionId
            ? { ...question, nextQuestion: routing }
            : question
        );

        // Update store
        this.store.update({ questions: updatedQuestions });

      } catch (error) {
        console.error('Failed to update question routing:', error);
        throw error;
      }
    }


  // Additional type-specific methods could be added
  /**
   * Sets the input type for an input question
   * @param questionId - The ID of the question to update
   * @param inputType - The new input type
   * @throws {Error} If question not found or is not an input question
   */
  setInputQuestionType(
      questionId: number,
      inputType: 'text' | 'email' | 'number'
    ): void {
      try {
        // Get and validate questions array
        const questions = this.store.state.questions;
        if (!Array.isArray(questions)) {
          throw new Error('Questions array is not initialized');
        }

        // Find and validate question
        const question = questions.find(q => q.id === questionId);
        if (!question) {
          throw new Error(`Question with ID ${questionId} not found`);
        }

        // Validate question type
        if (question.type !== QuestionType.INPUT) {
          throw new Error(`Question ${questionId} is not an input question`);
        }

        // Update questions array
        const updatedQuestions = questions.map(q => {
          if (q.id === questionId && q.type === QuestionType.INPUT) {
            return {
              ...q,
              inputType
            };
          }
          return q;
        });

        // Update store
        this.store.update({ questions: updatedQuestions });

      } catch (error) {
        console.error('Failed to set input question type:', error);
        throw error;
      }
    }


  /**
   * Sets the maximum number of selections for a multiple choice question
   * @param questionId - The ID of the question to update
   * @param maxSelections - The maximum number of selections allowed
   * @throws {Error} If validation fails
   */
  setMaxSelections(questionId: number, maxSelections: number): void {
      try {
        // Get and validate questions array
        const questions = this.store.state.questions;
        if (!Array.isArray(questions)) {
          throw new Error('Questions array is not initialized');
        }

        // Find and validate question
        const question = questions.find(q => q.id === questionId);
        if (!question) {
          throw new Error(`Question with ID ${questionId} not found`);
        }

        // Validate question type
        if (question.type !== QuestionType.MULTIPLE_CHOICE) {
          throw new Error(`Question ${questionId} is not a multiple choice question`);
        }

        // Update questions array
        const updatedQuestions = questions.map(q => {
          if (q.id === questionId && q.type === QuestionType.MULTIPLE_CHOICE) {
            return {
              ...q,
              maxSelections
            };
          }
          return q;
        });

        // Update store
        this.store.update({ questions: updatedQuestions });

      } catch (error) {
        console.error('Failed to set max selections:', error);
        throw error;
      }
    }


  /**
   * Generates a unique ID for a new question
   * @private
   * @returns A unique question ID
   */
  private generateQuestionId(): number {
    try {
      const questions = this.store.state.questions || [];
      return questions.length > 0 
        ? Math.max(...questions.map(q => q.id)) + 1 
        : 1;
    } catch (error) {
      console.error('Error generating question ID:', error);
      return Date.now(); // Fallback to timestamp if error occurs
    }
  }

  /**
   * Removes a question and updates related question IDs and routing
   * @param id - The ID of the question to remove
   * @throws {Error} If validation fails
   */
  removeQuestion(id: number): void {
    try {
      // Get and validate questions array
      const questions = this.store.state.questions;
      if (!Array.isArray(questions)) {
        throw new Error('Questions array is not initialized');
      }

      // Validate question exists
      if (!questions.find(q => q.id === id)) {
        throw new Error(`Question with ID ${id} not found`);
      }

      // Create backup for rollback
      const questionsBackup = [...questions];

      try {
        // Remove question and update remaining questions
        const updatedQuestions = this.updateQuestionsAfterRemoval(questions, id);

        // Update store
        this.store.update({ questions: updatedQuestions });

        // Emit question removed event
        this.emitQuestionRemoved(id);

      } catch (error) {
        // Rollback on failure
        this.rollbackQuestions(questionsBackup);
        throw error;
      }

    } catch (error) {
      console.error('Failed to remove question:', error);
      throw error;
    }
  }

  private updateQuestionsAfterRemoval(questions: Question[], removedId: number): Question[] {
    const filteredQuestions = questions.filter(q => q.id !== removedId);

    return filteredQuestions.map((question, index) => {
      const newId = index + 1;
      
      return {
        ...question,
        id: newId,
        nextQuestion: this.updateNextQuestion(question.nextQuestion, removedId, newId)
      };
    });
  }

  private updateNextQuestion(
    nextQuestion: Question['nextQuestion'],
    removedId: number,
    currentNewId: number
  ): Question['nextQuestion'] {
    if (!nextQuestion) return undefined;

    return {
      ...nextQuestion,
      default: this.updateRoutingId(nextQuestion.default, removedId),
      conditions: nextQuestion.conditions?.map(condition => ({
        ...condition,
        nextId: this.updateRoutingId(condition.nextId, removedId) ?? condition.nextId // Fallback to original nextId
      }))
    };
  }

  private updateRoutingId(routingId: number | undefined, removedId: number): number {
    if (typeof routingId === 'undefined') {
      return 1; // Or some default value that makes sense in your context
    }
    return routingId > removedId ? routingId - 1 : routingId;
  }


  /**
   * Rolls back questions to previous state
   * @private
   */
  private rollbackQuestions(backup: Question[]): void {
    try {
      this.store.update({ questions: backup });
    } catch (error) {
      console.error('Failed to rollback questions:', error);
    }
  }

  /**
   * Emits question removed event
   * @private
   */
  private emitQuestionRemoved(id: number): void {
    window.dispatchEvent(new CustomEvent('question-removed', {
      detail: { questionId: id }
    }));
  }
}



// Singleton export
export const formStateService = new FormStateService();

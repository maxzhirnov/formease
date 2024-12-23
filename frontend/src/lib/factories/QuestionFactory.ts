// src/lib/factories/QuestionFactory.ts

import type { SingleChoiceQuestion, MultipleChoiceQuestion, InputQuestion, Question } from '$lib/types';
import { QuestionType } from '$lib/types';

// Function to create a new question based on type
export function createQuestion(id: number, type: QuestionType): Question {
    const newQuestionId = id

    switch (type) {
        case QuestionType.SINGLE_CHOICE:
            return {
                id: newQuestionId,
                question: '',
                subtext: '',
                image: '',
                type: QuestionType.SINGLE_CHOICE,
                options: [],
                nextQuestion: { conditions: [], default: undefined }
            } as SingleChoiceQuestion;

        case QuestionType.MULTIPLE_CHOICE:
            return {
                id: newQuestionId,
                question: '',
                subtext: '',
                image: '',
                type: QuestionType.MULTIPLE_CHOICE,
                options: [],
                maxSelections: undefined,
                nextQuestion: { conditions: [], default: undefined }
            } as MultipleChoiceQuestion;

        case QuestionType.INPUT:
            return {
                id: newQuestionId,
                question: '',
                subtext: '',
                image: '',
                type: QuestionType.INPUT,
                placeholder: '',
                inputType: 'text', // Default input type
                validation: undefined,
                nextQuestion: { conditions: [], default: undefined }
            } as InputQuestion;

        default:
            throw new Error('Invalid question type');
    }
}

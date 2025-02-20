import type { ValidatorName } from "$lib/validators/formValidators";

export type Question = 
  | SingleChoiceQuestion 
  | MultipleChoiceQuestion 
  | InputQuestion 
  | RatingQuestion;

export enum QuestionType {
  SINGLE_CHOICE = 'single-choice',
  MULTIPLE_CHOICE = 'multiple-choice',
  INPUT = 'input',
  RATING = 'rating'
}

export interface BaseQuestion {
  id: number;
  question: string;
  subtext: string;
  image: string;
  nextQuestion?: {
    conditions: {
      answer: string;
      nextId: number;
    }[];
    default?: number;
  };
}

export interface QuestionOption {
  id: number;
  text: string;
  icon?: string;
  image?: string;
}
  
export interface SingleChoiceQuestion extends BaseQuestion {
  type: QuestionType.SINGLE_CHOICE;
  options: QuestionOption[];
}

export interface MultipleChoiceQuestion extends BaseQuestion {
  type: QuestionType.MULTIPLE_CHOICE;
  options: QuestionOption[];
  maxSelections?: number;
}

export interface InputQuestion extends BaseQuestion {
  type: QuestionType.INPUT;
  placeholder?: string;
  inputType: 'text' | 'email' | 'number';
  validation?: ValidatorName;
}

export interface RatingQuestion extends BaseQuestion {
  type: QuestionType.RATING;
  minValue: number;
  maxValue: number;
  step?: number;
  showLabels?: boolean;
  minLabel?: string;
  maxLabel?: string;
  icon?: string;  // For example: ⭐️ or 👍
}
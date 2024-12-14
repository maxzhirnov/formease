// types.ts
export interface QuestionOption {
    text: string;
    icon: string;
    color: string;
  }
  
interface BaseQuestion {
  id: number;
  question: string;
  subtext: string;
  image: string;
  gradient: string;
  bgColor: string;
  nextQuestion?: {
    conditions: {
      answer: string;
      nextId: number;
    }[];
    default?: number;
  };
}
  
  export interface SingleChoiceQuestion extends BaseQuestion {
    type: 'single-choice';
    options: QuestionOption[];
  }

  export interface MultipleChoiceQuestion extends BaseQuestion {
    type: 'multiple-choice';
    options: QuestionOption[];
    maxSelections?: number;
  }

  export interface InputQuestion extends BaseQuestion {
    type: 'input';
    placeholder?: string;
    inputType: 'text' | 'email' | 'number';
    validation?: RegExp;
  }
  
  
  export type Question = SingleChoiceQuestion | MultipleChoiceQuestion | InputQuestion;
  

  export interface Theme {
    backgroundColor?: string;
    fontColor?: string;
  }

  export interface Button {
    text: string;
    url?: string;
    newTab?: boolean;
  }

  export interface ThankYouMessage {
    title?: string;
    subtitle?: string;
    icon?: string;
    button?: Button;
  }

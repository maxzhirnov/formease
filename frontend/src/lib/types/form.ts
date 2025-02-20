import type { Question } from "./questions";
import type { ThankYouMessage } from "./thank";
import type { ThemeName } from "./theme";

export interface FormData {
  id: string;
  isDraft: boolean;
  name: string;
  theme: ThemeName;
  userId: string;
  floatingShapesTheme: string;
  questions: Question[];
  thankYouMessage: ThankYouMessage;
}

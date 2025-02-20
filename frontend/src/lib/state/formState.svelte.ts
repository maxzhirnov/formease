// src/lib/state/formState.ts

import { type FormData } from '$lib/types/form';
import { floatingShapesOptionsList } from '$lib/types/shapes';
import { defaultThankYouMessage } from '$lib/types/thank';
import { themeOptionsList } from '$lib/types/theme';

// Initial form data structure with default values
const initialFormData: FormData = {
  id: '',
  isDraft: true,
  name: '',
  userId: '',
  theme: themeOptionsList[1].value,
  floatingShapesTheme: floatingShapesOptionsList[0].value,
  questions: [],
  thankYouMessage: defaultThankYouMessage
};

function createFormState(initialFormData: FormData) {
  let formState = $state(initialFormData);
  
  return {
    get state(): Readonly<FormData> {
      return formState;
    },
    update(newState: Partial<FormData>) {
      formState = { ...formState, ...newState };
    },
    reset() {
      formState = { ...initialFormData };
    }
  };
}

// Create a writable store with initial form data
export const formState = createFormState(initialFormData);
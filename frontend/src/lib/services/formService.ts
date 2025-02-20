import { type FormApiService, formApiService } from './formApiService';
import { type FormStateService, formStateService } from './formStateService';

class FormService {
    constructor(
      public state: FormStateService,
      public api: FormApiService
    ) {}
  
  }
  
  export const formService = new FormService(formStateService, formApiService);
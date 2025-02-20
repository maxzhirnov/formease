<script lang="ts">
  import { Label, Input, Helper, Button } from 'flowbite-svelte';

  import { validate, type ValidatorName } from '$lib/validators/formValidators';
    import type { InputQuestion, Question } from '$lib/types/questions';
  
  export let question: InputQuestion;
  export let onAnswer: (value: string) => void;
  export let onNext: (question: Question) => void;

  const currentQuestion = question;
  let inputValue = '';
  let isValid = false;
  let errorMessage = '';

  function validateInput(value: string) {
    if (question.validation) {
      console.log('validateInput', value, question);
      isValid = validate(value, question.validation);
      errorMessage = isValid ? '' : `Please enter a valid ${question.inputType}`;
    } else {
      isValid = value.length > 0;
      errorMessage = isValid ? '' : 'This field is required';
    }
  }

  function handleInput(e: Event) {
    inputValue = (e.target as HTMLInputElement).value;
    validateInput(inputValue);
  }

  function handleContinue() {
    if (!isValid) return;
    onAnswer(inputValue);
    onNext(currentQuestion);
  }
</script>

<div class="space-y-4">
  <Label>
      <Input 
          type={question.inputType}
          placeholder={question.placeholder}
          color={!isValid && inputValue.length > 0 ? 'red' : 'base'}
          bind:value={inputValue}
          on:input={handleInput}
      />
  </Label>
  
  {#if errorMessage && inputValue.length > 0}
      <Helper color="red" class="mt-2">{errorMessage}</Helper>
  {/if}

  <Button 
      class="mt-4"
      color="blue"
      disabled={!isValid}
      on:click={handleContinue}
  >
      Next
  </Button>
</div>


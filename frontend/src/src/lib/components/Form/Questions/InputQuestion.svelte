<script lang="ts">
  import { fade } from 'svelte/transition';
  import type { InputQuestion, Question } from '$lib/types';
  
  export let question: InputQuestion;
  export let onAnswer: (value: string) => void;
  export let onNext: (question: Question) => void;

  const currentQuestion = question;
  let inputValue = '';
  let isValid = false;
  let errorMessage = '';

  function validateInput(value: string) {
    if (question.validation) {
      isValid = question.validation.test(value);
      errorMessage = isValid ? '' : 'Please enter a valid value';
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

<div class="input-container">
  <input 
    type={question.inputType}
    class="input-field"
    class:error={!isValid && inputValue.length > 0}
    placeholder={question.placeholder}
    value={inputValue}
    on:input={handleInput}
  />
  
  {#if errorMessage && inputValue.length > 0}
    <p class="error-message" in:fade>
      {errorMessage}
    </p>
  {/if}

  <button 
    class="continue-button"
    class:disabled={!isValid}
    disabled={!isValid}
    on:click={handleContinue}
    in:fade
  >
    Next
  </button>
</div>

<style>
  .input-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .input-field {
    padding: 1.2rem;
    background: rgba(255, 255, 255, 0.1);
    border: 1px solid rgba(105, 92, 92, 0.285);
    border-radius: 16px;
    color: inherit;
    font-size: 1.1rem;
    width: 100%;
    transition: all 0.3s ease;
  }

  .input-field.error {
    border-color: rgba(255,100,100,0.5);
  }

  .error-message {
    color: rgba(255,100,100,0.9);
    font-size: 0.9rem;
    margin: 0;
  }

  .continue-button {
    margin-top: 1rem;
    padding: 1.2rem;
    background: linear-gradient(to right, var(--gradient-colors));
    border: none;
    border-radius: 16px;
    color: inherit;
    font-size: 1.1rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    box-shadow: 0 4px 12px rgba(0,0,0,0.2);
  }

  .continue-button:hover:not(.disabled) {
    transform: translateY(-2px);
    box-shadow: 0 8px 20px rgba(0,0,0,0.3);
  }

  .continue-button.disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  @media (max-width: 768px) {
    .input-field {
      width: 90%;
    }
  }
</style>

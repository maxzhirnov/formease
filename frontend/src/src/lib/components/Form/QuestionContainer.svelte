<!-- QuestionContainer.svelte -->
<script lang="ts">
  import type { Question } from '$lib/types';
  import { fade, slide } from 'svelte/transition';
  import InputQuestion from './Questions/InputQuestion.svelte';
  import SingleChoiceQuestion from './Questions/SingleChoiceQuestion.svelte';
  import MultipleChoiceQuestion from './Questions/MultipleChoiceQuestion.svelte';

  export let question: Question;
  export let answers: string | string[] | undefined;
  export let onAnswer: (value: string) => void;
  export let onNavigate: (nextId: number) => void;

  const handleNext = (q: Question) => {
      if (q.nextQuestion?.default) {
          onNavigate(q.nextQuestion.default);
      }
  };
</script>

<div class="question-container" 
  in:fade={{ duration: 400 }}
  style="--gradient-colors: {question.gradient};"
>
  <div class="content">
      <h2 in:slide={{ duration: 400, delay: 200 }}>
          {question.question}
      </h2>
      <p class="subtext" in:fade={{ duration: 400, delay: 300 }}>
          {question.subtext}
      </p>

      <div class="options">
        {#if question.type === 'input'}
            <InputQuestion 
                {question}
                {onAnswer}
                onNext={handleNext}
            />
        {:else if question.type === 'single-choice'}
            <SingleChoiceQuestion
                {question}
                answer={answers as string}
                onSelect={onAnswer}
            />
        {:else if question.type === 'multiple-choice'}
            <MultipleChoiceQuestion
                {question}
                answers={answers as string[]}
                onSelect={onAnswer}
                onNext={handleNext}
            />
        {/if}
    </div>
    
  </div>

  <div class="image-container" in:fade={{ duration: 600, delay: 300 }}>
      <img src={question.image} alt="Question illustration" />
  </div>
</div>

<style>
  .question-container {
    display: grid;
    grid-template-columns: 1.2fr 1fr;
    gap: 4rem;
    align-items: center;
    min-height: 500px;
  }

  .image-container {
    position: relative;
    overflow: hidden;
    border-radius: 24px;
    box-shadow: 0 20px 40px rgba(0,0,0,0.3);
    aspect-ratio: 4/5;
  }

  .image-container::after {
    content: '';
    position: absolute;
    inset: 0;
    background: linear-gradient(to bottom right, var(--gradient-colors));
    opacity: 0.6;
    mix-blend-mode: soft-light;
  }

  .image-container img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.7s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .image-container:hover img {
    transform: scale(1.05);
  }

  .options {
    display: grid;
    gap: 1.2rem;
  }

  h2 {
    font-family: 'Cal Sans', serif;
    margin-bottom: 1rem;
    font-size: 2.5rem;
    line-height: 1.2;
    font-weight: 600;
    /* background: linear-gradient(120deg, #fff, rgba(255,255,255,0.8)); */
    -webkit-background-clip: text;
    background-clip: text;
    color: inherit;
    text-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }

  .subtext {
    color: inherit;
    font-size: 1.1rem;
    margin-bottom: 2.5rem;
    font-weight: 400;
  }

  
  @media (max-width: 768px) {
    .question-container {
      grid-template-columns: 1fr;
      gap: 2rem;
      min-height: auto;
    }

    .image-container {
      aspect-ratio: 16/9;
      order: -1;
    }

    h2 {
      font-size: 2rem;
      text-align: center;
    }

    .subtext {
      text-align: center;
      margin-bottom: 2rem;
    }
  }

  @media (max-width: 480px) {
    h2 {
      font-size: 1.75rem;
    }

    .subtext {
      font-size: 1rem;
    }
  }
</style>
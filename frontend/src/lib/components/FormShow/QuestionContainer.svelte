<script lang="ts">
  import { onMount } from 'svelte';
  import { fade, slide } from 'svelte/transition';
  import InputQuestion from './Questions/InputQuestion.svelte';
  import SingleChoiceQuestion from './Questions/SingleChoiceQuestion.svelte';
  import MultipleChoiceQuestion from './Questions/MultipleChoiceQuestion.svelte';
  import RateQuestion from './Questions/RateQuestion.svelte';
  import type { Question } from '$lib/types/questions';

  export let question: Question;
  export let answers: string | string[] | undefined;
  export let onAnswer: (value: string) => void;
  export let onNavigate: (nextId: number | undefined) => void;
  let imageExists = true;

  const handleNext = (q: Question) => {
      if (q.nextQuestion?.default) {
          onNavigate(q.nextQuestion.default);
      } else {
          onNavigate(undefined);
      }
  };

  onMount(() => {
        if (question.image) {
            // Check if image exists
            fetch(`${question.image}`)
                .then(response => {
                    imageExists = response.ok;
                })
                .catch(() => {
                    imageExists = false;
                });
        }
    });
</script>

<div 
  class="grid grid-cols-1 md:grid-cols-2 gap-16 min-h-[500px] items-center" 
  in:fade={{ duration: 400 }}
>
  <div class="space-y-6">
    <h2 
      class="font-oswald text-3xl md:text-4xl font-semibold" 
      in:slide={{ duration: 400, delay: 200 }}
    >
      {question.question}
    </h2>
    <p 
      class="text-lg" 
      in:fade={{ duration: 400, delay: 300 }}
    >
      {question.subtext}
    </p>

    <div class="grid gap-4">
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
      {:else if question.type === 'rating'}
          <RateQuestion
              {question}
              onAnswer={onAnswer}
              onNext={handleNext}
          />
      {/if}
    </div>
  </div>
  
  <div 
    class="relative overflow-hidden rounded-3xl shadow-2xl aspect-[4/5] md:aspect-[4/5] hover:scale-105 transition-transform duration-700 ease-cubic-bezier" 
    in:fade={{ duration: 600, delay: 300 }}
>
    {#if question.image && imageExists}
        <img 
            src={`${question.image}`} 
            alt="Question illustration" 
            class="absolute inset-0 w-full h-full object-cover"
        />
    {:else}
        <img 
            src={`https://api.dicebear.com/7.x/shapes/svg?seed=${question.id}-${question.type}-${question.question.slice(0, 5)}`}
            alt="Generated placeholder" 
            class="absolute inset-0 w-full h-full object-cover"
        />
    {/if}
    <div class="absolute inset-0 bg-gradient-to-br from-blue-500/60 to-purple-500/60 mix-blend-soft-light"></div>
  </div>
</div>

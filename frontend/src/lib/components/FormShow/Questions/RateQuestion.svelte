<script lang="ts">
    import type { Question, RatingQuestion } from '$lib/types/questions';
    import { Button } from 'flowbite-svelte';
    import { StarSolid, StarOutline } from 'flowbite-svelte-icons';
    
    
    export let question: RatingQuestion;
    console.log('question', question);
    export let onAnswer: (value: string) => void;
    export let onNext: (question: Question) => void;
  
    const currentQuestion = question;
    let selectedValue: number | null = null;
    let isHovering: number | null = null;
  
    function handleSelect(value: number) {
      selectedValue = value;
    }
  
    function handleHover(value: number | null) {
      isHovering = value;
    }
  
    function handleContinue() {
      if (selectedValue === null) return;
      onAnswer(selectedValue.toString());
      onNext(currentQuestion);
    }
  
    // Generate array of numbers from min to max with step
    $: ratingValues = Array.from(
      { length: Math.floor((question.maxValue - question.minValue) / (question.step || 1)) + 1 },
      (_, i) => question.minValue + i * (question.step || 1)
    );

    $: starSize = ratingValues.length > 7 
        ? ratingValues.length > 12
        ? 'w-4 h-4 sm:w-5 sm:h-5' // For 13+ stars
        : 'w-5 h-5 sm:w-6 sm:h-6' // For 8-12 stars
        : 'w-6 h-6 sm:w-8 sm:h-8'; // For 7 or fewer stars

    $: iconSize = ratingValues.length > 7
        ? ratingValues.length > 12
        ? 'text-lg sm:text-xl' // For 13+ stars
        : 'text-lg sm:text-xl' // For 8-12 stars
        : 'text-2xl sm:text-3xl'; // For 7 or fewer stars

    $: gapSize = ratingValues.length > 7 ? 'gap-1' : 'gap-2';
  </script>
  
  <div class="w-full max-w-2xl mx-auto">
    <div class="space-y-6">
      <!-- Rating Stars Container -->
      <div class="flex flex-col items-center gap-6 p-4">
        <!-- Stars Row -->
        <div class="flex flex-wrap justify-center {gapSize} max-w-md mx-auto">
            {#each ratingValues as value}
              <button
                class="p-1.5 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700 
                       transform hover:scale-110 transition-all duration-200 ease-in-out
                       focus:outline-none focus:ring-2 focus:ring-blue-300 dark:focus:ring-blue-600"
                on:click={() => handleSelect(value)}
                on:mouseenter={() => handleHover(value)}
                on:mouseleave={() => handleHover(null)}
              >
                {#if question.icon}
                  <span class="{iconSize} transition-opacity duration-200"
                        class:opacity-100={value <= (isHovering ?? selectedValue ?? 0)}
                        class:opacity-40={value > (isHovering ?? selectedValue ?? 0)}>
                    {question.icon}
                  </span>
                {:else}
                  {#if value <= (isHovering ?? selectedValue ?? 0)}
                    <StarSolid 
                      class="{starSize} text-yellow-300 transition-transform duration-200" 
                    />
                  {:else}
                    <StarOutline 
                      class="{starSize} text-gray-300 hover:text-gray-400 transition-colors duration-200" 
                    />
                  {/if}
                {/if}
              </button>
            {/each}
          </div>
  
  
        <!-- Labels -->
        {#if question.showLabels}
          <div class="flex justify-between w-full px-4 text-sm">
            <span class="text-inherit dark:text-gray-400 font-medium">
              {question.minLabel || question.minValue}
            </span>
            <span class="text-inherit dark:text-gray-400 font-medium">
              {question.maxLabel || question.maxValue}
            </span>
          </div>
        {/if}
  

      </div>
  
      <!-- Next Button -->
      <Button 
        class="w-full 
        bg-gradient-to-r 
        from-blue-500 
        to-purple-500 
        hover:from-blue-600 
        hover:to-purple-600 
        transition-all 
        duration-300 
        ease-cubic-bezier"
        disabled={selectedValue === null}
        on:click={handleContinue}
      >
        Continue
      </Button>
    </div>
</div>
  
  <style>
    button {
      appearance: none;
      border: none;
      background: none;
      cursor: pointer;
    }
  </style>
<script lang="ts">
    import type { MultipleChoiceQuestion, Question } from '$lib/types';
    import { Button } from 'flowbite-svelte';
    
    export let question: MultipleChoiceQuestion;
    export let answers: string[] = [];
    export let onSelect: (value: string) => void;
    export let onNext: (question: Question) => void;

    // Compute if selection is allowed
    function canSelect(optionText: string): boolean {
        // If no max selections set, allow unlimited
        if (!question.maxSelections) return true;

        // If option already selected, allow deselection
        if (answers.includes(optionText)) return true;

        // Check if we've reached max selections
        return answers.length < question.maxSelections;
    }
</script>

<div class="grid gap-4">
    {#each question.options as option, i}
        <Button 
            disabled={!canSelect(option.text)}
            class="w-full justify-start gap-4 
                   {answers.includes(option.text) 
                     ? 'bg-blue-500/30 text-white-500 border-blue-500/40 scale-[1.02]' 
                     : 'bg-white/10 hover:bg-white/20 border-white/10'}
                   {!canSelect(option.text) ? 'opacity-50 cursor-not-allowed' : ''}
                   hover:scale-[1.02] 
                   border 
                   shadow-md 
                   transition-all 
                   duration-300 
                   ease-cubic-bezier"
            on:click={() => onSelect(option.text)}
        >
            {#if option.icon}
                <span>{option.icon}</span>
            {/if}
            {option.text}
            {#if question.maxSelections}
                <span class="ml-auto text-xs text-gray-500">
                    {answers.length} / {question.maxSelections}
                </span>
            {/if}
        </Button>
    {/each}

    {#if answers.length > 0}
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
            on:click={() => onNext(question)}
        >
            Next
        </Button>
    {/if}
</div>

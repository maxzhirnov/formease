<script lang="ts">
    import type { SingleChoiceQuestion } from '$lib/types';
    import { cubicOut } from 'svelte/easing';
    import { fly } from 'svelte/transition';
    import { Button } from 'flowbite-svelte';
    
    export let question: SingleChoiceQuestion;
    export let answer: string | undefined;
    export let onSelect: (value: string) => void;
</script>

<div class="grid gap-4">
    {#each question.options as option, i}
        <Button 
            class="w-full justify-start gap-4 
                   bg-white/10 
                   hover:bg-white/20 
                   hover:scale-[1.02] 
                   border 
                   border-white/10 
                   shadow-md 
                   transition-all 
                   duration-300 
                   ease-cubic-bezier 
                   {answer === option.text ? 'bg-white/30 scale-[1.02] border-white/40' : ''}"
            on:click={() => onSelect(option.text)}
        >
            {#if option.icon}
                <span>{option.icon}</span>
            {/if}
            {option.text}
        </Button>
    {/each}
</div>

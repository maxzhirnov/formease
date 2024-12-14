<script lang="ts">
    import type { MultipleChoiceQuestion, Question } from '$lib/types';
    import { cubicOut } from 'svelte/easing';
    import { fade, fly } from 'svelte/transition';
    
    export let question: MultipleChoiceQuestion;
    export let answers: string[] = [];
    export let onSelect: (value: string) => void;
    export let onNext: (question: Question) => void;
</script>

<div class="options">
    {#each question.options as option, i}
        <button 
            class="option-button"
            class:selected={answers.includes(option.text)}
            style="background: linear-gradient(120deg, {option.color}33, {option.color}22);"
            on:click={() => onSelect(option.text)}
            in:fly={{ y: 20, delay: 400 + (i * 100), duration: 500, easing: cubicOut }}
        >
            <span class="option-icon">{option.icon}</span>
            {option.text}
        </button>
    {/each}

    {#if answers.length > 0}
        <button 
            class="next-button"
            on:click={() => onNext(question)}
            in:fade={{ duration: 400 }}
        >
            Next
        </button>
    {/if}
</div>

<style>
    .options {
        display: grid;
        gap: 1.2rem;
    }

    .option-button {
        padding: 1.2rem;
        border: none;
        background: rgba(255, 255, 255, 0.1);
        border-radius: 16px;
        box-shadow: 0 4px 12px rgba(0,0,0,0.2), inset 0 1px 2px rgba(255,255,255,0.1);
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        cursor: pointer;
        display: flex;
        align-items: center;
        gap: 1rem;
        font-size: 1.1rem;
        font-weight: 500;
        color: inherit;
        border: 1px solid rgba(255,255,255,0.1);
    }

    .option-button:hover {
        transform: translateY(-2px) scale(1.02);
        background: rgba(255, 255, 255, 0.2);
        box-shadow: 0 8px 20px rgba(0,0,0,0.3), inset 0 2px 4px rgba(255,255,255,0.1);
    }

    .option-button.selected {
        background: rgba(255, 255, 255, 0.3) !important;
        transform: scale(1.02);
        border-color: rgba(255, 255, 255, 0.4);
    }

    .next-button {
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
        width: 100%;
    }

    .next-button:hover {
        transform: translateY(-2px);
        box-shadow: 0 8px 20px rgba(0,0,0,0.3);
    }
</style>

<script lang="ts">
    import { Button, Alert } from 'flowbite-svelte';
    import { WandMagicSparklesSolid } from 'flowbite-svelte-icons';

    interface Props {
        onSuccess?: (formId: string) => void;
        createAIFormCallback: (description: string) => Promise<void>;
    }

    let { createAIFormCallback }: Props = $props();

    let description: string = $state('');
    let error: string | null = $state(null);
    let isGenerating: boolean = $state(false);

    async function handleGenerate() {
        if (!description.trim()) {
            error = 'Please provide a description for your form';
            return;
        }
        isGenerating = true;

        try {
            error = null;
            await createAIFormCallback(description);
        } catch (err) {
            error = err instanceof Error 
                ? err.message 
                : 'Failed to generate form with AI';
        } finally {
            isGenerating = false;
        }
    }
</script>

<div class="magic-container p-6">
    <div class="relative space-y-6">
        <div class="mb-0 sm:mb-4 text-center">
            <h3 class="text-2xl font-bold magic-title">
                AI Form Generator
            </h3>

            <div class="magic-stars">
                <div class="star magic-star1">⭐</div>
                <div class="star magic-star2">⭐</div>
                <div class="star magic-star3">⭐</div>
            </div>
            <p class="text-center mt-2 text-purple-600/80">Describe your form and let AI craft it for you</p>
        </div>

        <textarea 
            bind:value={description}
            placeholder="Describe the purpose and content of your form. For example: 'A customer satisfaction survey for a restaurant'"
            disabled={isGenerating}
            rows={5}
            class="magic-textarea"
        ></textarea>

        <button 
            class="magic-button {isGenerating ? 'generating' : ''}"
            onclick={handleGenerate} 
            disabled={isGenerating}
        >
            <span class="default-content">
                <WandMagicSparklesSolid class="w-5 h-5 mr-2" />
                Generate with AI
            </span>
            <span class="generating-content">
                <div class="magic-pulse"></div>
                Crafting Your Form...
            </span>
        </button>
    </div>
    
    {#if error}
        <Alert color="red" class="mt-4">
            {error}
        </Alert>
    {/if}
</div>

<style>
    .magic-container {
        background: linear-gradient(to bottom right, #ffffff, #faf5ff);
        border-radius: 1rem;
        box-shadow: 0 0 30px rgba(147, 51, 234, 0.1);
    }

    .magic-title {
        position: relative;
        background: linear-gradient(90deg, #7928CA, #FF0080);
        background-clip: text;
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        display: inline-block;
    }

    .magic-stars {
        pointer-events: none;
    }

    .magic-star1 {
        position: absolute;
        top: -10px;
        right: 20px;
        font-size: .8rem;

    }

    .magic-star2 {
        position: absolute;
        top: 35px;
        left: -10px;
        font-size: 1rem;
    }

    .magic-star3 {
        position: absolute;
        top: 60px;
        right: 30px;
        font-size: 1.2rem;

    }

    .star {
        animation: star-bounce 1.5s ease infinite;
    }

    .star:nth-child(1) { 
        animation-delay: 0s; 
    }
    .star:nth-child(2) { 
        animation-delay: 0.2s; 
    }
    .star:nth-child(3) { 
        animation-delay: 0.4s; 
    }

    @keyframes star-bounce {
        0%, 100% {
            transform: translateY(0) scale(1);
            opacity: 0.8;
        }
        50% {
            transform: translateY(-5px) scale(1.2);
            opacity: 1;
        }
    }

    textarea.magic-textarea {
        width: 100%;
        min-height: 120px;
        padding: 1rem;
        border: 2px solid rgba(147, 51, 234, 0.2);
        border-radius: 0.75rem;
        background: white;
        transition: all 0.3s ease;
        font-size: 1rem;
        line-height: 1.5;
        resize: none;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    }

    textarea.magic-textarea:focus {
        outline: none;
        border-color: #9333EA;
        box-shadow: 0 0 0 4px rgba(147, 51, 234, 0.1);
    }

    textarea.magic-textarea::placeholder {
        color: #a1a1aa;
    }

    textarea.magic-textarea:disabled {
        background-color: #f9fafb;
        cursor: not-allowed;
    }

    .magic-button {
        position: relative;
        width: 100%;
        padding: 0.875rem;
        border-radius: 0.75rem;
        background: linear-gradient(45deg, #7928CA, #FF0080);
        color: white;
        font-weight: 600;
        font-size: 1rem;
        border: none;
        cursor: pointer;
        overflow: hidden;
        transition: all 0.3s ease;
    }

    .magic-button:not(:disabled):hover {
        transform: translateY(-2px);
        box-shadow: 0 8px 20px rgba(147, 51, 234, 0.3);
    }

    .magic-button:not(:disabled):active {
        transform: translateY(0);
    }

    .magic-button:disabled {
        opacity: 0.7;
        cursor: not-allowed;
    }

    .default-content,
    .generating-content {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: 100%;
        transition: all 0.3s ease;
    }

    .generating-content {
        position: absolute;
        top: 0;
        left: 0;
        opacity: 0;
    }

    .magic-button.generating .default-content {
        opacity: 0;
    }

    .magic-button.generating .generating-content {
        opacity: 1;
    }

    .magic-pulse {
        width: 20px;
        height: 20px;
        background: white;
        border-radius: 50%;
        margin-right: 0.75rem;
        position: relative;
    }

    .magic-pulse::before,
    .magic-pulse::after {
        content: '';
        position: absolute;
        width: 100%;
        height: 100%;
        border-radius: 50%;
        background: white;
        animation: pulse 2s ease-out infinite;
    }

    .magic-pulse::after {
        animation-delay: 1s;
    }

    @keyframes pulse {
        0% {
            transform: scale(1);
            opacity: 0.8;
        }
        100% {
            transform: scale(3);
            opacity: 0;
        }
    }

    @keyframes star-bounce {
        0%, 100% {
            transform: translateY(0) scale(1);
            opacity: 0.8;
        }
        50% {
            transform: translateY(-5px) scale(1.2);
            opacity: 1;
        }
    }

    /* Add sparkles around button on hover */
    .magic-button:not(:disabled):hover::before {
        content: '✨';
        position: absolute;
        animation: sparkle 1s linear infinite;
    }

    @keyframes sparkle {
        0% {
            transform: translate(30px, 0) rotate(0deg);
            opacity: 0;
        }
        50% {
            opacity: 1;
        }
        100% {
            transform: translate(-30px, -30px) rotate(180deg);
            opacity: 0;
        }
    }

    /* Loading state animation */
    .magic-button.generating {
        background: linear-gradient(45deg, #7928CA, #FF0080, #7928CA);
        background-size: 200% 200%;
        animation: gradient 2s ease infinite;
    }

    @keyframes gradient {
        0% {
            background-position: 0% 50%;
        }
        50% {
            background-position: 100% 50%;
        }
        100% {
            background-position: 0% 50%;
        }
    }
</style>

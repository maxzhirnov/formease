<script lang="ts">
    import { Button, Input, Textarea, Alert, Spinner } from 'flowbite-svelte';
    import { WandMagicSparklesSolid } from 'flowbite-svelte-icons';

    interface Props {
        createWithAICallback: (description: string) => Promise<void>;
    }

    let { createWithAICallback }: Props = $props();

    let description: string = $state('');
    let error: string | null = $state(null);
    let isGenerating: boolean = $state(false);

    async function handleGenerate() {
        if (!description.trim()) {
            error = 'Please provide a description for your form';
            return;
        }

        try {
            isGenerating = true;
            error = null;
            await createWithAICallback(description);
        } catch (err) {
            error = err instanceof Error 
                ? err.message 
                : 'Failed to generate form with AI';
        } finally {
            isGenerating = false;
        }
    }
</script>

<div class="space-y-4">
    <div class="flex items-center space-x-2">
        <Textarea 
            bind:value={description}
            placeholder="Describe the purpose and content of your form. For example: 'A customer satisfaction survey for a restaurant'"
            disabled={isGenerating}
            class="flex-grow"
            rows={3}
        />
        <Button 
            color="purple" 
            size="md"
            on:click={handleGenerate} 
            disabled={isGenerating}
            class="h-full w-[180px]"
        >
            {#if isGenerating}
                <Spinner class="mr-2" size="4" />
                Generating...
            {:else}
                <WandMagicSparklesSolid class="mr-2" />
                AI Generate 
            {/if}
        </Button>
    </div>
    
    {#if error}
        <Alert color="red">
            {error}
        </Alert>
    {/if}
</div>

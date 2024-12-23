<script lang="ts">
    import { showError } from '$lib/utils/errorHandler';
    import { Button, Input, Alert, Modal } from 'flowbite-svelte';
    import { CirclePlusSolid } from 'flowbite-svelte-icons';
    import CreateAi from './CreateAI.svelte';

    interface Props {
        createFormCallback: (name: string) => Promise<void>;
        createAICallback: (description: string) => Promise<void>;
    }

    let { createFormCallback, createAICallback }: Props = $props();

    

    let showModal = $state(false);
    let formName: string = $state('')
    let error: string | null = $state(null)
    let isSubmitting: boolean = $state(false)

    $effect(() => {
        if (error !== null) {
            showError(error)
        }
    })

    async function handleSubmit() {
        if (!formName.trim()) {
            error = 'Form name is required';
            return;
        }

        try {
            isSubmitting = true;
            error = null;
            createFormCallback(formName)
            formName = '';
            showModal = false;
        } catch (err) {
            error = err instanceof Error ? err.message : 'Failed to create form';
        } finally {
            isSubmitting = false;
        }
    }

    function handleCancel() {
        showModal = false;
        formName = '';
        error = null;
    }
</script>

<Modal 
    bind:open={showModal}
    title="Create new form"
    size="md"
    placement="center"
    outsideclose={true}
    >
    <form 
            onsubmit={(e) => {
                e.preventDefault(); 
                handleSubmit(); 
            }} 
            class="flex items-center space-x-2"
        >
            <Input 
                type="text"
                bind:value={formName}
                placeholder="Enter form name"
                required
                disabled={isSubmitting}
                class="flex-grow"
            />
            <Button 
                type="submit" 
                color="green" 
                disabled={isSubmitting}
            >
                {isSubmitting ? 'Creating...' : 'Create'}
            </Button>
            <Button 
                type="button" 
                color="light" 
                on:click={handleCancel}
                disabled={isSubmitting} 
            >
                Cancel
            </Button>
        </form>
        <CreateAi createWithAICallback={ createAICallback }/>
</Modal>

<Button 
    size="md"
    color="blue" 
    on:click={() => {
        showModal = true
        }
    }
>
    <CirclePlusSolid class="mr-2" />
    New Form
</Button>

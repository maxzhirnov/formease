<script lang="ts">
    import { showError } from '$lib/utils/errorHandler';
    import { Button, GradientButton, Input, Modal, Tabs, TabItem } from 'flowbite-svelte';
    import { CirclePlusSolid, PenSolid, WandMagicSparklesOutline } from 'flowbite-svelte-icons';
    import CreateAi from './CreateAI.svelte';

    interface Props {
        createFormCallback: (name: string) => Promise<void>;
        createAIFormCallback: (description: string) => Promise<void>;
    }

    let { createFormCallback, createAIFormCallback }: Props = $props();

    
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
    <Tabs tabStyle="underline">
        <TabItem open>
            <div slot="title" class="flex items-center gap-2">
                <PenSolid size="md" />
                Manual Creation
            </div>
            <form 
                onsubmit={(e) => {
                    e.preventDefault(); 
                    handleSubmit(); 
                }} 
                class="flex items-center space-x-2 mt-4"
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
        </TabItem>
        
        <TabItem>
            <div slot="title" class="flex items-center gap-2">
                <WandMagicSparklesOutline size="md" />
                AI Assistant
            </div>
            <CreateAi createAIFormCallback={createAIFormCallback}/>
        </TabItem>
    </Tabs>
</Modal>

<GradientButton 
    size="md"
    color="purpleToBlue" 
    on:click={() => {
        showModal = true
        }
    }
>
    <CirclePlusSolid class="mr-1" />
    New Form
</GradientButton>

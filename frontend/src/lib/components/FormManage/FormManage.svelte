<script lang="ts">
    import { Alert, Button, Heading, Skeleton } from 'flowbite-svelte';
    import { CirclePlusSolid } from 'flowbite-svelte-icons';

    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';

    import { formService } from '$lib/services/formService';
    import { listForms } from '$lib/api/forms';
    import { showError } from '$lib/utils/errorHandler';
    import type { FormData } from '$lib/types';

    import Create from './Create.svelte';
    import Card from './Card.svelte';
    import CreateAi from './CreateAI.svelte';
    
    let forms: FormData[] = $state([]);
    let isLoading = $state(true);
    let error: string | null = $state(null);

    async function loadForms() {
        try {
            isLoading = true;
            error = null;
            forms = await listForms();
        } catch (err) {
            const errMessage = err instanceof Error ? err.message : 'Failed to load forms'
            error = errMessage;
            showError(errMessage);
        } finally {
            isLoading = false;
        }
    }

    async function handleCreateForm(name: string) {
        formService.resetForm();
        formService.updateFormName(name);
        await formService.create()
        loadForms();
    }

    async function handleCreateWithAI(description: string) {
        try {
            const newForm = await formService.generateAIForm(description);
            loadForms();
        } catch (error) {
            // Handle generation errors
            const errorMessage =  error instanceof Error ? error.message : 'Failed to generate form with AI'
            showError(errorMessage);
        }
    }
    
    async function handleDeleteForm(id: string) {
        if (!confirm('Are you sure you want to delete this form?')) {
            return;
        }
        try {
            error = null;
            await formService.delete(id)
        } catch (err) {
            error = err instanceof Error ? err.message : 'Failed to delete form';
            // showError("eerrrrorr");
        } finally {
            loadForms()
        }
    }

    function handleEditForm(id: string) {
        goto(`/app/form/${id}/edit`);
    }

    onMount(() => {
        loadForms();
    });
</script>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    
    <!-- Title and refresh -->
    <div class="flex justify-between items-center mb-8 flex-col sm:flex-row">
        <div class="mb-4 sm:mb-0">
            <Heading 
            tag="h2" 
            class="text-3xl font-extrabold text-gray-900 dark:text-white">
            Your Forms
            </Heading>
        </div>
        <div class="flex items-center space-x-2">
            <div>
                <Create 
                    createFormCallback={handleCreateForm} 
                    createAICallback={handleCreateWithAI}
                 />
            </div>
    
            <div>
                <Button 
                    color="blue" 
                    class="w-[140px]" 
                    onclick={loadForms} 
                    disabled={isLoading}>
                    {#if isLoading}
                        Loading...
                    {:else}
                        Refresh
                    {/if}
                </Button>
            </div>
        </div>
    </div>

    <!-- Error message -->
    {#if error}
        <Alert color="red" class="mt-4">
            <Button 
                color="red" 
                class="mr-4" 
                onclick={loadForms}
            >
                Try Again
            </Button>
            {error}
        </Alert>
    {/if}
    
    <!-- Loading state -->
    {#if isLoading}
    <!-- Skeleton -->
        <div class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-3 gap-6">
            {#each Array(6) as _}
                <div class="space-y-4">
                    <Skeleton/>
                    <div class="space-y-2">
                        <Skeleton size="md" />
                        <Skeleton size="sm" />
                    </div>
                </div>
            {/each}
        </div>
    {:else if !forms}
    <!-- Empty state -->
        <div class="text-center py-16">
            <p class="text-xl text-gray-500 mb-6">
                You haven't created any forms yet.
            </p>
            <Button 
                color="blue" 
                onclick={() => goto('/app/create')}
                class="mx-auto"
            >
                <CirclePlusSolid class="mr-2" />
                Create Your First Form
            </Button>
        </div>
    {:else}
    <!-- List of forms -->
        <div class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-3 gap-6">
            {#each forms as form}
                <Card 
                    form={form} 
                    deleteFormCallback={() => handleDeleteForm(form.id)}
                    editFormCallback={() => handleEditForm(form.id)}
                />
            {/each}
        </div>
    {/if}
</div>
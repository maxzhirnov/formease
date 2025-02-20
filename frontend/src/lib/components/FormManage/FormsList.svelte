<script lang="ts">
    import { Alert, Button, GradientButton, Heading, Skeleton, Spinner } from 'flowbite-svelte';
    import { CirclePlusSolid, RefreshOutline } from 'flowbite-svelte-icons';

    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';

    import { formService } from '$lib/services/formService';
    import { listForms, generateFormWithAI, type AIFormRequest, getForm } from '$lib/api/forms';
    import { showError } from '$lib/utils/errorHandler';
    import type { FormData } from '$lib/types/form';

    import Create from './Create.svelte';
    import Card from './Card.svelte';
    
    let forms: FormData[] = $state([]);
    let isLoading = $state(true);
    let error: string | null = $state(null);

    async function loadForms() {
        try {
            isLoading = true;
            error = null;
            forms = await listForms();
        } catch (err) {
            const errMessage = err instanceof Error ? err.message : 'Failed to load forms';
            error = errMessage;
            showError(errMessage);
        } finally {
            isLoading = false;
        }
    }

    // Add a new function to update a single form
    async function updateSingleForm(formId: string) {
        try {
            const updatedForm = await getForm(formId, {}); // Assuming you have this API endpoint
            forms = forms.map(form => 
                form.id === formId ? updatedForm : form
            );
        } catch (err) {
            const errMessage = err instanceof Error ? err.message : 'Failed to update form';
            error = errMessage;
            showError(errMessage);
        }
    }
    async function handleCreateForm(name: string) {
        formService.state.resetForm();
        formService.state.updateFormName(name);
        await formService.api.create()
        loadForms();
    }

    async function handleCreateWithAI(description: string) {
        try {
            // Convert description to AI form request
            const aiRequest : AIFormRequest = {
                topic: description,
                formType: 'custom',  // You might want to detect this from the description
                numQuestions: 5,     // Default value, could be extracted from description
                preferences: []      // Could be extracted from description using AI
            };

            const generatedForm = await generateFormWithAI(aiRequest);

            // Handle success
            // Default behavior: redirect to form editor
            goto(`/app/form/${generatedForm.id}/edit`);
        } catch (error) {
            // Handle generation errors
            const errorMessage =  error instanceof Error ? error.message : 'Failed to generate form with AI'
            showError(errorMessage);
        }
    }
    
    async function handleDeleteForm(id: string) {
        try {
            error = null;
            await formService.api.delete(id)
            forms = forms.filter(form => form.id !== id);
        } catch (err) {
            error = err instanceof Error ? err.message : 'Failed to delete form';
            showError(error);
        } 
    }

    async function handleToggleDraft(id: string) {
        try {
            error = null;
            await formService.api.toggleDraft(id);
        } catch (err) {
            error = err instanceof Error ? err.message : 'Failed to delete form';
            showError(error);
        } finally {
            updateSingleForm(id);
        }
    }

    function handleEditForm(id: string) {
        goto(`/app/form/${id}/edit`);
    }


    onMount(() => {
        loadForms();
    });
    
</script>


    
<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    
    <!-- Title, Cretate and refresh -->
    <div class="flex justify-between items-center mb-2 flex-col sm:flex-row">
        <div class="mb-4 sm:mb-0">
            <Heading 
            tag="h2" 
            class="text-3xl font-extrabold text-gray-900 dark:text-white">
            My Forms
            </Heading>
        </div>
        <div class="flex items-center space-x-2">
                <Create 
                    createFormCallback={handleCreateForm} 
                    createAIFormCallback={handleCreateWithAI}
                 />
    
                <GradientButton 
                    outline
                    size="md"
                    color="purpleToBlue"
                    onclick={loadForms} 
                    disabled={isLoading}>
                    {#if isLoading}
                        <Spinner  size="4" />
                    {:else}
                        <RefreshOutline class="w-4 h-4" />
                    {/if}
                </GradientButton>
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
            {#each Array(4) as _}
                <div class="space-y-4">
                    <Skeleton/>
                    <div class="space-y-2">
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
        <div class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-3 gap-6">
            {#each forms as form (form.id)}
                <Card 
                    form={form} 
                    deleteFormCallback={() => handleDeleteForm(form.id)}
                    editFormCallback={() => handleEditForm(form.id)}
                    toggleDraftCallback={() => handleToggleDraft(form.id)}
                />
            {/each}
        </div>
    {/if}
    
</div>
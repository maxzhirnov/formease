<!-- src/routes/dashboard/form/[id]/edit/+page.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import FormEditor from '$lib/components/FormManage/EditorForm.svelte';
    import { formService } from '$lib/services/formService';
    import { showError } from '$lib/utils/errorHandler';
    import { Spinner } from 'flowbite-svelte';

    let loading = true;
    let error: string | null = null;

    onMount(async () => {
        try {
            let form = await formService.api.fetch($page.params.id)
            formService.state.setForm(form)
        } catch (e) {
            showError(e)
        }
        loading = false;
    });
</script>
    
{#if loading}
    <Spinner/>
    <p>Loading...</p>
{:else}
    <FormEditor />  
{/if}
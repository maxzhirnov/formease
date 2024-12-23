<!-- src/routes/dashboard/form/[id]/edit/+page.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { formStore } from '$lib/stores/formStore';
    import { getForm } from '$lib/api/forms';
    import FormEditor from '$lib/components/FormManage/EditorForm.svelte';

    let loading = true;
    let error: string | null = null;

    onMount(async () => {
        try {
            const form = await getForm($page.params.id, {}, fetch);
            formStore.set(form);
        } catch (e: unknown) {
            // Type guard for Error objects
            if (e instanceof Error) {
                error = e.message;
            } else {
                error = 'An unknown error occurred';
            }
        } finally {
            loading = false;
        }
    });
</script>

{#if loading}
    <div>Loading...</div>
{:else if error}
    <div>Error: {error}</div>
{:else}
    <!-- <pre>
        {JSON.stringify($formStore, null, 2)}
    </pre> -->
    <FormEditor />
{/if}

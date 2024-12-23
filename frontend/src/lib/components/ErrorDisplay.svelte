<script lang="ts">
    import { errorStore } from '$lib/stores/errorStore';
    import { onDestroy } from 'svelte';
    import { Toast } from 'flowbite-svelte';
    import { slide } from 'svelte/transition';
    import { CloseCircleSolid } from 'flowbite-svelte-icons';

    let errorMessage: string | null = null;

    const unsubscribe = errorStore.subscribe((value: string | null) => {
        errorMessage = value;

        // Auto close the error message after 5 seconds
        if (value) {
            setTimeout(() => {
                errorStore.set(null);
            }, 5000);
        }
    });

    onDestroy(() => {
        unsubscribe();
    });
</script>

{#if errorMessage}
    <Toast 
        transition={slide} 
        position="top-right" 
        color="red" 
        class="fixed top-4 right-4 z-50"
    >
        <svelte:fragment slot="icon">
            <CloseCircleSolid class="w-5 h-5" />
            <span class="sr-only">Error icon</span>
         </svelte:fragment>
        <span class="font-medium">{errorMessage}</span>

    </Toast>
{/if}

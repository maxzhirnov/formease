<script lang="ts">
    import '../app.css';
    import ErrorDisplay from '$lib/components/ErrorDisplay.svelte';
    import type { Snippet } from 'svelte';

    let { children }: {children: Snippet} = $props()
</script>

<main class="max-w-7xl mx-auto">
    <ErrorDisplay/>
    {@render children()}
</main>


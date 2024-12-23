<script lang="ts">
    import '../app.css';
    import ErrorDisplay from '$lib/components/ErrorDisplay.svelte';
    import type { Snippet } from 'svelte';

    let { children }: {children: Snippet} = $props()
</script>

<ErrorDisplay/>
{@render children()}


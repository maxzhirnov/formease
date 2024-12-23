<script lang="ts">
    import type { PageData } from './$types';
    import Form from "$lib/components/FormShow/Form.svelte";
    import { Spinner } from 'flowbite-svelte'
    import { randomBrightColor } from '$lib/theme/utils';
    
    export let data: PageData;

    let bgColor = randomBrightColor();
</script>
<div class={`min-h-screen ${bgColor}`}>

{#if data.form?.questions}
    <Form 
        questions={data.form.questions}
        thankYouMessage={data.form.thankYouMessage || {}}
        theme={data.form.theme}
        floatingShapesTheme={data.form.floatingShapesTheme}
    />
{:else}
    <div class="min-h-screen flex items-center justify-center">
        <Spinner size="16" color="blue" />
    </div>
{/if}
</div>
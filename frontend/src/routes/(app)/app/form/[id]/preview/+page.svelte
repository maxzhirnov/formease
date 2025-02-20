<script lang="ts">
    import Form from "$lib/components/FormShow/Form.svelte";
    import { randomBrightColor } from '$lib/theme/utils';
    import Preview from '$lib/components/FormShow/Preview.svelte';
    
    export let data;
    let bgColor = randomBrightColor();
    
</script>

<div class={`min-h-screen ${bgColor}`}>
    <Preview/>
    {#if data.form}
        <Form 
            questions={data.form.questions}
            thankYouMessage={data.form.thankYouMessage}
            theme={data.form.theme}
            floatingShapesTheme={data.form.floatingShapesTheme}
        />
    {/if}
</div>

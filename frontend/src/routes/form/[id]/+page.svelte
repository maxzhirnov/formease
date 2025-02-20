<script lang="ts">
    import type { PageData } from './$types';
    import Form from "$lib/components/FormShow/Form.svelte";
    import { Button, Spinner } from 'flowbite-svelte';
    import { randomBrightColor } from '$lib/theme/utils';
    import { browser } from '$app/environment';
    import AlreadySubmitted from '$lib/components/FormShow/AlreadySubmitted.svelte';
    import { page } from '$app/stores';
    import Preview from '$lib/components/FormShow/Preview.svelte';
    import Unpublished from '$lib/components/FormShow/Unpublished.svelte';
    
    export let data: PageData;
    let bgColor = randomBrightColor();
    let isFormSubmitted = false;
    let isChecking = true; // Add loading state

    let isPreview = $page.url.searchParams.get('preview')

    // Move check outside onMount
    if (browser && data.form?.id) {
        const submittedForms = localStorage.getItem('submittedForms');
        if (submittedForms) {
            const forms = JSON.parse(submittedForms);
            isFormSubmitted = forms.includes(data.form.id);
        }
        isChecking = false;
    }

    function saveFormSubmission() {
        if (!data.form?.id) return;
        
        const submittedForms = localStorage.getItem('submittedForms');
        let forms = submittedForms ? JSON.parse(submittedForms) : [];
        forms.push(data.form.id);
        localStorage.setItem('submittedForms', JSON.stringify(forms));
    }


    
</script>

{#if !data.form?.isDraft}
    <div class={`min-h-screen ${bgColor}`}>
        {#if isChecking}
            <div class="min-h-screen flex items-center justify-center">
                <Spinner size="16" color="purple" />
            </div>
        {:else if data.form?.questions}
            {#if isFormSubmitted}
                <AlreadySubmitted/>
            {:else}
                {#if isPreview === "true"}
                    <Preview/>
                {/if}
                <Form 
                    questions={data.form.questions}
                    thankYouMessage={data.form.thankYouMessage}
                    theme={data.form.theme}
                    floatingShapesTheme={data.form.floatingShapesTheme}
                    onSubmit={saveFormSubmission}
                />
            {/if}
        {:else}
            <div class="min-h-screen flex items-center justify-center">
                <Spinner size="16" color="purple" />
            </div>
        {/if}
    </div>
{:else}
    <Unpublished />
{/if}
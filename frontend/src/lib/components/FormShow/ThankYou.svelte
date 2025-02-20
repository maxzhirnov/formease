<script lang="ts">
    import { goto } from "$app/navigation";
    import { fade, fly, slide } from "svelte/transition";
    import { Button } from 'flowbite-svelte';
    import type { ThankYouMessage } from "$lib/types/thank";

    export let thankYouMessage : ThankYouMessage;

    const handleButtonClick = () => {
        if (thankYouMessage.button) {
            if (thankYouMessage.button.newTab) {
                window.open(thankYouMessage.button.url, "_blank");
            } else {
                goto(thankYouMessage.button.url as string);
            }
        }
    };
</script>

<div 
    class="text-center px-8 py-16 space-y-8" 
    in:fade={{ duration: 400 }}
>
    <h2 
        class="text-4xl font-bold" 
        in:slide={{ duration: 400, delay: 200 }}
    >
        {thankYouMessage.title}
    </h2>
    
    <p 
        class="text-xl " 
        in:fade={{ duration: 400, delay: 300 }}
    >
        {thankYouMessage.subtitle}
    </p>
    
    <div 
        class="text-6xl my-8 drop-shadow-md" 
        in:fly={{ y: 20, duration: 500, delay: 400 }}
    >
        {thankYouMessage.icon}
    </div>
    
    {#if thankYouMessage.button}
        <Button 
            on:click={handleButtonClick}
            class="bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600"
        >
            {thankYouMessage.button.text}
        </Button>
    {/if}
</div>

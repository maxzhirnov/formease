<script lang="ts">
    import { goto } from "$app/navigation";
    import type { ThankYouMessage } from "$lib/types";
    import { fade, fly, slide } from "svelte/transition";
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

<div class="thank-you-container" in:fade={{ duration: 400 }}>
    <h2 in:slide={{ duration: 400, delay: 200 }}>
        {thankYouMessage.title}
    </h2>
    <p class="subtext" in:fade={{ duration: 400, delay: 300 }}>
        {thankYouMessage.subtitle}
    </p>
    <div class="checkmark" in:fly={{ y: 20, duration: 500, delay: 400 }}>
        {thankYouMessage.icon}
    </div>
    {#if thankYouMessage.button}
        <button in:fly={{ y: 20, duration: 500, delay: 600 }} on:click={handleButtonClick}>
            {thankYouMessage.button.text}
        </button>
    {/if}
</div>

<style>
    .thank-you-container {
        text-align: center;
        padding: 4rem 2rem;
    }

    .checkmark {
        font-size: 4rem;
        margin: 2rem 0;
        filter: drop-shadow(0 2px 4px rgba(0,0,0,0.2));
    }
    button {
        padding: 1rem 2rem;
        background: linear-gradient(to right, var(--gradient-colors));
        border: none;
        border-radius: 12px;
        color: white;
        font-size: 1.2rem;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.3s ease;
        box-shadow: 0 4px 12px rgba(0,0,0,0.2);
    }

    button:hover {
        transform: translateY(-2px);
        box-shadow: 0 6px 16px rgba(0,0,0,0.3);
    }
</style>
<script lang="ts">
    import type { ThankYouMessage } from "$lib/types/thank";


    import { Label, Input, Checkbox, Button, P } from "flowbite-svelte";

    interface Props {
        thankYouScreen: ThankYouMessage;
        handleSaveCallback: () => Promise<void>;
    }

    let { thankYouScreen, handleSaveCallback }: Props = $props();
    

</script>
<div class="space-y-6 p-6 mb-6 bg-white/10 backdrop-blur-lg rounded-2xl border border-white/10">
    <h2 class="text-2xl font-bold mb-6 text-gray-800">Thank You Screen</h2>

    <div class="grid md:grid-cols-2 gap-6">
        <Label>
            Title
            <Input 
                type="text" 
                placeholder="Enter thank you title" 
                bind:value={thankYouScreen.title}
            />
        </Label>

        <Label>
            Subtitle
            <Input 
                type="text" 
                placeholder="Enter thank you subtitle" 
                bind:value={thankYouScreen.subtitle}
            />
        </Label>

        <Label>
            Icon
            <div class="flex items-center gap-2">
                <Input 
                    type="text" 
                    placeholder="Enter emoji or icon" 
                    bind:value={thankYouScreen.icon}
                    class="flex-grow"
                />
                <span class="text-3xl">{thankYouScreen.icon}</span>
            </div>
        </Label>
    </div>

    <div class="space-y-4">
        <h3 class="text-xl font-semibold">Button Configuration</h3>
        <div class="grid md:grid-cols-2 gap-6">
            <Label>
                Button Text
                <Input 
                    type="text" 
                    placeholder="Enter button text" 
                    bind:value={thankYouScreen.button.text} 
                />
            </Label>

            <Label>
                Button URL
                <Input 
                    type="url" 
                    placeholder="Enter redirect URL" 
                    bind:value={thankYouScreen.button.url}
                />
            </Label>

            <Label class="flex items-center gap-2">
                <Checkbox 
                    bind:checked={thankYouScreen.button.newTab}
                />
                Open in New Tab
            </Label>
        </div>
    </div>

    <div class="mt-6 bg-gray-50 dark:bg-gray-800 p-4 rounded-lg">
        <h3 class="text-lg font-semibold mb-4">Preview</h3>
        <div class="flex flex-col items-center text-center space-y-4">
            <div class="text-6xl mb-4">{thankYouScreen.icon}</div>
            <h2 class="text-2xl font-bold">{thankYouScreen.title}</h2>
            <p class="text-gray-600 dark:text-gray-300">{thankYouScreen.subtitle}</p>
            <Button 
                color="blue" 
                href={thankYouScreen.button.url} 
                target={thankYouScreen.button.newTab ? '_blank' : undefined}
            >
                {thankYouScreen.button.text}
            </Button>
        </div>
    </div>
    
    <div class="text-center mt-6">
        <Button color="green" on:click={handleSaveCallback}>Save</Button>
    </div>
    
</div>
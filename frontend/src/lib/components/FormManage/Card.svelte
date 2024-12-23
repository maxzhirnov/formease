<script lang="ts">
    import { Badge, Button, ButtonGroup, Heading, Card, P } from 'flowbite-svelte';
    import { FilePenSolid, TrashBinSolid } from 'flowbite-svelte-icons'
    import { goto } from '$app/navigation';
    import type { FormData } from '$lib/types';

    interface Props {
        form: FormData;
        deleteFormCallback: (id: string) => void;
        editFormCallback: (id: string) => void;
    }

    let { form, deleteFormCallback, editFormCallback }: Props = $props();
    let isDeleting = $state(false);
</script>

<Card 
    class="
        relative 
        w-auto
        hover:scale-[1.01] 
        transition-transform 
        duration-300 
        bg-white/10 
        backdrop-blur-lg border border-white/10 
        shadow-xl
        " 
    padding="md"
    size="md"
>
    <div class="flex justify-between items-center mb-4">
        <Badge color="dark" class="text-xs">{form.id}</Badge>
        <Badge color={form.isDraft ? 'yellow' : 'green'}>
            {form.isDraft ? 'Draft' : 'Published'}
        </Badge>
    </div>

    <div class="space-y-3">
        <Heading 
            tag="h5" 
            class="text-xl font-bold text-gray-900 dark:text-white truncate"
        >
            {form.name}
        </Heading>
        
        <P class="text-gray-500 text-sm flex items-center gap-2">
            <span class="text-lg">📝</span>
            {#if form.questions}
                {form.questions.length} Questions
            {:else}
                No Questions
            {/if}
        </P>
    </div>

    <ButtonGroup class="mt-6 grid grid-cols-3">
        <Button 
            pill 
            color="light" 
            class="col-span-1" 
            onclick={() => goto(`/form/${form.id}`)}
        >
            Preview
        </Button>
        <Button 
            pill 
            color="light" 
            class="col-span-1 flex items-center justify-center" 
            onclick={() => editFormCallback(form.id)}
        >
            <FilePenSolid class="mr-1 w-4 h-4" />Edit
        </Button>
        <Button 
            pill 
            outline 
            color="red" 
            class="col-span-1" 
            onclick={() => deleteFormCallback(form.id)} 
            disabled={isDeleting}
        >
            <TrashBinSolid class="w-4 h-4" />
        </Button>
    </ButtonGroup>
</Card>

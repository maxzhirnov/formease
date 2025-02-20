<script lang="ts">
    import { Badge, Button, ButtonGroup, Heading, Card, P, Modal, Tooltip, Popover } from 'flowbite-svelte';
    import { EditOutline, TrashBinOutline, EyeOutline, UploadOutline, DownloadOutline } from 'flowbite-svelte-icons'
    import { goto } from '$app/navigation';
    import type { FormData } from '$lib/types/form';
    import ShareUrlButton from './ShareURLButton.svelte';
    import { PUBLIC_URL } from '$env/static/public';
    import { page } from '$app/stores';

    interface Props {
        form: FormData;
        deleteFormCallback: (id: string) => void;
        editFormCallback: (id: string) => void;
        toggleDraftCallback: (id: string) => void;
    }

    let { form, deleteFormCallback, editFormCallback, toggleDraftCallback }: Props = $props();
    let isDeleting = $state(false);
    let showDeleteModal = $state(false);
    let hasQuestions = $derived((form.questions?.length ?? 0) > 0);
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
        <Badge color="dark" class="text-xs truncate">{form.id}</Badge>
        <Badge color={form.isDraft ? 'yellow' : 'green'}>
            {form.isDraft ? 'Draft' : 'Published'}
        </Badge>
    </div>

    <div class="">
        <Heading 
            tag="h5"
            customSize="text-lg font-semibold"
            class="text-gray-700 truncate"
        >
            {form.name}
        </Heading>

        <P class="text-gray-500 text-xs flex items-center gap-2">
            <span class="text-lg">📝</span>
            {#if form.questions}
                {form.questions.length} Questions
            {:else}
                No Questions
            {/if}
        </P>
    </div>

    {#if !form.isDraft}
    <ShareUrlButton 
        url={`http://${$page.url.host}/form/${form.id}`} 
        class="ml-2 text-gray-500 hover:text-gray-700 w-10 h-10"
        />
    {:else}
    <!-- spacer -->
        <div class="ml-2 text-gray-500 hover:text-gray-700 w-10 h-10"></div>
    {/if}

    <ButtonGroup class="mt-1 grid grid-cols-4">
            <Button 
                disabled={!hasQuestions}
                color="light" 
                class="col-span-1" 
                on:click={() => goto(`/app/form/${form.id}/preview`)}
            >
                <div class="flex flex-col xl:flex-row items-center justify-center">
                    <div><EyeOutline class="w-4 h-4" /></div>
                    <div class="ml-1 text-xs hidden md:inline">Preview</div>
                </div>
            </Button>
            <Tooltip>Preview form</Tooltip>
    
            {#if form.isDraft}
                <Button 
                    color="light" 
                    class="col-span-1 flex items-center justify-center" 
                    on:click={() => toggleDraftCallback(form.id)}
                >   
                    <div class="flex flex-col xl:flex-row items-center justify-center">
                        <div><UploadOutline class="w-4 h-4" /></div>
                        <div class="ml-1 text-xs hidden md:inline">Publish</div>
                    </div>
                </Button>
                <Tooltip>Publish form</Tooltip>
            {:else}
                <Button 
                    color="light" 
                    class="col-span-1 flex items-center justify-center" 
                    on:click={() => toggleDraftCallback(form.id)}
                >
                    <div class="flex flex-col xl:flex-row items-center justify-center">
                        <div><DownloadOutline class="w-4 h-4" /></div>
                        <div class="ml-1 text-xs hidden md:inline">Unpublish</div>
                    </div>
                </Button>
                <Tooltip>Move to draft</Tooltip>
            {/if}
    
            <Button 
                color="light" 
                class="col-span-1 flex items-center justify-center" 
                on:click={() => editFormCallback(form.id)}
            >
                <div class="flex flex-col xl:flex-row items-center justify-center">
                    <div><EditOutline class="w-4 h-4" /></div>
                    <div class="ml-1 text-xs hidden md:inline">Edit</div>
                </div>
            </Button>
            <Tooltip>Edit form</Tooltip>
    
            <Button 
                outline 
                color="red" 
                class="col-span-1" 
                on:click={() => showDeleteModal = true} 
                disabled={isDeleting}
            >
                <div class="flex flex-col xl:flex-row items-center justify-center">
                    <div><TrashBinOutline class="w-4 h-4" /></div>
                    <div class="ml-1 text-xs hidden md:inline">Delete</div>
                </div>
            </Button>
            <Tooltip>Delete form</Tooltip>
    </ButtonGroup>
</Card>

<Modal title="" bind:open={showDeleteModal} autoclose size="sm" class="w-full">
    <svg class="text-gray-400 dark:text-gray-500 w-11 h-11 mb-3.5 mx-auto" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" /></svg>
    <p class="mb-4 text-gray-500 dark:text-gray-300 text-center">Are you sure you want to delete this item?</p>
    <div class="flex justify-center items-center space-x-4">
      <Button color="light" on:click={() => showDeleteModal = false}>No, cancel</Button>
      <Button color="red" on:click={() => deleteFormCallback(form.id)}>Yes, I'm sure</Button>
    </div>
</Modal>
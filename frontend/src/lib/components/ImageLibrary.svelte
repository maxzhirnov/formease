<script lang="ts">
    import { Button, Alert } from 'flowbite-svelte';
    import { TrashBinSolid, CheckCircleSolid } from 'flowbite-svelte-icons';
    import { getImages, deleteImage } from '$lib/api/forms';
    import { showError } from '$lib/utils/errorHandler';
    import ImageUploader from './ImageUploader.svelte';
    import { onMount } from 'svelte';

    interface Props {
        onImageChosen: (imageUrl: string) => void;
    }

    let { onImageChosen }: Props = $props();
    
    let images: any[] = $state([]);
    let loading = $state(true);
    let page = $state(1);
    let totalImages = $state(0);
    let limit = $state(20);
    let deleting = $state(false);

    let totalPages = $derived(Math.ceil(totalImages / limit));
    
    async function loadImages(pageNum = 1) {
        try {
            loading = true;
            const response = await getImages(pageNum, limit);
            images = response.images;
            totalImages = response.total;
            page = response.page;
        } catch (error) {
            console.error('Error loading images:', error);
            showError('Failed to load images');
        } finally {
            loading = false;
        }
    }

    async function handleDeleteImage(imageId: string) {
        try {
            deleting = true;
            await deleteImage(imageId);
            
            // Reload current page after deletion
            await loadImages(page);
            
            // Clear any selected images
        } catch (error) {
            console.error('Error deleting image:', error);
            showError('Failed to delete image');
        } finally {
            deleting = false;
        }
    }

    function handleImageChosen(image: any) {
        onImageChosen(image.url);
        return;
    }

    function handleRemoveImage() {
        onImageChosen("");
        return;
    }

    async function open() {
        await loadImages();
    }


    async function handleUploadSuccess(imageUrl: string) {
        await loadImages();
    }

    onMount(async () => {
        await loadImages();
    })
</script>

{#if loading}
    <div class="flex justify-center p-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
    </div>
    {:else}
    <Button 
        color="red"
        onclick={handleRemoveImage}>
        Remove
    </Button>
    {#if !images || images.length === 0}
        <div class="space-y-4">
            <div 
                class="relative aspect-square cursor-pointer group flex items-center justify-center hover:bg-gray-50 transition-colors max-w-xs"
            >
                <ImageUploader 
                    onUploadSuccessCallback={handleUploadSuccess}
                    className="w-full h-full flex items-center justify-center"
                />
            </div>
            <Alert color="yellow">
                No images found. Start uploading!
            </Alert>
        </div>
    {:else}
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div 
                class="relative aspect-square cursor-pointer group flex items-center justify-center hover:bg-gray-50 transition-colors"
            >
                <ImageUploader 
                    onUploadSuccessCallback={handleUploadSuccess}
                    className="w-full h-full flex items-center justify-center"
                />
            </div>
            {#each images as image}
                <!-- svelte-ignore a11y_no_static_element_interactions -->
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                <div 
                class="relative aspect-square  group"
                >
                <img 
                    src={`${image.url}`}
                    alt={image.title || 'Uploaded image'}
                    class="w-full h-full object-cover rounded-lg"
                />
                <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-20 transition-all rounded-lg">

                        <div class="flex gap-2 absolute bottom-2 right-2 opacity-100 md:opacity-50 group-hover:opacity-100 transition-opacity">
                        <Button 
                            color="green"
                            class="w-10 h-10"
                            onclick={() => handleImageChosen(image)}    
                            >
                            <CheckCircleSolid  />
                        </Button>  
                        <Button 
                            color="red" 
                            class="w-10 h-10"
                            disabled={deleting}
                            onclick={() => handleDeleteImage(image.id)}
                        >
                            <TrashBinSolid  />
                        </Button>
                        </div>
                </div>
                </div>
            {/each}
        </div>

        <!-- Pagination -->
        {#if totalImages > limit}
            <div class="flex justify-center space-x-2 mt-4">
                <Button 
                    size="sm"
                    disabled={page === 1}
                    on:click={() => loadImages(page - 1)}
                >
                    Previous
                </Button>
                
                <span class="px-4 py-2">
                    Page {page} of {totalPages}
                </span>
                
                <Button 
                    size="sm"
                    disabled={page === totalPages}
                    on:click={() => loadImages(page + 1)}
                >
                    Next
                </Button>
            </div>
        {/if}
    {/if}
{/if}

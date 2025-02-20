<script lang="ts">
    import { uploadImage } from '$lib/api/forms';
    import { showError } from '$lib/utils/errorHandler';
    import { Dropzone } from 'flowbite-svelte';

    interface Props {
        className?: string;
        onUploadSuccessCallback: (imageUrl: string) => void;
    }

    let { className, onUploadSuccessCallback }: Props = $props();

    let fileName = $state('');
    let uploading = $state(false);

    async function processFile(file: File) {
        if (!file || !file.type.startsWith('image/')) {
            showError('Please upload a valid image file');
            return;
        }
        
        try {
            uploading = true;
            const imageUrl = await uploadImage(file);
            onUploadSuccessCallback(imageUrl);
            fileName = file.name;
        } catch (error) {
            if (error instanceof Error) {
                showError(error.message);
            } else {
                showError('An unknown error occurred');
            }
        } finally {
            uploading = false;
        }
    }

    function handleDrop(event: DragEvent) {
        event.preventDefault();
        const files = event.dataTransfer?.files;
        
        if (files && files.length > 0) {
            processFile(files[0]);
        }
    }

    function handleChange(event: Event) {
        const input = event.target as HTMLInputElement;
        const files = input.files;
        
        if (files && files.length > 0) {
            processFile(files[0]);
        }
    }

    function truncateFileName(name: string, maxLength = 40): string {
        return name.length > maxLength 
            ? name.slice(0, maxLength) + '...' 
            : name;
    }
</script>

<div class={`${className} relative w-full h-full`}>
    <Dropzone
        id="dropzone"
        class="absolute inset-0 flex flex-col items-center justify-center h-full w-full"
        on:drop={handleDrop}
        on:dragover={(event) => event.preventDefault()}
        on:change={handleChange}
    >
        <svg 
            aria-hidden="true" 
            class="mb-3 w-10 h-10 text-gray-400" 
            fill="none" 
            stroke="currentColor" 
            viewBox="0 0 24 24" 
            xmlns="http://www.w3.org/2000/svg"
        >
            <path 
                stroke-linecap="round" 
                stroke-linejoin="round" 
                stroke-width="2" 
                d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" 
            />
        </svg>

        <div class="text-center px-2">
            {#if !fileName}
                <p class="mb-2 text-sm text-gray-500 dark:text-gray-400">
                    <span class="font-semibold">Click to upload</span> <span class="hidden md:inline">or drag and drop</span>
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400">
                    SVG, PNG, JPG or GIF
                </p>
            {:else}
                <p>{truncateFileName(fileName)}</p>
            {/if}
        </div>
    </Dropzone>

    {#if uploading}
        <div class="absolute inset-0 bg-white/50 flex flex-col items-center justify-center">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
            <p class="mt-2 text-sm text-gray-600">Uploading...</p>
        </div>
    {/if}
</div>

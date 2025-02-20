<script lang="ts">

    interface Props {
        url: string;
        buttonText?: string;
        class?: string;
    }

    let { url, buttonText  = '', class: className = '' }: Props = $props();
    let isModalOpen = $state(false);
  
    function copyToClipboard(): Promise<void> {
      return navigator.clipboard.writeText(url);
    }
  
    function openInNewTab(): void {
      window.open(url, '_blank');
    }
</script>
  
  <button
    onclick={() => isModalOpen = true}
    class={`inline-flex items-center p-2 rounded-lg hover:bg-gray-100 transition-colors duration-200 ${className}`}
  >
    <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
    </svg>
    {buttonText}
  </button>
  
  {#if isModalOpen}
    <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 rounded-md px-2">
      <div class="bg-white rounded-lg p-6 w-full max-w-md shadow-xl">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold">Share Link</h3>
          <!-- svelte-ignore a11y_consider_explicit_label -->
          <button
            onclick={() => isModalOpen = false}
            class="text-gray-500 hover:text-gray-700"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
          </button>
        </div>
        
        <div class="flex gap-1">
          <input
            type="text"
            readonly
            value={url}
            class="flex-1 px-3 py-2 border rounded-lg bg-gray-50 text-xs"
          />
          
          <!-- svelte-ignore a11y_consider_explicit_label -->
          <button
            onclick={copyToClipboard}
            class="px-4 py-2 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors duration-200"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
            </svg>
          </button>
          
          <!-- svelte-ignore a11y_consider_explicit_label -->
          <button
            onclick={openInNewTab}
            class="px-4 py-2 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors duration-200"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  {/if}
  
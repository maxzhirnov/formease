<script>
    import { goto } from "$app/navigation";


    import { Button, Spinner } from "flowbite-svelte";

</script>
<div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-purple-50 to-white">
    <div class="p-8 backdrop-blur-lg bg-white/80 rounded-2xl shadow-xl border border-purple-100 flex flex-col items-center gap-4">
        <!-- Icon Container with Purple Glow -->
        <div class="relative mb-4">
            <div class="absolute inset-0 bg-purple-200 rounded-full blur-xl opacity-50"></div>
            <Spinner size="16" color="purple" class="w-12 h-12 relative" />
        </div>
        
        <!-- Messages with Better Typography -->
        <h2 class="text-2xl font-medium text-gray-800">
            Form is not published
        </h2>
        
        <p class="text-center text-gray-600 max-w-sm">
            You can only view unpublished forms in preview mode
        </p>
        
        <!-- Styled Button -->
        <Button 
            color="purple" 
            class="mt-6 px-6 py-2 shadow-lg shadow-purple-100 hover:shadow-purple-200 transition-all duration-300"
            onclick={() => goto('/app')}
        >
            Return to Dashboard
        </Button>
    </div>
</div>

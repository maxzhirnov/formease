<script lang="ts">
    import { browser } from "$app/environment";
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { Button } from "flowbite-svelte";
    import { ArrowLeftOutline, PenSolid, EyeSolid } from 'flowbite-svelte-icons'; // 

    function handleBack(): void {
        if (browser) {
            window.history.back();
        }
    }

    function handleEdit(): void {
        goto(`/app/form/${$page.params.id}/edit`);
    }
</script>

<style lang="postcss">
    .preview-stripes {
        background-image: linear-gradient(
            135deg,
            rgba(255, 255, 255, 0.05) 25%,
            transparent 25%,
            transparent 50%,
            rgba(255, 255, 255, 0.05) 50%,
            rgba(255, 255, 255, 0.05) 75%,
            transparent 75%
        );
        background-size: 16px 16px;
        animation: moveStripes 30s linear infinite;
    }

    @keyframes moveStripes {
        0% {
            background-position: 0 0;
        }
        100% {
            background-position: 50px 50px;
        }
    }

    .glow-effect {
        animation: glow 2s ease-in-out infinite alternate;
    }

    @keyframes glow {
        from {
            box-shadow: 0 0 5px rgba(167, 139, 250, 0.3),
                        0 0 10px rgba(167, 139, 250, 0.3);
        }
        to {
            box-shadow: 0 0 10px rgba(167, 139, 250, 0.6),
                        0 0 20px rgba(167, 139, 250, 0.6);
        }
    }
</style>

<div class="fixed bottom-0 left-0 right-0 z-50">
    <div class="preview-stripes bg-gradient-to-r from-purple-900/90 to-indigo-900/90 backdrop-blur-sm shadow-lg relative overflow-hidden">
        <!-- Animated dots - Add negative z-index -->
        <div class="absolute inset-0 opacity-10 -z-10">
            <div class="absolute h-1 w-1 bg-white rounded-full top-1 left-[10%] animate-ping"></div>
            <div class="absolute h-1 w-1 bg-white rounded-full top-8 left-[20%] animate-ping [animation-delay:0.5s]"></div>
            <div class="absolute h-1 w-1 bg-white rounded-full top-2 left-[80%] animate-ping [animation-delay:1s]"></div>
        </div>

        <div class="container mx-auto px-4 relative z-10">
            <div class="flex items-center justify-between h-12">
                <!-- Left section -->
                <div class="flex items-center space-x-3">
                    <Button 
                        class="!p-2 !bg-transparent hover:!bg-white/10 transition-colors relative group"
                        on:click={handleBack}
                    >
                        <ArrowLeftOutline class="w-4 h-4 group-hover:scale-110 transition-transform" />
                        <div class="absolute -bottom-1 left-1/2 w-0 h-0.5 bg-white transform -translate-x-1/2 group-hover:w-full transition-all"></div>
                    </Button>
                    
                    <!-- Preview badge -->
                    <div class="flex items-center space-x-2 bg-white/10 rounded-full px-3 py-1 glow-effect">
                        <EyeSolid class="w-4 h-4 text-purple-200 animate-pulse" />
                        <span class="text-sm font-medium text-purple-100">Preview Mode</span>
                    </div>
                </div>

                <!-- Right section -->
                <div class="flex items-center">
                    <Button 
                        class="flex items-center space-x-2 !bg-white/10 hover:!bg-white/20 transition-all hover:scale-105 hover:shadow-lg hover:shadow-purple-500/20"
                        on:click={handleEdit}
                    >
                        <PenSolid class="w-4 h-4" />
                        <span>Edit</span>
                    </Button>
                </div>
            </div>
        </div>

        <!-- Decorative corner elements -->
        <div class="absolute top-0 left-0 w-2 h-2 border-l-2 border-t-2 border-purple-300/30"></div>
        <div class="absolute top-0 right-0 w-2 h-2 border-r-2 border-t-2 border-purple-300/30"></div>
        <div class="absolute bottom-0 left-0 w-2 h-2 border-l-2 border-b-2 border-purple-300/30"></div>
        <div class="absolute bottom-0 right-0 w-2 h-2 border-r-2 border-b-2 border-purple-300/30"></div>
    </div>
</div>


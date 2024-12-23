<script>
    import { Card } from "flowbite-svelte";
</script>

<div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-3">
    {#each ['Active Forms', 'Total Responses', 'Forms Created'] as label}
        <Card 
            class="hover:scale-105 transition-transform duration-300 
                   shadow-md hover:shadow-xl bg-white border border-gray-200"
        >
            <div class="text-center">
                <div class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-purple-500 mb-2">
                    0
                </div>
                <div class="text-sm text-gray-500">
                    {label}
                </div>
            </div>
        </Card>
    {/each}
</div>
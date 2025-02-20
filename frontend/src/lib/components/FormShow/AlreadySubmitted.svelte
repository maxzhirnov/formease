<script lang="ts">
    import { Button, Card } from 'flowbite-svelte';
    import { HomeSolid, CheckCircleSolid } from 'flowbite-svelte-icons';
    import { goto } from '$app/navigation';
</script>

<div class="min-h-screen flex items-center justify-center p-4">
    <Card class="max-w-lg w-full text-center" padding="xl">
        <div class="flex flex-col items-center gap-6">
            <!-- Success Icon -->
            <div class="w-20 h-20 rounded-full bg-purple-100 flex items-center justify-center">
                <CheckCircleSolid class="w-12 h-12 text-purple-500" />
            </div>

            <!-- Message -->
            <div class="space-y-3">
                <h2 class="text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-blue-500">
                    Спасибо за участие!
                </h2>
                <p class="text-lg text-gray-600">
                    Вы уже заполнили эту форму. Ваши ответы были успешно сохранены.
                </p>
            </div>

            <!-- Divider -->
            <div class="w-full max-w-xs mx-auto border-t border-gray-200"></div>

            <!-- Action Button -->
            <Button 
                size="lg"
                color="light"
                class="mt-2"
                on:click={() => goto('/app')}
            >
                <HomeSolid class="w-4 h-4 mr-2" />
                Вернуться на главную
            </Button>
        </div>
    </Card>
</div>

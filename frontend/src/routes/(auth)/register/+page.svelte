<script lang="ts">
    import { register } from '$lib/api/auth';
    import { goto } from '$app/navigation';
    import { Button, Input, Label, Alert } from 'flowbite-svelte';

    let email = '';
    let password = '';
    let error = '';
    let isLoading = false;

    async function handleSubmit() {
        try {
            isLoading = true;
            await register(email, password);
            goto('/login');
        } catch (err: unknown) {
            if (err instanceof Error) {
                error = err.message;
            } else {
                error = 'An unknown error occurred';
            }
        } finally {
            isLoading = false;
        }
    }
</script>

<div class="relative min-h-screen flex items-center justify-center p-4 overflow-hidden">
    <div 
        class="absolute top-[-20%] right-[-10%] w-[500px] h-[500px] 
               rounded-full bg-gradient-to-r from-blue-500/10 to-purple-500/10 
               blur-[80px] opacity-50 pointer-events-none"
    ></div>
    <div 
        class="absolute bottom-[-30%] left-[-15%] w-[500px] h-[500px] 
               rounded-full bg-gradient-to-r from-purple-500/10 to-blue-500/10 
               blur-[80px] opacity-50 pointer-events-none rotate-45"
    ></div>

    <div class="w-full max-w-md bg-white/80 backdrop-blur-lg rounded-2xl p-8 shadow-lg relative z-10">
        <h1 class="text-3xl font-bold text-center mb-6 text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-purple-500">
            Create Account
        </h1>

        <form on:submit|preventDefault={handleSubmit} class="space-y-4">
            <div>
                <Label for="email" class="mb-2">Email</Label>
                <Input 
                    id="email"
                    type="email" 
                    bind:value={email} 
                    placeholder="your@email.com" 
                    required
                    autocomplete="email"
                />
            </div>

            <div>
                <Label for="password" class="mb-2">Password</Label>
                <Input 
                    id="password"
                    type="password" 
                    bind:value={password} 
                    placeholder="••••••••" 
                    required
                    autocomplete="new-password"
                />
            </div>

            <Button 
                type="submit" 
                color="blue" 
                class="w-full bg-gradient-to-r from-blue-500 to-purple-500"
                disabled={isLoading}
            >
                {isLoading ? 'Creating Account...' : 'Create Account'}
            </Button>
        </form>
        
        {#if error}
            <Alert color="red" class="mt-4">
                {error}
            </Alert>
        {/if}
        
        <div class="text-center mt-6 text-sm text-gray-600">
            Already have an account? 
            <a 
                href="/login" 
                class="text-blue-600 hover:underline font-medium"
            >
                Sign In
            </a>
        </div>
    </div>
</div>

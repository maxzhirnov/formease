<script lang="ts">
    import { Button, Input, Label, Alert } from 'flowbite-svelte';

    import { goto } from '$app/navigation';

    import { login } from '$lib/api/auth';

    let email = '';
    let password = '';
    let error = '';
    let isSubmitting = false;

    async function handleSubmit() {
        try {
            await login({ email, password });
            goto('/app');
        } catch (err) {
            error = err instanceof Error ? err.message : 'Login failed';
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
            Welcome Back
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
                    autocomplete="current-password"
                />
            </div>

            <Button 
                type="submit" 
                color="blue" 
                class="w-full bg-gradient-to-r from-blue-500 to-purple-500"
                disabled={isSubmitting}
            >
                {isSubmitting ? 'Logging in...' : 'Login'}
            </Button>
        </form>
        
        {#if error}
            <Alert color="red" class="mt-4">
                {error}
            </Alert>
        {/if}
        
        <div class="text-center mt-6 text-sm text-gray-600">
            Don't have an account? 
            <a 
                href="/register" 
                class="text-blue-600 hover:underline font-medium"
            >
                Register
            </a>
        </div>
    </div>
</div>

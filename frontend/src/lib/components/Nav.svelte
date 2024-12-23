<script lang="ts">
    import { logout } from '$lib/api/auth';
    import { Navbar, NavBrand, Button } from 'flowbite-svelte';

    async function handleLogout() {
        await logout();
    }
</script>

<Navbar 
    class="bg-transparent"
    fluid={true} 
>
    <NavBrand href="/">
        <span class="self-center whitespace-nowrap text-2xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-purple-500">
            FormEase
        </span>
    </NavBrand>
    
    <div class="flex items-center space-x-4">
        <a 
            href="/app" 
            class="text-sm font-medium text-gray-700 hover:text-blue-600 transition-colors"
        >
            Dashboard
        </a>
        <Button color="red" outline 
            on:click={handleLogout}
        >
            Sign Out
        </Button>
    </div>
</Navbar>

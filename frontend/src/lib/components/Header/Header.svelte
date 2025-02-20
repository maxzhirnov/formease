<script>
    import { Navbar, NavBrand, NavHamburger, NavUl, NavLi, Button } from 'flowbite-svelte';
    import { logout } from '$lib/api/auth';
    import Logo from './Logo.svelte';
  

    async function handleLogout() {
        await logout();
    }
  </script>
  
  <header class="flex flex-wrap justify-between items-center mx-auto px-0 md:px-4">
    <Navbar let:hidden let:toggle fluid={true} class="py-4 bg-transparent">
      <div class="flex items-center gap-8">
        
        <Logo />

        <NavUl {hidden} divClass="hidden md:flex" ulClass="flex flex-col font-medium md:flex-row md:mt-0">
          <NavLi class="text-sm" href="/app">My Forms</NavLi>
        </NavUl>
      </div>
      
      <div class="flex items-center">
        <Button class="ml-2" color="red" onclick={handleLogout}>Logout</Button>
        <NavHamburger 
          on:click={toggle} 
          class="inline-flex items-center p-2 ml-1 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" 
        />
      </div>
      
    </Navbar>
</header>


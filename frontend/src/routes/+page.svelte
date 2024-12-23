<script lang="ts">
  import { goto } from '$app/navigation';
  import { Button } from 'flowbite-svelte';

  function handleGetStarted() {
          goto('/login');
  }
</script>

<div class="relative min-h-screen flex items-center justify-center overflow-hidden">
  <div 
      class="absolute top-[-20%] right-[-10%] w-[600px] h-[600px] 
             rounded-full bg-gradient-to-r from-blue-500/10 to-purple-500/10 
             blur-[80px] opacity-50 pointer-events-none"
  ></div>
  <div 
      class="absolute bottom-[-30%] left-[-15%] w-[600px] h-[600px] 
             rounded-full bg-gradient-to-r from-purple-500/10 to-blue-500/10 
             blur-[80px] opacity-50 pointer-events-none rotate-45"
  ></div>
  
  <div class="text-center max-w-2xl px-4 relative z-10">
      <h1 class="text-4xl md:text-5xl font-extrabold mb-6 text-gray-900">
          Welcome to <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-purple-500">
              FormEase
          </span>
      </h1>
      
      <p class="text-xl text-gray-600 mb-10 max-w-xl mx-auto">
          Create beautiful, interactive forms in minutes. Share your vision, collect responses, 
          and gain meaningful insights with our intuitive form builder.
      </p>
      
      <Button 
          size="xl" 
          class="bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600"
          on:click={handleGetStarted}
      >
          Get Started
      </Button>
  </div>
</div>

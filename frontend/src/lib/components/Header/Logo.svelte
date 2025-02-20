<script>

    import { NavBrand } from "flowbite-svelte";

</script>

<NavBrand href="/">
    <svg 
        class="w-12 h-12 mr-0 transition-transform duration-300 ease-out hover:scale-105" 
        viewBox="0 0 32 32" 
        fill="none" 
        xmlns="http://www.w3.org/2000/svg"
    >
        <!-- Abstract Form Shape -->
        <path 
            d="M8 6L24 6C26 6 27 7 27 9L27 23C27 25 26 26 24 26L8 26C6 26 5 25 5 23L5 9C5 7 6 6 8 6Z" 
            class="fill-gradient"
            stroke="url(#logoGradient)" 
            stroke-width="1.5" 
            stroke-linejoin="round"
        >
            <animate 
                attributeName="stroke-dasharray" 
                from="0 100" 
                to="100 100" 
                dur="1s" 
                begin="0s" 
                fill="freeze"
            />
        </path>

        <!-- Minimalistic Check Mark -->
        <path 
            d="M13 16L15 18L19 14" 
            stroke="url(#accentGradient)" 
            stroke-width="1.5" 
            stroke-linecap="round" 
            stroke-linejoin="round"
        >
            <animate 
                attributeName="stroke-dasharray" 
                from="0 100" 
                to="100 100" 
                dur="0.6s" 
                begin="0.4s" 
                fill="freeze"
            />
        </path>

        <!-- Gradient Definitions -->
        <defs>
            <linearGradient id="logoGradient" x1="5" y1="6" x2="27" y2="26" gradientUnits="userSpaceOnUse">
                <stop offset="0%" class="text-indigo-500" stop-color="currentColor" />
                <stop offset="100%" class="text-purple-500" stop-color="currentColor" />
            </linearGradient>
            <linearGradient id="accentGradient" x1="13" y1="14" x2="19" y2="18" gradientUnits="userSpaceOnUse">
                <stop offset="0%" class="text-purple-400" stop-color="currentColor" />
                <stop offset="100%" class="text-indigo-400" stop-color="currentColor" />
            </linearGradient>
        </defs>
    </svg>
    <span class="self-center text-2xl font-medium tracking-tight bg-gradient-to-r from-indigo-500 to-purple-500 bg-clip-text text-transparent transition-colors duration-300">
        formease
    </span>
</NavBrand>

<style>
    .fill-gradient {
        fill: rgba(99, 102, 241, 0.02);
    }

    :global(.dark) .fill-gradient {
        fill: rgba(99, 102, 241, 0.08);
    }
</style>

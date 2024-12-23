<script lang="ts">
  export let currentQuestion: number;
  export let totalQuestions: number;
</script>

<div 
  class="h-1.5 bg-white/10 rounded-full mb-12 md:mb-10 sm:mb-8"
>
  <div 
    class="h-full bg-white rounded-full transition-all duration-500 ease-cubic-bezier" 
    style="
      width: {Math.min(((currentQuestion + 1) / totalQuestions) * 100, 100)}%;
      box-shadow: 0 0 20px rgba(255,255,255,0.3);
    "
  ></div>
</div>

<script lang="ts">
  export let currentQuestion: number;
  export let totalQuestions: number;
  export let color: string = 'bg-black'
  let bgColor = `bg-${color}`
  let opaqueColor = `${bgColor}/10`
</script>

<div 
  class={`h-1.5 ${opaqueColor} rounded-full mb-12 md:mb-10 sm:mb-8`}
>
  <div 
    class={`h-full ${bgColor} rounded-full transition-all duration-500 ease-cubic-bezier`}
    style="
      width: {Math.min(((currentQuestion + 1) / totalQuestions) * 100, 100)}%;
      box-shadow: 0 0 20px rgba(255,255,255,0.3);
    "
  ></div>
</div>

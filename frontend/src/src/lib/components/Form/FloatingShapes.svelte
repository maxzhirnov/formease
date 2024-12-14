<script lang="ts">
    import { themes, type FloatingShapesTheme, type Shape } from '$lib/themes/floatingShapes';
  
    export let theme: FloatingShapesTheme = 'default';
    export let customShapes: Shape[] | undefined = undefined;
    export let blendMode = "screen";
    export let zIndex = -1;
  
    $: shapes = customShapes || themes[theme];
  </script>
  
  <style>
    .floating-shapes {
      position: fixed;
      inset: 0;
      pointer-events: none;
      overflow: hidden;
    }
  
    .shape {
      position: absolute;
      border-radius: 50%;
    }
  
    @keyframes pattern1 {
      0%, 100% { transform: translate(0, 0) rotate(0deg); }
      33% { transform: translate(100px, 50px) rotate(10deg); }
      66% { transform: translate(50px, 100px) rotate(-10deg); }
    }
  
    @keyframes pattern2 {
      0%, 100% { transform: translate(0, 0) rotate(0deg); }
      33% { transform: translate(-100px, -50px) rotate(-10deg); }
      66% { transform: translate(-50px, -100px) rotate(10deg); }
    }
  
    @keyframes pattern3 {
      0%, 100% { transform: translate(0, 0) rotate(0deg); }
      33% { transform: translate(-50px, 100px) rotate(10deg); }
      66% { transform: translate(100px, -50px) rotate(-10deg); }
    }
  </style>
  
  <div class="floating-shapes" style="z-index: {zIndex};">
    {#each shapes as shape, i}
      <div
        class="shape"
        style="
          width: {shape.width}px;
          height: {shape.height}px;
          background: {shape.color};
          top: {shape.position.top};
          left: {shape.position.left};
          right: {shape.position.right};
          bottom: {shape.position.bottom};
          filter: blur({shape.blur}px);
          opacity: {shape.opacity};
          mix-blend-mode: {blendMode};
          animation: {shape.animationPattern} {shape.animationDuration}s infinite ease-in-out;
          animation-delay: {i * -5}s;
        "
      ></div>
    {/each}
  </div>
  
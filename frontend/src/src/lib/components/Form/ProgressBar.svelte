<script lang="ts">
  export let currentQuestion: number;
  export let totalQuestions: number;
  export let gradient: string;
</script>

<div class="progress">
  <div 
    class="progress-bar" 
    style="
      width: {Math.min(((currentQuestion + 1) / totalQuestions) * 100, 100)}%;
      --gradient-colors: {gradient};
    "
  ></div>
</div>
<style>
    .progress {
      height: 6px;
      background: rgba(255,255,255,0.1);
      margin-bottom: 3rem;
      border-radius: 3px;
      overflow: hidden;
    }
  
    .progress-bar {
      height: 100%;
      background-color: #fff;
      /* background: linear-gradient(to right, var(--gradient-colors)); */
      border-radius: 3px;
      transition: width 0.5s cubic-bezier(0.4, 0, 0.2, 1);
      box-shadow: 0 0 20px rgba(255,255,255,0.3);
    }

    @media (max-width: 480px) {
      .progress {
        margin-bottom: 2rem;
      }
    }
</style>
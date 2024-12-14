<script lang="ts">
    import ProgressBar from '$lib/components/Form/ProgressBar.svelte';
    import QuestionContainer from '$lib/components/Form/QuestionContainer.svelte';
    import FloatingShapes from './FloatingShapes.svelte';
    import ThankYou from './ThankYou.svelte';
    import { predefinedThemes, type ThemeName } from '$lib/themes/themes';
    import type { Question, ThankYouMessage, Theme } from '$lib/types';
    import type { FloatingShapesTheme } from '$lib/themes/floatingShapes';
    
    export let questions : Question[];
    export let thankYouMessage : ThankYouMessage
    export let theme: ThemeName | Theme = 'dark';
    export let floatingShapesTheme : FloatingShapesTheme;

    $: activeTheme = typeof theme === 'string' ? predefinedThemes[theme] : theme;
    
    let isFormSubmitted = false;
    let currentQuestion = 0;
    let answers: Array<string | string[]> = [];

    const handleNavigate = (nextId: number | undefined) => {
        if (nextId) {
            currentQuestion = questions.findIndex(q => q.id === nextId);
        } else {
            // No next question defined, submit the form
            handleSubmit();
        }
    };

    const handleSelection = (answer: string) => {
      if (questions[currentQuestion].type === 'multiple-choice') {
          if (!Array.isArray(answers[currentQuestion])) {
              answers[currentQuestion] = [];
          }
          const currentAnswers = answers[currentQuestion] as string[];
          if (currentAnswers.includes(answer)) {
              answers[currentQuestion] = currentAnswers.filter(a => a !== answer);
          } else {
              answers[currentQuestion] = [...currentAnswers, answer];
          }
      } else {
          answers[currentQuestion] = answer;
          const nextId = questions[currentQuestion].nextQuestion?.conditions.find(c => c.answer === answer)?.nextId 
              || questions[currentQuestion].nextQuestion?.default;
          handleNavigate(nextId); // Pass undefined if no next question
      }
  };

    const handleSubmit = async () => {
        console.log('Form Answers:', answers);
        isFormSubmitted = true;
        // submit logic here
    };

    const isLastQuestion = (questionId: number): boolean => {
        const question = questions.find(q => q.id === questionId);
        return !question?.nextQuestion;
    };
  </script>
  
  <div class="container" style="background-color: {activeTheme.backgroundColor}; color: {activeTheme.fontColor};">

    <FloatingShapes theme={floatingShapesTheme} zIndex={0} />
    
    <div class="questionnaire">
      {#if !isFormSubmitted}
        <ProgressBar 
            currentQuestion={Math.min(currentQuestion, questions.length - 1)}
            totalQuestions={questions.length}
            gradient={questions[Math.min(currentQuestion, questions.length - 1)].gradient}
        />

        {#key currentQuestion}
            <QuestionContainer
                question={questions[currentQuestion]}
                answers={answers[currentQuestion]}
                onAnswer={handleSelection}
                onNavigate={handleNavigate}
            />
        {/key}
      {:else}
          <ThankYou thankYouMessage={thankYouMessage} />
    {/if}
    </div>
</div>
  
  <style>
  .container {
      display: flex;
      justify-content: center;
      min-height: 100vh;
      align-items: center;
      font-family: 'Outfit', sans-serif;
      z-index: -2;
  }

  .questionnaire {
      width: 90%;
      max-width: 1000px;
      margin: 2rem auto;
      padding: 3rem;
      background: rgba(255, 255, 255, 0.1);
      backdrop-filter: blur(10px);
      border-radius: 30px;
      box-shadow: 
        0 20px 60px rgba(0,0,0,0.3),
        inset 0 1px 2px rgba(255,255,255,0.1);
      border: 1px solid rgba(255,255,255,0.1);
  }
  
  
  @media (max-width: 1100px) {
    .container{
      padding: 1rem;
    }
  }

  @media (max-width: 768px) {
    .questionnaire {
        width: 95%;
        padding: 1.5rem;
        margin: 1rem auto;
    }
  }

  @media (max-width: 480px) {
    .questionnaire {
      padding: 1rem;
    }
  }
  </style>

<script lang="ts">
    import ProgressBar from '$lib/components/FormShow/ProgressBar.svelte';
    import QuestionContainer from '$lib/components/FormShow/QuestionContainer.svelte';
    import FloatingShapes from './FloatingShapes.svelte';
    import ThankYou from './ThankYou.svelte';
    
    import { predefinedThemes, type ThemeName } from '$lib/types';
    import type { Question, ThankYouMessage, Theme } from '$lib/types';
    import type { FloatingShapesTheme } from '$lib/types';
    
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
            
            // Check if nextQuestion exists first
            const nextQuestion = questions[currentQuestion].nextQuestion;
            if (!nextQuestion) {
                // If no nextQuestion property exists, this is the last question
                handleSubmit();
                return;
            }

            // If nextQuestion exists, find the next question ID
            const nextId = nextQuestion.conditions?.find(c => c.answer === answer)?.nextId 
                || nextQuestion.default;
            
            handleNavigate(nextId);
        }
    };

    const handleSubmit = async () => {
        console.log('Form Answers:', answers);
        isFormSubmitted = true;
        // submit logic here
    };

  </script>
  
  <div 
    class="min-h-screen flex justify-center items-center font-sans" 
    style="background-color: {activeTheme.backgroundColor}; color: {activeTheme.fontColor};">

    <FloatingShapes theme={floatingShapesTheme} zIndex={0} />
    
    <div class="
        w-[90%] 
        max-w-[1000px]
        p-12 
        bg-white/10 
        backdrop-blur-lg 
        rounded-3xl 
        shadow-2xl 
        shadow-black/30 
        border 
        border-white/10
    ">
      {#if !isFormSubmitted}
        <ProgressBar 
            currentQuestion={Math.min(currentQuestion, questions.length - 1)}
            totalQuestions={questions.length}
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
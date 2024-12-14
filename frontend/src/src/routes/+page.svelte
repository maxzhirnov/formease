<script lang="ts">
    import Form from "$lib/components/Form/Form.svelte";
    import type { ThemeName } from '$lib/themes/themes';
    import type { FloatingShapesTheme } from "$lib/themes/floatingShapes";
    import type { ThankYouMessage, Question, Theme } from "$lib/types";

    const questions : Question[] = [
      {
        id: 1,
        type: 'single-choice',
        question: "What brings you here today?",
        subtext: "Help us personalize your experience",
        image: "/img/demo.jpg",
        gradient: "from-emerald-500 to-blue-500",
        bgColor: "#10B981",
        options: [
          { text: "Building a Website", icon: "🎨", color: "#34D399" },
          { text: "Mobile App", icon: "📱", color: "#60A5FA" },
          { text: "Just Exploring", icon: "🔍", color: "#818CF8" },
        ],
        nextQuestion: {
          conditions: [
            { answer: "Building a Website", nextId: 2 },
            { answer: "Mobile App", nextId: 3 },
            { answer: "Just Exploring", nextId: 4 }
          ],
        },
      },
      {
        id: 2,
        type: 'multiple-choice',
        question: "What features do you need?",
        subtext: "Select up to 4 key features",
        image: "/img/demo.jpg",
        gradient: "from-pink-500 to-purple-500",
        bgColor: "#EC4899",
        maxSelections: 4,
        options: [
          { text: "E-commerce", icon: "🛍️", color: "#F472B6" },
          { text: "Blog", icon: "✍️", color: "#E879F9" },
          { text: "Authentication", icon: "🔒", color: "#C084FC" },
          { text: "Analytics", icon: "📊", color: "#A855F7" },
          { text: "SEO Tools", icon: "🎯", color: "#D946EF" },
        ],
        nextQuestion: {
          conditions: [],
          default: 5
        }
      },
      {
        id: 3,
        type: 'multiple-choice',
        question: "Which platforms are you targeting?",
        subtext: "Select all that apply",
        image: "/img/demo.jpg",
        gradient: "from-orange-500 to-red-500",
        bgColor: "#F97316",
        maxSelections: 3,
        options: [
          { text: "iOS", icon: "🍎", color: "#FB923C" },
          { text: "Android", icon: "🤖", color: "#F87171" },
          { text: "Both", icon: "📱", color: "#FCA5A5" },
        ],
        nextQuestion: {
          conditions: [],
          default: 5
        }
      },
      {
        id: 4,
        type: 'input',
        question: "What's your main interest?",
        subtext: "Tell us what you'd like to learn more about",
        image: "/img/demo.jpg",
        gradient: "from-teal-500 to-cyan-500",
        bgColor: "#14B8A6",
        inputType: 'text',
        placeholder: "Enter your interest",
        validation: /.+/,
        nextQuestion: {
          conditions: [],
          default: 5
        }
      },
      {
        id: 5,
        type: 'input',
        question: "What's your email?",
        subtext: "We'll send you personalized recommendations",
        image: "/img/demo.jpg",
        gradient: "from-indigo-500 to-purple-500",
        bgColor: "#6366F1",
        inputType: 'email',
        placeholder: "your@email.com",
        validation: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
        nextQuestion: {
          conditions: [],
          default: 6
        }
      },
      {
        id: 6,
        type: 'single-choice',
        question: "Would you like to receive our newsletter?",
        subtext: "Get weekly updates and tips",
        image: "/img/demo.jpg",
        gradient: "from-rose-500 to-pink-500",
        bgColor: "#F43F5E",
        options: [
          { text: "Yes, please!", icon: "✨", color: "#FB7185" },
          { text: "Maybe later", icon: "⏰", color: "#FDA4AF" },
        ]
      }
    ];

    const thankYouMessage : ThankYouMessage = {
        title: "Спасибо вам огромное!", 
        subtitle: "За то что прошли нашу форму!.", 
        icon: "✨",
        button: {
            text: "Перейти",
            url: "https://hemaks.org",
            newTab: true
        }
    }


    let theme: ThemeName = 'tech';
    const floatingShapesTheme : FloatingShapesTheme = 'spring';
</script>

<Form 
    questions={questions} 
    thankYouMessage={thankYouMessage} 
    theme={theme}
    floatingShapesTheme={floatingShapesTheme}
/>

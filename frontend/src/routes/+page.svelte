<script lang="ts">
    import { goto } from '$app/navigation';
    import Logo from '$lib/components/Header/Logo.svelte';
    import { Button, Card, Timeline, TimelineItem, Accordion, AccordionItem, FooterLinkGroup, FooterIcon, FooterCopyright, FooterLink, FooterBrand, Footer } from 'flowbite-svelte';
    import { Section, Cta } from "flowbite-svelte-blocks"
    import { ArrowRightOutline, WandMagicSparklesSolid, ChartMixedDollarSolid, UsersSolid } from 'flowbite-svelte-icons';
  
    const features = [
      {
        title: "Создание форм с помощью AI",
        description: "Позвольте нашему AI-ассистенту создавать профессиональные формы на основе вашего описания.",
        icon: WandMagicSparklesSolid
      },
      {
        title: "Аналитика в реальном времени",
        description: "Получайте мгновенный анализ ответов с помощью наглядных визуализаций.",
        icon: ChartMixedDollarSolid
      },
      {
        title: "Командная работа",
        description: "Работайте вместе с членами вашей команды над созданием форм и анализом результатов.",
        icon: UsersSolid
      }
    ];
  
    const faqs = [
      {
        question: "Как работает создание форм с помощью ИИ?",
        answer: "Просто опишите назначение вашей формы, и наш ИИ создаст профессиональную форму с релевантными вопросами и правильной структурой."
      },
      {
        question: "Можно ли настраивать сгенерированные формы?",
        answer: "Да! Хотя ИИ создает начальную структуру, у вас есть полный контроль над изменением, добавлением или удалением вопросов по необходимости."
      },
      {
        question: "Есть ли бесплатный тариф?",
        answer: "Да, мы предлагаем щедрый бесплатный тариф, который включает генерацию форм с помощью ИИ и базовую аналитику."
      }
    ];
  
    function handleGetStarted() {
      goto('/register');
    }
</script>
  
<div class="relative min-h-screen overflow-hidden">
    <!-- Background Effects -->
    <div class="absolute top-[-20%] right-[-10%] w-[600px] h-[600px] rounded-full bg-gradient-to-r from-blue-500/10 to-purple-500/10 blur-[80px] opacity-50 pointer-events-none"></div>
    <div class="absolute bottom-[-30%] left-[-15%] w-[600px] h-[600px] rounded-full bg-gradient-to-r from-purple-500/10 to-blue-500/10 blur-[80px] opacity-50 pointer-events-none rotate-45"></div>
  
    <!-- Hero Section -->
    <section class="py-40 px-4 relative z-10">
      <div class="text-center max-w-3xl mx-auto">
        <h1 class="text-5xl md:text-6xl font-extrabold mb-6 text-gray-900">
          Создавайте формы с помощью 
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-purple-500">
            магии AI
          </span>
        </h1>
        
        <p class="text-xl text-gray-600 mb-10 max-w-2xl mx-auto">
          Мгновенно превращайте ваши идеи в профессиональные формы. Наша платформа с AI делает создание форм простым,
          позволяя сосредоточиться на самом важном - сборе ценной информации.
        </p>
        
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <Button 
            size="xl" 
            class="bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600"
            on:click={handleGetStarted}
          >
            Начать бесплатно
            <ArrowRightOutline class="w-4 h-4 ml-2" />
          </Button>
          
          <Button 
            size="xl" 
            color="light"
            href="#features"
          >
            Узнать больше
          </Button>
        </div>
      </div>
    </section>
  
    <!-- Features Section -->
    <section id="features" class="py-20 px-4 bg-gray-50/50 relative z-10">
      <div class="max-w-6xl mx-auto">
        <h2 class="text-3xl font-bold text-center mb-12">
          Всё необходимое для создания идеальных форм
        </h2>
  
        <div class="grid md:grid-cols-3 gap-8">
          {#each features as feature}
            <Card class="text-center hover:shadow-lg transition-shadow">
              <svelte:component this={feature.icon} class="w-12 h-12 mx-auto mb-4 text-purple-500" />
              <h3 class="text-xl font-bold mb-2">{feature.title}</h3>
              <p class="text-gray-600">{feature.description}</p>
            </Card>
          {/each}
        </div>
      </div>
    </section>
  
    <!-- How It Works Section -->
    <section class="py-20 px-4 relative z-10">
      <div class="max-w-4xl mx-auto">
        <h2 class="text-3xl font-bold text-center mb-12">
          Как это работает
        </h2>
  
        <Timeline class="border-l border-purple-200">
          <TimelineItem title="Опишите вашу форму" date="Шаг 1">
            Расскажите нашему ИИ, какая форма вам нужна - от опросов до регистрационных форм.
          </TimelineItem>
          <TimelineItem title="ИИ создаёт вашу форму" date="Шаг 2">
            Наблюдайте, как ИИ создает профессиональную форму с релевантными вопросами и структурой.
          </TimelineItem>
          <TimelineItem title="Настройте при необходимости" date="Шаг 3">
            Доработайте форму с помощью нашего интуитивного редактора или используйте как есть.
          </TimelineItem>
          <TimelineItem title="Делитесь и собирайте ответы" date="Шаг 4">
            Поделитесь формой и начните собирать ответы мгновенно.
          </TimelineItem>
        </Timeline>
      </div>
    </section>
  
    <!-- FAQ Section -->
    <section class="py-20 px-4 bg-gray-50/50 relative z-10">
      <div class="max-w-3xl mx-auto">
        <h2 class="text-3xl font-bold text-center mb-12">
          Часто задаваемые вопросы
        </h2>
  
        <Accordion>
          {#each faqs as faq}
            <AccordionItem>
              <svelte:fragment slot="header">{faq.question}</svelte:fragment>
              <p class="mb-2 text-gray-600">{faq.answer}</p>
            </AccordionItem>
          {/each}
        </Accordion>
      </div>
    </section>
  
    <!-- CTA Section -->
    <Section name="cta" sectionClass="py-[200px] px-4 relative z-10">
        <Cta ctatype="heading">
          <svelte:fragment slot="h2">Готовы изменить процесс создания форм?</svelte:fragment>
          <p class="mb-12 font-light text-gray-500 dark:text-gray-400 md:text-lg">Присоединяйтесь к тысячам пользователей, которые уже создают потрясающие формы с помощью AI.</p>
          <Button 
          size="xl" 
          class="bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600"
          on:click={handleGetStarted}
        >
          Начать создавать формы
          <ArrowRightOutline class="w-4 h-4 ml-2" />
        </Button>
        </Cta>
    </Section>

    <!-- Footer Section -->
    <script>
        import { Section } from 'flowbite-svelte-blocks';
        import { Facebook, Instagram, Github, Twitter } from '../utils';
        import { Footer, FooterBrand, FooterLinkGroup, FooterLink, FooterCopyright, FooterIcon } from 'flowbite-svelte';
      </script>
      
      <Footer footerType="socialmedia">
        <div class="grid gap-12 lg:grid-cols-4 lg:gap-24">
          <div class="col-span-2">
            <div class="flex items-center mb-4">
              <Logo/>
            </div>
            <p class="text-gray-600 dark:text-gray-400">
              FormEase - это современная платформа для создания форм с использованием искусственного интеллекта. 
              Создавайте профессиональные формы за считанные минуты и получайте ценные данные от ваших пользователей.
            </p>
          </div>
          
          <div>
            <h2 class="mb-6 text-sm font-semibold text-gray-400 uppercase dark:text-white">Поддержка</h2>
            <FooterLinkGroup>
              <FooterLink liClass="mb-4" href="/contact">Связаться с нами</FooterLink>
              <FooterLink liClass="mb-4" href="/docs">Документация</FooterLink>
              <FooterLink liClass="mb-4" href="/faq">Частые вопросы</FooterLink>
            </FooterLinkGroup>
          </div>
      
          <div>
            <h2 class="mb-6 text-sm font-semibold text-gray-400 uppercase dark:text-white">Правовая информация</h2>
            <FooterLinkGroup>
              <FooterLink liClass="mb-4" href="/privacy">Конфиденциальность</FooterLink>
              <FooterLink liClass="mb-4" href="/terms">Условия использования</FooterLink>
              <FooterLink liClass="mb-4" href="/cookies">Политика cookies</FooterLink>
            </FooterLinkGroup>
          </div>
        </div>
      
        <hr class="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
      
        <div class="sm:flex sm:items-center sm:justify-between">
          <FooterCopyright href="/" by="FormEase" year={2024}>
            Все права защищены
          </FooterCopyright>
          
          <div class="flex mt-4 space-x-6 sm:justify-center sm:mt-0">
            <FooterIcon href="https://github.com"class="text-gray-400 hover:text-gray-900" />
            <FooterIcon href="https://discord.com" class="text-gray-400 hover:text-gray-900" />
            <FooterIcon href="https://twitter.com" class="text-gray-400 hover:text-gray-900" />
            <FooterIcon href="https://t.me"class="text-gray-400 hover:text-gray-900" />
          </div>
        </div>
      </Footer>
</div>

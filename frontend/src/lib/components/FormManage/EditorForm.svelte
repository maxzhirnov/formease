<script lang="ts">
    import { page } from '$app/stores';
    import { formStore } from '$lib/stores/formStore';
    import { QuestionType, type Question } from '$lib/types';
    import { formService } from '$lib/services/formService'; 
    import { themeOptionsList } from '$lib/types';
    import { Modal, Button, Label, Input, Select, Hr, Card } from 'flowbite-svelte';
    import QuestionEditor from './EditorQuestion.svelte';

    // State management
    let questionType = $state(QuestionType.SINGLE_CHOICE);
    let selectedQuestionId: number | null = $state(null);
    let openModal = $state(false);

    // Computed properties from service
    let currentForm = $state(formService.getCurrentForm())
    let selectedTheme = $state(currentForm.theme || 'light')
    
    const questionTypeOptions = Object.entries(QuestionType).map(([key, value]) => ({
        value,
        name: key.split('_').map(word => 
            word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
        ).join(' ')
    }));
    
    // Set user ID when authenticated
    // if (userState.id) {
    //     formService.setUserId(userState.id);
    // }
    
    function handleAddQuestion() {
        formService.addQuestion(questionType);
    }

    function selectQuestion(id: number) {
        selectedQuestionId = id;
        openModal = true;
    }

    function handleQuestionUpdate(updatedQuestion: Question) {
        formService.updateQuestion(updatedQuestion.id, updatedQuestion);
        openModal = false;
    }

    function handleNameChange(event: Event) {
        const target = event.target as HTMLInputElement;
        formService.updateFormName(target.value);
    }

    function handleChangeTheme(event: Event) {
        const target = event.target as HTMLSelectElement;
        formService.updateTheme(target.value)
    }
    
    async function handleSave() {
        const currentForm = formService.getCurrentForm();
        
        if (!currentForm.questions) {
            return;
        }
        
        const hasEmptyQuestions = currentForm.questions.some(q => !q.question.trim());
        if (hasEmptyQuestions) {
            alert('Please ensure all questions have text before saving.');
            return;
        }

        try {
            await formService.update($page.params.id);
        } catch (error: unknown) {
            if (error instanceof Error) {
                alert(error.message);
            } else {
                alert('An error occurred while saving the form');
            }
            console.error('Error saving form:', error);
        }
    }
</script>

<Modal 
    bind:open={openModal}
    title="Edit Question"
    size="xl"
    placement="center"
    outsideclose={true}
>
    {#if selectedQuestionId !== null}
        {@const selectedQuestion = $formStore.questions!.find(q => q.id === selectedQuestionId)}
        {#if selectedQuestion}
            <QuestionEditor 
                question={selectedQuestion}
                updateQuestionCallback={handleQuestionUpdate}
            />
        {/if}
    {/if}
</Modal>

<div class="w-auto max-w-[1000px] mx-auto p-6 space-y-6">
    <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-purple-500">
            Edit Form
        </h1>
    </div>

    <Card class="max-w-none w-full">
        <div class="grid md:grid-cols-2 gap-6">
            <div>
                <Label for="formName" class="mb-2">Form Name</Label>
                <Input 
                    id="formName" 
                    type="text"
                    value={$formStore.name} 
                    on:input={handleNameChange}
                    required
                    placeholder="Enter form name" 
                />
            </div>
            <div>
                <Label>Theme</Label>
                <Select 
                    items={themeOptionsList} 
                    bind:value={selectedTheme}
                    on:change={handleChangeTheme} 
                />
            </div>
        </div>

        <Hr/>

        <div class="flex items-end space-x-4">
            <div class="flex-grow">
                <Label>Question Type</Label>
                <Select
                    items={questionTypeOptions} 
                    bind:value={questionType}
                />
            </div>
            <Button color="blue" on:click={handleAddQuestion}>
                Add Question
            </Button>
        </div>

        {#if $formStore.questions && $formStore.questions.length > 0}
            <div class="space-y-4">
                <h2 class="text-xl font-semibold mt-6">Questions</h2>
                <div class="grid gap-4">
                    {#each $formStore.questions as question (question.id)}
                        <div 
                            class="p-4 bg-gray-100 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors 
                                   {selectedQuestionId === question.id ? 'ring-2 ring-blue-500' : ''}"
                            onclick={() => selectQuestion(question.id)}
                            onkeydown={(e) => e.key === 'Enter' && selectQuestion(question.id)}
                            role="button"
                            tabindex="0"
                        >
                            <div class="flex justify-between items-center">
                                <div>
                                    <span class="font-bold mr-2">Q{question.id}</span>
                                    <span>{question.question || 'Untitled Question'}</span>
                                </div>
                                <span class="text-sm text-gray-500">{question.type}</span>
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        {/if}

        <div class="text-center mt-6">
            <Button color="green" on:click={handleSave}>
                Update Form
            </Button>
        </div>
    </Card>
</div>

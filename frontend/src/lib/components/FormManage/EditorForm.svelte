<script lang="ts">
    import { Modal, Button, Label, Input, Select, Hr, Card, ButtonGroup, Checkbox, Spinner } from 'flowbite-svelte';
    import { ArrowLeftOutline, EyeSolid, TrashBinSolid, CheckCircleSolid } from 'flowbite-svelte-icons';

    import { page } from '$app/stores';
    import { goto } from '$app/navigation';

    import { formService } from '$lib/services/formService'; 
    import { QuestionType, type Question } from '$lib/types/questions';
    import  { type ThankYouMessage, defaultThankYouMessage } from '$lib/types/thank';
    import {getThemeNameFromTheme, themeOptionsList } from '$lib/types/theme';

    import QuestionEditor from './EditorQuestion.svelte';
    import ThankYouScreen from './ThankYouScreen.svelte';
    import { floatingShapesOptionsList, getFloatingShapesThemeNameFromTheme } from '$lib/types/shapes';

    // State management
    let currentForm = $derived(formService.state.getCurrentForm())
    let selectedTheme = $derived(currentForm.theme || 'light')
    let selectedFloatingShapesTheme = $derived(currentForm.floatingShapesTheme || 'default')
    let questionType = $state(QuestionType.SINGLE_CHOICE);
    let selectedQuestionId: number | null = $state(null);
    let selectedQuestion = $derived(currentForm.questions!.find(q => q.id === selectedQuestionId))
    let hasQuestions = $derived((currentForm.questions?.length ?? 0) > 0);
    let isSaving = $state(false);
    let saveSuccess = $state(false);

    // Delete question modal
    let questionToDelete = $state(0);
    let showDeleteModal = $state(false);

    let thankYouScreen: ThankYouMessage  = $derived(currentForm.thankYouMessage || defaultThankYouMessage);

    // Modals
    let openEditQuestionModal = $state(false);
    let openEditThankYouModal = $state(false);

    const questionTypeOptions = Object.entries(QuestionType).map(([key, value]) => ({
        value,
        name: key.split('_').map(word => 
            word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
        ).join(' ')
    }));
    
    function handleAddQuestion() {
        formService.state.addQuestion(questionType);
        console.log('formState', formService.state.getCurrentForm());
        handleSave()
    }

    function handleDeleteQuestion(id: number) {
        if (questionToDelete === 0) {
            return
        }
        formService.state.removeQuestion(id);
        formService.api.update(currentForm.id);
        questionToDelete = 0;
        showDeleteModal = false;
    }

    function selectQuestion(id: number) {
        selectedQuestionId = id;
        openEditQuestionModal = true;
    }

    function handleNameChange(event: Event) {
        const target = event.target as HTMLInputElement;
        formService.state.updateFormName(target.value);
        // handleSave() // do not save on any change to reduce server load
    }

    function handleChangeTheme(event: Event) {
        const target = event.target as HTMLSelectElement;
        formService.state.updateTheme(getThemeNameFromTheme(target.value));
        handleSave();
    }

    function handleChangeFloatingShapesTheme(event: Event) {
        const target = event.target as HTMLSelectElement;
        formService.state.updateFloatingShapesTheme(getFloatingShapesThemeNameFromTheme(target.value));
        handleSave();
    }

    function handleQuestionUpdate(updatedQuestion: Question) {
        formService.state.updateQuestion(updatedQuestion.id, updatedQuestion);
        openEditQuestionModal = false;
        handleSave()
    }

    function handleBackToForms() {
        goto('/app');
    }
    
    async function handleSave() {
        isSaving = true;
        saveSuccess = false;


        const currentForm = formService.state.getCurrentForm();
        
        if (!currentForm.questions) {
            return;
        }

        try {
            await formService.api.update($page.params.id);
            saveSuccess = true;

            setTimeout(() => {
                saveSuccess = false;
            }, 2000);
        } catch (error: unknown) {
            if (error instanceof Error) {
                alert(error.message);
            } else {
                alert('An error occurred while saving the form');
            } 
            console.error('Error saving form:', error);
        } finally {
            isSaving = false;
        }
    }
</script>

<!-- Edit Question Modal -->
<Modal 
    bind:open={openEditQuestionModal}
    title="Edit Question"
    size="xl"
    placement="center"
    outsideclose={true}
>
    {#if selectedQuestionId !== null}
        {#if selectedQuestion}
            <QuestionEditor 
                question={selectedQuestion}
                updateQuestionCallback={handleQuestionUpdate}
            />
        {/if}
    {/if}
</Modal>
<!-- End Edit Question Modal -->

<!-- Control Buttons -->
{#snippet controlButtons()}
    <div class="text-center">
        <ButtonGroup class="flex-row">
            <Button 
                color="light" 
                on:click={handleBackToForms}
                class="flex items-center"
            >
                <ArrowLeftOutline class="w-4 h-4 mr-2" />
                <span>Back</span>
            </Button>
            
            
            <Button 
                color="light" 
                disabled={!hasQuestions}
                on:click={() => goto(`/app/form/${currentForm.id}/preview`)}
            >
                <EyeSolid class="w-4 h-4" />
                <span class="ml-1  md:inline">Preview</span>
            </Button> 

            <Button 
                color='green'
                on:click={handleSave}
                disabled={isSaving}
                class="min-w-[130px] transition-all duration-200"
            >
                {#if isSaving}
                    <Spinner class="mr-2" size="4" />
                    Saving...
                {:else if saveSuccess}
                    <CheckCircleSolid class="w-4 h-4 mr-2" />
                    Saved!
                {:else}
                    Save Changes
                {/if}
            </Button>
        </ButtonGroup>
    </div>
{/snippet}
<!-- End Control Buttons -->

<div class="w-auto max-w-[1000px] mx-auto p-6 space-y-6">
    <div class="text-center mb-8">
        <h1 class="text-2xl text-gray-600">
            Edit Form
        </h1>
    </div>

    {@render controlButtons()}
    <Card class="max-w-none w-full">
        <!-- Name and Theme -->
        <h2 class="text-xl font-semibold mt-6">Form name and theme</h2>
        <div class="grid mt-2 md:grid-cols-2 gap-6">
            <div>
                <Label for="formName">Form Name</Label>
                <Input 
                    id="formName" 
                    type="text"
                    value={currentForm.name} 
                    on:input={handleNameChange}
                    required
                    placeholder="Enter form name" 
                />
            </div>
            <div>
                <Label for="theme">Theme</Label>
                <Select 
                    id="theme"
                    items={themeOptionsList} 
                    value={selectedTheme}
                    on:change={handleChangeTheme} 
                />
                <Label for="theme">Background Theme</Label>
                <Select 
                    id="floatingShapesTheme"
                    items={floatingShapesOptionsList} 
                    value={selectedFloatingShapesTheme}
                    on:change={handleChangeFloatingShapesTheme}
                />
            </div>
        </div>
        <!-- End Name and Theme -->

        <Hr/>

        <!-- Add Question -->
        <h2 class="text-xl font-semibold mt-6">Add Question</h2>
        <div class="flex items-end space-x-4 mt-4">

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
        <!-- End Add Question -->

        <!-- Questions -->
        {#if currentForm.questions && formService.state.hasQuestions()}
            <div class="space-y-4">
                <h2 class="text-xl font-semibold mt-6">Questions</h2>
                <div class="grid gap-4">
                    {#each currentForm.questions as question (question.id)}
                        <div class="flex items-center space-x-2">
                            <div 
                                class="flex-1 p-2 bg-gray-100 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors 
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
                                    <span class="text-sm text-gray-500 mr-2">{question.type}</span>
                                        
                                </div>
                            </div>
                            <Button color=red class="w-8 h-10" onclick={() => {
                                questionToDelete = question.id;
                                showDeleteModal = true;
                            }}>
                                <TrashBinSolid class="w-4 h-4" />
                            </Button>
                        </div>
                    {/each}
                </div>
            </div>
        {/if}
        <!-- End Questions -->

        <Hr/>
        
        <!-- Thank You Screen -->
         <h2 class="text-xl font-semibold mt-6">Thank you screen</h2>
         <div class="my-8">
            <Button color="blue" on:click={() => openEditThankYouModal = true}>Edit Thank You Screen</Button>
         </div>
         <Modal
            bind:open={openEditThankYouModal}
            title="Edit Question"
            size="xl"
            placement="center"
            autoclose
            outsideclose={true}>
            <ThankYouScreen thankYouScreen={thankYouScreen} handleSaveCallback={handleSave}/>
        </Modal>
        <!-- End Thank You Screen -->
    </Card>
    {@render controlButtons()}
</div>

<Modal title="" bind:open={showDeleteModal} autoclose size="sm" class="w-full">
    <svg class="text-gray-400 dark:text-gray-500 w-11 h-11 mb-3.5 mx-auto" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" /></svg>
    <p class="mb-4 text-gray-500 dark:text-gray-300 text-center">Are you sure you want to delete this question?</p>
    <div class="flex justify-center items-center space-x-4">
      <Button color="light" on:click={() => showDeleteModal = false}>No, cancel</Button>
      <Button color="red" on:click={() => handleDeleteQuestion(questionToDelete)}>Yes, I'm sure</Button>
    </div>
</Modal>
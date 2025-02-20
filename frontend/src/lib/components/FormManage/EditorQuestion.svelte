<script lang="ts">
    import { Label, Input, Button, Select, Checkbox, Img, Modal, TableBodyRow, TableBody, TableBodyCell, Table, TableHeadCell, TableHead } from 'flowbite-svelte';
    import { CirclePlusSolid, ImageSolid, TrashBinSolid, UploadSolid } from 'flowbite-svelte-icons';
    import type { ValidatorName } from '$lib/validators/formValidators';
    import { validators } from '$lib/validators/formValidators';
    import { formService } from '$lib/services/formService';
    import { QuestionType, type InputQuestion, type MultipleChoiceQuestion, type Question, type QuestionOption } from '$lib/types/questions';
    import ImageLibrary from '../ImageLibrary.svelte';
    import { showError } from '$lib/utils/errorHandler';

    interface Props {
        question: Question;
        updateQuestionCallback: (question: Question) => void;
    }

    let { question, updateQuestionCallback }: Props = $props();

    // Local state for editing
    let questionText = $state(question.question);
    let subtext = $state(question.subtext || '');
    let imageUrl = $derived(question.image || '');
    let validation: ValidatorName = $state((question as InputQuestion)?.validation || 'text');
    let selectedOptionId: number = $state(0);
    let showImageLibraryModal = $state(false);
    let showImageOptionModal = $state(false);
    let defaultNextId = $state(question.nextQuestion?.default);
    
    // For INPUT type
    let inputType = $state((question as InputQuestion).inputType || 'text');
    let placeholder = $state((question as InputQuestion).placeholder || '');

    // For MULTIPLE_CHOICE type
    let maxSelections = $state((question as MultipleChoiceQuestion).maxSelections || 1);

    const validationOptions = Object.entries(validators).map(([key, pattern]) => ({
        value: key as ValidatorName,
        label: key.charAt(0).toUpperCase() + key.slice(1).replace(/([A-Z])/g, ' $1'),
        pattern: pattern.toString().slice(1, -1)
    }));

    const inputTypes = [
        { value: 'text', name: 'Text' },
        { value: 'email', name: 'Email' },
        { value: 'number', name: 'Number' }
    ];

    function handleAddOption(): void {
        formService.state.addQuestionOption(question.id, '', '');
    }

    function removeOption(index: number): void {
        formService.state.removeQuestionOption(question.id, index);
    }

    function handleAddCondition(): void {
        if (question.type === QuestionType.SINGLE_CHOICE || question.type === QuestionType.MULTIPLE_CHOICE) {
            const conditions = question.nextQuestion?.conditions || [];
            question.nextQuestion = {
                ...question.nextQuestion,
                conditions: [...conditions, { answer: '', nextId: 1 }]
            };
        }
    }

    function handleRemoveCondition(index: number): void {
        if (question.nextQuestion?.conditions) {
            question.nextQuestion.conditions = question.nextQuestion.conditions.filter((_, i) => i !== index);
        }
    }

    function applyChanges(): void {
        const updatedQuestion: Question = {
            ...question,
            question: questionText,
            subtext,
            image: imageUrl,
            nextQuestion: {
                conditions: question.nextQuestion?.conditions || [], 
                default: defaultNextId
            }
        };

        // Add type-specific updates
        if (question.type === QuestionType.INPUT) {
            console.log('validation: ', validation);
            (updatedQuestion as InputQuestion).inputType = inputType;
            (updatedQuestion as InputQuestion).placeholder = placeholder;
            (updatedQuestion as InputQuestion).validation = validation;
        }

        if (question.type === QuestionType.MULTIPLE_CHOICE) {
            (updatedQuestion as MultipleChoiceQuestion).maxSelections = maxSelections;
        }

        updateQuestionCallback?.(updatedQuestion);
    }

    async function handleQuesionImageChosen(imageUrl: string) {
        // console.log('handleImageUploadSuccess', imageUrl);
        try {
            // Update question with image
            formService.state.updateQuestion(question.id, {
                image: imageUrl
            });

            // Wait a short moment to ensure state is updated
            await new Promise(resolve => setTimeout(resolve, 100));

            // Get the current form state after update
            const currentForm = formService.state.getCurrentForm();

            // Ensure form ID exists before updating
            if (!currentForm.id) {
                console.error('Form ID is missing');
                return;
            }

            // Update form with new state
            await formService.api.update(currentForm.id);
        } catch (error) {
            console.error('Failed to handle image upload:', error);
            showError(error, 'Failed to handle image upload');
            // Optionally show error to user
        } finally {
            showImageLibraryModal = false;
        }
    }
    

    function startOptionImageSelect(optionId: number) {
        selectedOptionId = optionId;
        showImageOptionModal = true;
    }

    async function handleOptionImageChosen(imageUrl: string) {
        console.log('handleOptionImageChosen', imageUrl);
        showImageOptionModal = false;
        try {
            formService.state.updateQuestionOption(question.id, selectedOptionId, {
                image: imageUrl
            });

            // Wait a short moment to ensure state is updated
            await new Promise(resolve => setTimeout(resolve, 100));

            // Get the current form state after update
            const currentForm = formService.state.getCurrentForm();

            // Ensure form ID exists
            if (!currentForm.id) {
                console.error('Form ID is missing');
                showError('Form ID is missing');
                return;
            }

            // Update form with new state
            await formService.api.update(currentForm.id);
        } catch (error) {
            console.error('Failed to handle image upload:', error);
            showError(error, 'Failed to handle image upload');
            
        } finally {
            selectedOptionId = 0;
            showImageOptionModal = false;
        }
    }
</script>
<Modal 
    title="Image Library" 
    autoclose={false}
    bind:open={showImageLibraryModal} 
    size="xl"
>
    <ImageLibrary onImageChosen={handleQuesionImageChosen}/>
</Modal>

<Modal 
    title="Image Library" 
    autoclose={false}
    bind:open={showImageOptionModal} 
    size="xl"
>
    <ImageLibrary onImageChosen={handleOptionImageChosen}/>
</Modal>

<div class="space-y-6 p-6 bg-white/10 backdrop-blur-lg rounded-2xl border border-white/10">
    <h2 class="text-2xl font-bold mb-6 text-gray-800">
        Edit Question {question.id}
    </h2>

    <div class="space-y-4">
        <Label>
            Question Text
            <Input 
                type="text" 
                placeholder="Enter question text" 
                bind:value={questionText}
                required 
            />
        </Label>

        <Label>
            Subtext
            <Input 
                type="text" 
                placeholder="Enter subtext" 
                bind:value={subtext}
            />
        </Label>
        
        <div class="space-y-3">
            <!-- <Label class="space-y-1">
                <span class="text-sm font-medium text-gray-700">Image URL</span>
                <Input 
                    disabled
                    type="url" 
                    placeholder="Enter image URL" 
                    value={imageUrl}
                    class="bg-gray-50"
                />
            </Label> -->
            <Label class="space-y-1">Image</Label>
            <div class="flex flex-col md:flex-row gap-3 items-start bg-gray-50 rounded-lg p-3 border border-gray-200">
                <div class="w-full md:w-auto">
                    <Button 
                        color="blue" 
                        class="w-full" 
                        on:click={() => showImageLibraryModal = true}
                    >
                        <ImageSolid class="w-4 h-4 mr-2" />
                        Pick Image
                    </Button>
                </div>
                
                {#if imageUrl}
                    <div class="flex justify-center">
                        <div class="relative group">
                            <Img 
                                size="max-w-xs" 
                                src={imageUrl} 
                                alt="Question illustration" 
                                class="h-32 w-full object-contain rounded-lg border border-gray-200" 
                            />
                           
                        </div>
                    </div>
                {/if}
            </div>
        </div>
        

        <!-- Options for Single/Multiple Choice -->
        {#if question.type === QuestionType.SINGLE_CHOICE || question.type === QuestionType.MULTIPLE_CHOICE}
            <div class="space-y-4">
                <h3 class="text-xl font-semibold">Options</h3>
                <Table>
                    <TableHead>
                        <TableHeadCell class="w-20">Icon</TableHeadCell>
                        <TableHeadCell>Option Text</TableHeadCell>
                        <TableHeadCell class="w-20">Image</TableHeadCell>
                        <TableHeadCell class="w-20">Delete</TableHeadCell>
                    </TableHead>
                    <TableBody>
                        {#each question.options as option, index}
                            <TableBodyRow>
                                <TableBodyCell>
                                    <Input
                                        type="text"
                                        size="sm"
                                        class="w-full"
                                        bind:value={option.icon}
                                    />
                                </TableBodyCell>
                                <TableBodyCell tdClass="min-w-[50px]">
                                    <Input 
                                        type="text"     
                                        size="sm"
                                        placeholder="Enter option text" 
                                        class="w-full"
                                        bind:value={option.text} 
                                        required    
                                    />
                                </TableBodyCell>
                                <TableBodyCell>
                                    {#if !option.image || option.image == ''}
                                        <Button color="green" class="w-12 h-12" onclick={() => startOptionImageSelect(option.id)}><UploadSolid/></Button>
                                    {:else}
                                        <div class="relative group w-12 h-12">
                                            <Img 
                                                src={option.image} 
                                                alt="Option image" 
                                                class="w-12 h-12 object-cover rounded-lg" 
                                            />
                                            <Button 
                                                color="green" 
                                                class="absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity"
                                                onclick={() => startOptionImageSelect(option.id)}
                                            >
                                                <UploadSolid/>
                                            </Button>
                                        </div>
                                    {/if}
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Button class="w-12 h-12" outline color="red" on:click={() => removeOption(index)}>
                                        <TrashBinSolid/>
                                    </Button>
                                </TableBodyCell>
                            </TableBodyRow>
                        {/each}
                    </TableBody>
                </Table>
                
                <Button color="light" on:click={handleAddOption}>
                    <CirclePlusSolid class="mr-2" /> Add Option
                </Button>

                {#if question.type === QuestionType.MULTIPLE_CHOICE}
                    <Label>
                        Max Selections
                        <Input 
                            type="number" 
                            placeholder="Enter max selections" 
                            bind:value={maxSelections}
                            min="1"
                        />
                    </Label>
                {/if}
            </div>

            <!-- Next Question Logic -->
            <div class="space-y-4">
                <h3 class="text-xl font-semibold">Next Question Logic</h3>
                <Table>
                    <TableHead>
                        <TableHeadCell>Answer</TableHeadCell>
                        <TableHeadCell>Next Question ID</TableHeadCell>
                        <TableHeadCell class="w-20">Delete</TableHeadCell>
                    </TableHead>
                    <TableBody>
                        {#each question.nextQuestion?.conditions || [] as condition, index}
                            <TableBodyRow>
                                <TableBodyCell>
                                    <Select 
                                        items={question.options.map(opt => ({
                                            value: opt.text, 
                                            name: opt.text
                                        }))} 
                                        bind:value={condition.answer}
                                        placeholder="Select Answer" 
                                        required 
                                        class="w-full"
                                    />
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Input 
                                        type="number" 
                                        placeholder="Next Question ID" 
                                        bind:value={condition.nextId}
                                        min="1"
                                        required 
                                        class="w-full"
                                    />
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Button class="w-12 h-12" outline color="red" on:click={() => handleRemoveCondition(index)}>
                                        <TrashBinSolid/>
                                    </Button>
                                </TableBodyCell>
                            </TableBodyRow>
                        {/each}
                    </TableBody>
                </Table>
                
                
                <Button color="light" on:click={handleAddCondition}>
                    <CirclePlusSolid class="mr-2" /> Add Condition
                </Button>
            </div>
        {/if}

        <!-- Input Question Specific Fields -->
        {#if question.type === QuestionType.INPUT}
            <div class="space-y-4">
                <Label>
                    Placeholder
                    <Input 
                        type="text" 
                        placeholder="Enter placeholder text" 
                        bind:value={placeholder}
                    />
                </Label>

                <Label>
                    Input Type
                    <Select 
                        items={inputTypes}
                        bind:value={inputType}
                    />
                </Label>

                <Label>
                    Validation Type
                    <Select 
                        class="mt-2"
                        bind:value={validation}
                    >
                        {#each validationOptions as option}
                            <option value={option.value} title={option.pattern}>
                                {option.label}
                            </option>
                        {/each}
                    </Select>
                </Label>
            </div>
        {/if}

        {#if question.type === QuestionType.RATING}
            <div class="space-y-4">
                <div class="grid grid-cols-2 gap-4">
                    <Label>
                        Minimum Value
                        <Input 
                            type="number" 
                            bind:value={question.minValue}
                            min="1"
                        />
                    </Label>
                    <Label>
                        Maximum Value
                        <Input 
                            type="number" 
                            bind:value={question.maxValue}
                            min={question.minValue + 1}
                        />
                    </Label>
                </div>

                <Label>
                    Step
                    <Input 
                        type="number" 
                        bind:value={question.step}
                        min="0.1"
                        step="0.1"
                    />
                </Label>

                <div class="grid grid-cols-2 gap-4">
                    <Label>
                        Minimum Label
                        <Input 
                            type="text" 
                            bind:value={question.minLabel}
                            placeholder="e.g., Poor"
                        />
                    </Label>
                    <Label>
                        Maximum Label
                        <Input 
                            type="text" 
                            bind:value={question.maxLabel}
                            placeholder="e.g., Excellent"
                        />
                    </Label>
                </div>

                <Label>
                    Rating Icon
                    <Input 
                        type="text" 
                        bind:value={question.icon}
                        placeholder="e.g., ⭐️"
                    />
                </Label>

                <Label class="flex items-center gap-2">
                    <Checkbox 
                        type="checkbox" 
                        bind:checked={question.showLabels}
                        />
                    Show Labels
                </Label>
            </div>
        {/if}

        <Label>
            Default Next Question ID
            <Input 
                type="number" 
                placeholder="Enter default next question ID" 
                bind:value={defaultNextId}
                min="1"
            />
        </Label>
    </div>

    <Button 
        color="blue" 
        class="mt-6 w-full" 
        on:click={applyChanges}
    >
        Apply Changes
    </Button>
</div>

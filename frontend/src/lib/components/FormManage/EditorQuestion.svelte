<script lang="ts">
    import { QuestionType } from '$lib/types';
    import type { Question, SingleChoiceQuestion, MultipleChoiceQuestion, InputQuestion } from '$lib/types';
    import { Label, Input, Button, Select } from 'flowbite-svelte';
    import { CirclePlusSolid } from 'flowbite-svelte-icons';

    interface Props {
        question: Question;
        updateQuestionCallback: (updatedQuestion: Question) => void;
    }

    let { question, updateQuestionCallback }: Props = $props();

    // Deep clone function to create a completely new object
    function deepClone<T>(obj: T): T {
        return JSON.parse(JSON.stringify(obj));
    }

    // State for managing question properties
    let questionText = $state(deepClone(question).question);
    let subtext = $state(deepClone(question).subtext);
    let image = $state(deepClone(question).image);
    let nextQuestionConditions = $state(deepClone(question).nextQuestion?.conditions || []);
    let defaultNextQuestion = $state(deepClone(question).nextQuestion?.default);

    // State for SingleChoice and MultipleChoice questions
    let options = $state(
        'options' in question 
            ? deepClone(question).options.map(opt => ({...opt})) 
            : []
    );
    let maxSelections = $state(
        'maxSelections' in question 
            ? deepClone(question).maxSelections 
            : undefined
    );

    // State for Input questions
    let placeholder = $state(
        'placeholder' in question 
            ? deepClone(question).placeholder 
            : ''
    );
    let inputType = $state(
        'inputType' in question 
            ? deepClone(question).inputType 
            : 'text'
    );
    let validation = $state(
        'validation' in question 
            ? deepClone(question).validation 
            : undefined
    );

    // Function to apply all changes
    function applyChanges() {
        const updatedQuestion = {
            ...question,
            question: questionText,
            subtext,
            image,
            nextQuestion: {
                conditions: nextQuestionConditions,
                default: defaultNextQuestion
            }
        };

        if ('options' in question) {
            (updatedQuestion as SingleChoiceQuestion | MultipleChoiceQuestion).options = options;
        }

        if ('maxSelections' in question) {
            (updatedQuestion as MultipleChoiceQuestion).maxSelections = maxSelections;
        }

        if ('placeholder' in question) {
            (updatedQuestion as InputQuestion).placeholder = placeholder;
        }

        if ('inputType' in question) {
            (updatedQuestion as InputQuestion).inputType = inputType;
        }

        if ('validation' in question) {
            (updatedQuestion as InputQuestion).validation = validation;
        }

        updateQuestionCallback(updatedQuestion);
    }
</script>

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

        <Label>
            Image URL
            <Input 
                type="url" 
                placeholder="Enter image URL" 
                bind:value={image} 
            />
        </Label>

        <!-- Options for Single/Multiple Choice -->
        {#if question.type === QuestionType.SINGLE_CHOICE || question.type === QuestionType.MULTIPLE_CHOICE}
            <div class="space-y-4">
                <h3 class="text-xl font-semibold">Options</h3>
                {#each options as option}
                    <Input 
                        type="text" 
                        placeholder="Enter option text" 
                        bind:value={option.text} 
                        required 
                    />
                {/each}
                <Button 
                    color="light" 
                    on:click={() => options = [...options, { text: '', icon: '' }]}
                >
                    <CirclePlusSolid class="mr-2" /> Add Option
                </Button>

                {#if question.type === QuestionType.MULTIPLE_CHOICE}
                    <Label>
                        Max Selections
                        <Input 
                            type="number" 
                            placeholder="Enter max selections" 
                            bind:value={maxSelections} 
                        />
                    </Label>
                {/if}
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
                        bind:value={inputType}
                    >
                        <option value="text">Text</option>
                        <option value="email">Email</option>
                        <option value="number">Number</option>
                    </Select>
                </Label>

                <Label>
                    Validation Regex
                    <Input 
                        type="text" 
                        placeholder="Enter validation regex" 
                        bind:value={validation} 
                    />
                </Label>
            </div>
        {/if}

        <!-- Next Question Logic -->
        <div class="space-y-4">
            <h3 class="text-xl font-semibold">Next Question Logic</h3>
            {#each nextQuestionConditions as condition}
                <div class="flex space-x-4">
                    <Input 
                        type="text" 
                        placeholder="Answer" 
                        bind:value={condition.answer} 
                        required 
                    />
                    <Input 
                        type="number" 
                        placeholder="Next Question ID" 
                        bind:value={condition.nextId} 
                        required 
                    />
                </div>
            {/each}
            <Button 
                color="light" 
                on:click={() => nextQuestionConditions = [...nextQuestionConditions, { answer: '', nextId: 0 }]}
            >
                <CirclePlusSolid class="mr-2" /> Add Condition
            </Button>
        </div>

        <Label>
            Default Next Question ID
            <Input 
                type="number" 
                placeholder="Enter default next question ID" 
                bind:value={defaultNextQuestion} 
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

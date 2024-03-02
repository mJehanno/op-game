<script setup lang="ts">
import InputNumber from 'primevue/inputnumber';
import FloatLabel from 'primevue/floatlabel';

import Prompter from '@/components/Prompter.vue';

const props = defineProps<{
    prompt: string;
    expected: number
}>()

const emit = defineEmits<{
    succeed: []
    failed: []
}>()




function handleInput(event: Event) {
    const v = (event.target as HTMLInputElement).value;
    check(parseInt(v));
    //(event.target as HTMLInputElement).value = null;
}

function check(val: number | undefined){
    if (val === props.expected) {
        emit("succeed")
    }else {
        emit("failed")
    }
}


</script>
<template>
    <div>
        <div class="flex flex-column align-items-center">
            <div class="flex flex-row align-items-center">
                <div id="game">
                    <Prompter id="prompt" :prompt="props.prompt" />
                    <FloatLabel id="input" >
                        <InputNumber :input-props="{'autofocus': true}"  id="answer-input"   @keyup.enter="handleInput($event); $event.target.value=null;" />
                        <label for="answer-input">Answer</label>
                    </FloatLabel>
                </div>
            </div>
        </div>

    </div>
</template>
<style>
    #game{
        margin: 2em;
        display: flex;
        flex-direction: row;
        align-items:center ;
        justify-content: space-evenly;
    }
    #prompt{
        font-size: 5em;
    }
    .p-dialog-header, .p-dialog-footer{
        padding: 1em;
    }
</style>


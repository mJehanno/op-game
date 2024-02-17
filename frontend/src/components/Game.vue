<script setup lang="ts">
import { computed, ref} from 'vue';
import { onMounted } from 'vue'

import InputNumber from 'primevue/inputnumber';
import FloatLabel from 'primevue/floatlabel';
import Dialog from 'primevue/dialog';
import Prompter from '@/components/Prompter.vue';
import Timer from '@/components/Timer.vue';



onMounted(() => {
    generatePrompt();
    //answer.value?.focus();
});

let value = ref<number | null>(null) ;
const emits = defineEmits(["gameOver"])
const x = ref(0);
const y = ref(0);
const streak = ref(0);
const prompt = computed<string>(() => x.value + " X " + y.value + " = ");
let result = ref(0);
const visible = ref(false);
const seconds = ref(20);
const timer = ref<HTMLInputElement | null>(null)

function generateRandom(): number{
    return Math.ceil(Math.random() * 10);
}

function openEndingDialog() {
    visible.value =true;
}

function gameOver(){
    emits("gameOver");
}
function generatePrompt() {
    x.value = generateRandom();
    y.value = generateRandom();
}
function check(val: number | null)  {
    result = computed(() => x.value * y.value);
    if (val != (x.value*y.value)){
        openEndingDialog()
    }else {
        streak.value ++;
        if (streak.value %5 == 0 && seconds.value > 4) {
            seconds.value -= 4;
        }
        value.value=null;
        generatePrompt()
        timer.value?.reset(seconds.value)
        console.log(value.value)
    }
}
</script>
<template>
    <div>
    <Timer @timeout="check(value)" :sec="seconds" ref="timer"/>
    <span id="streak">Current streak : {{ streak }}</span>
    <div id="game">
        <Prompter id="prompt" :prompt="prompt" />
        <FloatLabel id="input" >
            <InputNumber :model-value="value" v-focustrap :input-props="{'autofocus': true}"  id="answer-input"   @keyup.enter="check($event.target.value); $event.target.value = null;" />
            <label for="answer-input">Answer</label>
        </FloatLabel>
    </div>
    </div>

    <Dialog v-model:visible="visible" modal header="Game Over" v-on:hide="gameOver">
        <div class="flex flex-column">
            <span class="p-text-secondary block ">The answer to {{ prompt }} was : {{ result }}</span>
            <Button label="Back" class="m-2" severity="success" @click="gameOver"/>
        </div>
    </Dialog>
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
    #streak{
        font-size: 3em;
    }
</style>


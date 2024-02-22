<script setup lang="ts">
import { computed, ref} from 'vue';
import { onMounted } from 'vue'

import InputNumber from 'primevue/inputnumber';
import FloatLabel from 'primevue/floatlabel';
import Dialog from 'primevue/dialog';
import InlineMessage from 'primevue/inlinemessage';

import Prompter from '@/components/Prompter.vue';
import Timer from '@/components/Timer.vue';
import { useRouter } from 'vue-router';
import { AddScore } from '../../wailsjs/go/score/ScoreService';



onMounted(() => {
    generatePrompt();
});

let user = ref<string | null>(null);
let value = ref<number | null>(null);
const router = useRouter();
const x = ref(0);
const y = ref(0);
const streak = ref(0);
const prompt = computed<string>(() => x.value + " X " + y.value + " = ");
let result = ref(0);
const visible = ref(false);
const seconds = ref(20);
const timer = ref<HTMLInputElement | null>(null)
const err = ref<string| null>(null)

function generateRandom(): number{
    return Math.ceil(Math.random() * 10);
}

function openEndingDialog() {
    visible.value =true;
}

function gameOver(){
    if (user.value?.length > 0 && user.value?.length <=3) {
        AddScore({username: user.value, score: streak.value})
        router.push('/');
    }else if (user.value?.length > 3){
        err.value= "username should contains max 3 characters"
    }else {
        err.value = "username is required"
    }
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
        generatePrompt()
        timer.value?.reset(seconds.value)
    }
}
</script>
<template>
    <div class="flex flex-column align-items-center">
    <Timer @timeout="check(value)" :sec="seconds" ref="timer"/>
    <span id="streak">Current streak : {{ streak }}</span>
    <div id="game">
        <Prompter id="prompt" :prompt="prompt" />
        <FloatLabel id="input" >
            <InputNumber :model-value="value" :input-props="{'autofocus': true}"  id="answer-input"   @keyup.enter="check($event.target.value); $event.target.value = null;" />
            <label for="answer-input">Answer</label>
        </FloatLabel>
    </div>
    </div>

    <Dialog v-model:visible="visible" modal header="Game Over" :closable="false" v-on:hide="gameOver">
        <div class="flex flex-column">
            <span class="p-text-secondary block ">The answer to {{ prompt }} was : {{ result }}</span>
            <div class="flex flex-column m-3 flex-wrap">
                <span class="m-1">Register your score !</span>
                <FloatLabel>
                    <InputText id="username" v-model="user" mask="aaa" placeholder="aaa" aria-describedby="username-help"/>
                    <label for="username">Username</label>
                </FloatLabel>
                <small id="username-help">Can't contain more than 3 characters.</small>
                <InlineMessage class="m-1" v-if="err" severity="error">{{ err }}</InlineMessage>
            </div>
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


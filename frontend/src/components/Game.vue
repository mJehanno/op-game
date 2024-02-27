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
import { DifficultyLevel } from '@/models/game';
import {score} from '../../wailsjs/go/models';
import { useRoute } from 'vue-router';

onMounted(() => {
    generatePrompt();
});

const route = useRoute()
let user = ref<string | null>(null);
let value = ref<number | null>(null);
const router = useRouter();
const x = ref(0);
const y = ref(0);
const streak = ref(0);
const prompt = computed<string>(() => x.value + " X " + y.value + " = ");
let result = ref(0);
const visible = ref(false);
const timer = ref<typeof Timer | null>(null)
const err = ref<string| null>(null)
const level = ref<DifficultyLevel>(route.params.difficultyLevel as DifficultyLevel)
const currentlife = ref(3);


function generateRandom(): number{
    return Math.ceil(Math.random() * 10);
}

function openEndingDialog() {
    visible.value =true;
}

function gameOver(){
    if (user.value && user.value?.length > 0 && user.value?.length <=3) {
        AddScore(score.Score.createFrom({username: user.value, score: streak.value}))
        router.push('/');
    }else if (user.value && user.value?.length > 3){
        err.value= "username should contains max 3 characters"
    }else {
        err.value = "username is required"
    }
}
function generatePrompt() {
    x.value = generateRandom();
    y.value = generateRandom();
}


// result ok, pas timeout -> continue
// result ok, timemout -> increase streak, but end
// result pas ok -> depends on lvl
// if timeout -> end
function check(val: number | null, timeout: boolean)  {
    result = computed(() => x.value * y.value);

    if (val == result.value) {
        streak.value ++;
        if (level.value == DifficultyLevel.Hard && timer.value && streak.value %5 == 0 && timer.value?.defaultSec.value > 4) {
            timer.value.defaultSec.value -= 4;
        }
    }else {
        switch(level.value) {
        case DifficultyLevel.Easy:
            if (!timeout) {
                return;
            }
            break;
        case DifficultyLevel.Medium:
            currentlife.value --;
            if (currentlife.value == 0) {
                openEndingDialog()
            }
            if (!timeout) {
                return;
            }
            break;
        case DifficultyLevel.Hard:
            openEndingDialog();
            return;
        }
    }

    if (timeout) {
        openEndingDialog();
        return;
    }
    generatePrompt();
    timer.value?.reset();
}
</script>
<template>
    <div class="flex flex-column align-items-center">
    <Timer @timeout="check(value, true)" ref="timer"/>
    <div class="flex flex-row align-items-center w-full justify-content-evenly">
        <span id="streak">Current streak : {{ streak }}</span>
        <div v-if="level == DifficultyLevel.Medium">
            <i class="pi pi-heart-fill text-red-700 text-2xl" v-for="n in currentlife"></i>
            <i class="pi pi-heart text-red-700 text-2xl" v-for="n in (3- currentlife)"></i>
        </div>
    </div>

    <div id="game">
        <Prompter id="prompt" :prompt="prompt" />
        <FloatLabel id="input" >
            <InputNumber :model-value="value" :input-props="{'autofocus': true}"  id="answer-input"   @keyup.enter="check($event.target.value, false); $event.target.value = null;" />
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
            <Button label="Submit" class="m-2" severity="success" @click="gameOver"/>
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


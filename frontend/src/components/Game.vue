<script setup lang="ts">
import { computed, reactive, ref} from 'vue';
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
import type {GamePrompt, GameState} from '@/models/game';
import {score} from '../../wailsjs/go/models';
import { useRoute } from 'vue-router';

onMounted(() => {
    generatePrompt();
});

const route = useRoute()
const router = useRouter();

const prompt = computed<string>(() => game_prompt.x + " X " + game_prompt.y + " = ");
const result = computed<number>(() => game_prompt.x * game_prompt.y);
const timer = ref<typeof Timer | null>(null)
const gameState: GameState  = reactive({
    streak: 0,
    endingDialogVisible : false,
    level: route.params.difficultyLevel as DifficultyLevel,
    currentlife: route.params.difficultyLevel == 'medium' ? 3 : 0,
})

const game_prompt: GamePrompt = reactive({
    x : 0,
    y : 0,
    result: result,
    prompt: prompt,
})

function generateRandom(): number{
    return Math.ceil(Math.random() * 10);
}

function openEndingDialog() {
    gameState.endingDialogVisible = true;
}

function gameOver(){
    if (gameState.user && gameState.user?.length > 0 && gameState.user?.length <=3) {
        AddScore(score.Score.createFrom({username: gameState.user, score: gameState.streak, difficulty: gameState.level}))
        router.push('/');
    }else if (gameState.user && gameState.user?.length > 3){
        gameState.err = "username should contains max 3 characters"
    }else {
        gameState.err = "username is required"
    }
}
function generatePrompt() {
    game_prompt.x = generateRandom();
    game_prompt.y = generateRandom();
}


function check(val: number | undefined, timeout: boolean)  {
    if (val == game_prompt.result) {
        gameState.streak ++;
        if (gameState.level == DifficultyLevel.Hard && timer.value && gameState.streak %5 == 0 && timer.value?.defaultSec.value > 4) {
            timer.value.defaultSec.value -= 4;
        }
    }else {
        switch(gameState.level) {
        case DifficultyLevel.Easy:
            if (!timeout) {
                return;
            }
            break;
        case DifficultyLevel.Medium:
            gameState.currentlife --;
            if (gameState.currentlife == 0) {
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
    <Timer @timeout="check(gameState.answer, true)" ref="timer"/>
    <div class="flex flex-row align-items-center w-full justify-content-evenly">
        <span id="streak">Current streak : {{ gameState.streak }}</span>
        <div v-if="gameState.level == DifficultyLevel.Medium">
            <i class="pi pi-heart-fill text-red-700 text-2xl" v-for="n in gameState.currentlife"></i>
            <i class="pi pi-heart text-red-700 text-2xl" v-for="n in (3- gameState.currentlife)"></i>
        </div>
    </div>

    <div id="game">
        <Prompter id="prompt" :prompt="game_prompt.prompt" />
        <FloatLabel id="input" >
            <InputNumber :model-value="gameState.answer" :input-props="{'autofocus': true}"  id="answer-input"   @keyup.enter="check($event.target.value, false); $event.target.value = null;" />
            <label for="answer-input">Answer</label>
        </FloatLabel>
    </div>
    </div>

    <Dialog v-model:visible="gameState.endingDialogVisible" modal header="Game Over" :closable="false" v-on:hide="gameOver">
        <div class="flex flex-column">
            <span class="p-text-secondary block ">The answer to {{ game_prompt.prompt }} was : {{ game_prompt.result }}</span>
            <div class="flex flex-column m-3 flex-wrap">
                <span class="m-1">Register your score !</span>
                <FloatLabel>
                    <InputText id="username" v-model="gameState.user" mask="aaa" placeholder="aaa" aria-describedby="username-help"/>
                    <label for="username">Username</label>
                </FloatLabel>
                <small id="username-help">Can't contain more than 3 characters.</small>
                <InlineMessage class="m-1" v-if="gameState.err" severity="error">{{ gameState.err }}</InlineMessage>
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


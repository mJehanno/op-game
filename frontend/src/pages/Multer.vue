<script setup lang="ts">
import Timer from '@/components/Timer.vue';
import Streak from '@/components/Streak.vue';
import ScoreBoard from '@/components/ScoreBoard.vue';
import GameOverDialog from '@/components/GameOverDialog.vue';
import LifeBar from '@/components/LifeBar.vue';
import Game from '@/components/Game.vue';
import { computed, onMounted, reactive, ref } from 'vue';
import {score} from '../../wailsjs/go/models';
import { DifficultyLevel, type GameState } from '@/models/game';
import { useRoute, useRouter } from 'vue-router';
import {Mode} from '@/models/scoreboard';
import { AddScore, GetScore } from '../../wailsjs/go/score/ScoreService';
import type {GamePrompt} from '@/models/game';

const route = useRoute();
const router = useRouter();

const gameState = reactive<GameState>({
    maxSec: 20,
    endingDialogVisible: false,
    streak: 0,
    currentLife: 3,
    maxLife: 3,
});

const prompt = computed<string>(() => gamePrompt.x + " X " + gamePrompt.y + " = ");
const result = computed<number>(() => gamePrompt.x * gamePrompt.y);

const gamePrompt: GamePrompt = reactive({
    x : 0,
    y : 0,
    result: result,
    prompt: prompt,
})


const scores = ref<score.Score[]>([])
const timer = ref<typeof Timer | null>(null)

const current: score.Score = reactive(
    {id: -1,username: 'xxx', score: 0, difficulty: route.params.level, created_at: new Date()} as score.Score
)

function generateRandom(): number{
    return Math.ceil(Math.random() * 10);
}

function generatePrompt() {
    gamePrompt.x = generateRandom();
    gamePrompt.y = generateRandom();
}

function checkDifficulty(lvl: DifficultyLevel): boolean {
    return route.params.level == lvl;
}

function endGame() {
    gameState.endingDialogVisible = true;
}

function handleSucceed() {
    gameState.streak ++;
    current.score ++;
    if (checkDifficulty(DifficultyLevel.Hard) && gameState.streak %5 === 0 && gameState.maxSec > 4) {
        gameState.maxSec -= 4;
    }
    generatePrompt();
    timer.value.reset()
}

function handleFailed() {
    if (checkDifficulty(DifficultyLevel.Hard)) {
        endGame();
        return;
    }
    if (checkDifficulty(DifficultyLevel.Medium)) {
        gameState.currentLife --;
        if (gameState.currentLife === 0) {
            endGame();
        }
    }
}


function quitGame(username: string) {
    AddScore(score.Score.createFrom({username: username, score: gameState.streak, difficulty: route.params.level}));
    router.push('/');
}

onMounted(()=> {

    generatePrompt();

    GetScore(route.params.level).then((scs: score.Score[]) => {
        scs.push(current);
        scores.value = scs;
    }).catch(err => {

    })
})

</script>
<template>
    <div class="flex flex-column align-items-center m-2">
        <Timer :seconds="gameState.maxSec" @timeout="endGame" ref="timer"/>
        <div class="flex flex-row justify-content-around w-full align-items-center m-3">
            <Streak class="text-3xl" :streak="gameState.streak" />
            <LifeBar v-if="checkDifficulty(DifficultyLevel.Medium)" :max-health="gameState.maxLife" :current="gameState.currentLife" />
        </div>
        <div class="flex flex-row align-content-center justify-content-evenly w-full m-3">
            <ScoreBoard size="small" :scores="scores" :mode="Mode.Live" />
            <Game :prompt="gamePrompt.prompt" :expected="gamePrompt.result" @succeed="handleSucceed" @failed="handleFailed" />
        </div>
        <GameOverDialog :visible="gameState.endingDialogVisible" :prompt="gamePrompt.prompt" :answer="gamePrompt.result" @user-created="quitGame" /> 
    </div>
</template>
<script setup lang="ts">
import { DifficultyLevel, type GameState } from '@/models/game';
import { computed, onBeforeMount, onMounted, reactive, ref } from 'vue';
import Timer from '@/components/game-components/Timer.vue';
import Streak from '@/components/game-components/Streak.vue';
import LifeBar from '@/components/game-components/LifeBar.vue';
import ScoreBoard from '@/components/ScoreBoard.vue';
import GameBoard from '@/components/game-components/GameBoard.vue';
import GameOverDialog from '@/components/game-components/GameOverDialog.vue';
import { Mode } from '@/models/scoreboard';
import { useGameInfosStore } from '@/stores/game-infos';
import { useScoresStore } from '@/stores/scores';
import { useGameStore } from '@/stores/game';
import { useRouter } from 'vue-router';

const router = useRouter();
const gameInfosStore = useGameInfosStore();
const scoresStore = useScoresStore();
const gameStore = useGameStore();

const gameState = reactive<GameState>({
    maxSec: 20,
    endingDialogVisible: false,
    currentLife: 3,
    maxLife: 3,
});
const timer = ref<typeof Timer | null>(null);

function checkDifficulty(lvl: DifficultyLevel): boolean {
    return gameInfosStore.difficultyLevel  == lvl;
}

const scores = computed(() => {
    const arr = scoresStore.scores.get(gameInfosStore.selectedGame)?.get(gameInfosStore.difficultyLevel);
    arr?.push(scoresStore.current);
    return arr;
})

function endGame() {
    gameState.endingDialogVisible = true;
}

function handleSucceed() {
    scoresStore.increaseCurrent();
    if (checkDifficulty(DifficultyLevel.Hard) && scoresStore.current.score %5 === 0 && gameState.maxSec > 4) {
        gameState.maxSec -= 4;
    }
    //gameStore.generatePrompt();
    gameStore.generate()
    timer.value?.reset();
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

onBeforeMount(() => {
    scoresStore.getScoreByGameAndLevel(gameInfosStore.selectedGame, gameInfosStore.difficultyLevel).then();
})

onMounted(() => {
    gameStore.generate();
})

function quitGame(username: string) {
    scoresStore.setCurrentScoreUsername(username)
    scoresStore.saveScore();
    router.push('/');
}
</script>
<template>
    <div class="flex flex-column align-items-center m-2">
        <Timer :seconds="gameState.maxSec" @timeout="endGame" ref="timer"/>
        <div class="flex flex-row justify-content-around w-full align-items-center m-3">
            <Streak class="text-3xl" />
            <LifeBar v-if="checkDifficulty(DifficultyLevel.Medium)" :max-health="gameState.maxLife" :current="gameState.currentLife" />
        </div>
        <div class="flex flex-row align-content-center justify-content-evenly w-full m-3">
            <ScoreBoard size="small" :scores="scores" :mode="Mode.Live" />
            <GameBoard :prompt="gameStore.game.prompt" :expected="gameStore.game.result" @succeed="handleSucceed" @failed="handleFailed" />
        </div>
        <GameOverDialog :visible="gameState.endingDialogVisible" :prompt="gameStore.game.prompt" :answer="gameStore.game.result" @user-created="quitGame" /> 
    </div>
</template>


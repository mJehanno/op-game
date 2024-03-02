import { GameMode, type DifficultyLevel } from "@/models/game";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useGameInfosStore  = defineStore("game-info", ()=> {
    const selectedGame = ref(GameMode.Mult);
    const difficultyLevel = ref();
    function changeSelectedGame(newGame: GameMode) {
        selectedGame.value = newGame;
    }

    function changeDifficultyLevel(level: DifficultyLevel) {
        difficultyLevel.value = level;
    }

    return {selectedGame, difficultyLevel, changeSelectedGame, changeDifficultyLevel}
})

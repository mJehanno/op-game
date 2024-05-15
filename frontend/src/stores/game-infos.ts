import { GameMode, type DifficultyLevel } from "@/models/game";
import { defineStore } from "pinia";
import { ref } from "vue";
import { useScoresStore } from "./scores";

export const useGameInfosStore  = defineStore("game-info", ()=> {
    const score = useScoresStore()
    const selectedGame = ref(GameMode.Mult);
    changeSelectedGame(selectedGame.value)
    const difficultyLevel = ref();
    function changeSelectedGame(newGame: GameMode) {
        selectedGame.value = newGame;
        score.setGame(newGame);
    }

    function changeDifficultyLevel(level: DifficultyLevel) {
        difficultyLevel.value = level;
        score.setDifficulty( level);
    }

    return {selectedGame, difficultyLevel, changeSelectedGame, changeDifficultyLevel}
})

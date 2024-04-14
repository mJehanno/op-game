import { defineStore } from "pinia";
import { useGameInfosStore } from "./game-infos";
import { computed, reactive, ref } from "vue";
import { GameMode } from "@/models/game";

function generateRandom(max: number): number{
    return Math.ceil(Math.random() * max);
}


export const useGameStore = defineStore('game', () => {
    const gameInfosStore = useGameInfosStore();
    const game = reactive({x: 0, y: 0, prompt: "",  result: 0})
    
    function generate() {
        switch(gameInfosStore.selectedGame) {
            case GameMode.Divid:
                game.y = generateRandom(10);
                game.x = generateRandom(10) * game.y;
                game.prompt = game.x + " / " + game.y;
                game.result = game.x / game.y;
                break;
            case GameMode.Mult:
                game.x = generateRandom(10);
                game.y = generateRandom(10);
                game.prompt = game.x + " X " + game.y;
                game.result = game.x * game.y;
            default:
        }
    }
    


    return { game, generate}
})
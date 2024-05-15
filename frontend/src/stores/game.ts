import { defineStore } from "pinia";
import { useGameInfosStore } from "./game-infos";
import { reactive } from "vue";
import { GameMode } from "@/models/game";

function generateRandom(min: number, max: number): number{
    return Math.ceil(Math.random() * (max-min) +min);
}


export const useGameStore = defineStore('game', () => {
    const gameInfosStore = useGameInfosStore();
    const game = reactive({x: 0, y: 0, prompt: "",  result: 0})
    
    function generate() {
        game.y = generateRandom(0,10);
        game.x = generateRandom(0,10);
        switch(gameInfosStore.selectedGame) {
            case GameMode.Divid:
                game.x *= game.y;
                game.prompt = game.x + " / " + game.y;
                game.result = game.x / game.y;
                break;
            case GameMode.Add:
                const coin = generateRandom(0,1); 
                if (coin === 1) {
                    game.x = generateRandom(5,10);
                }else {
                    game.y = generateRandom(5,10);
                }
                game.prompt = game.x + " + " + game.y;
                game.result = game.x + game.y;
                break;
            case GameMode.Min:
                game.x = generateRandom(5, 15);
                game.x = generateRandom(game.y, 20);
                game.prompt = game.x + " - " + game.y;
                game.result = game.x - game.y;
                break;
            case GameMode.Mult:
            default:
                game.prompt = game.x + " X " + game.y;
                game.result = game.x * game.y;
        }
        game.prompt += " = ";
    }
    


    return { game, generate}
})
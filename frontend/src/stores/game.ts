import { defineStore } from "pinia";
import { useGameInfosStore } from "./game-infos";
import { computed, ref } from "vue";
import { GameMode } from "@/models/game";

function generateRandom(max: number): number{
    return Math.ceil(Math.random() * max);
}


export const useGameStore = defineStore('game', () => {
    const gameInfosStore = useGameInfosStore();
    const x = ref(0);
    const y = ref(0);


    const computedX  = computed<number>(() => {
            switch(gameInfosStore.selectedGame) {
                case GameMode.Divid:
                    x.value = generateRandom(10) * computedY.value;
                    return x.value;
                case GameMode.Mult:
                    x.value = generateRandom(10);
                    return x.value;
                default:
                    return 0;
            }
    });

    const computedY = computed<number>(() => {
            switch(gameInfosStore.selectedGame) {
                case GameMode.Mult:
                case GameMode.Divid:
                    y.value = generateRandom(10);
                    return y.value;
                default:
                    return 0;
            }
    });


    function generatePrompt() {
            switch(gameInfosStore.selectedGame) {
                case GameMode.Mult:
                    x.value = generateRandom(10);
                    y.value = generateRandom(10);
                case GameMode.Divid:
                    y.value = generateRandom(10);
                    x.value = generateRandom(10) * y.value;
                default:
                    return 0;
            }
    }

    const prompt = computed(() => {
        switch(gameInfosStore.selectedGame) {
            case GameMode.Divid:
                return computedX.value + " / " + computedY.value
            case GameMode.Mult:
            default:
                return computedX.value + " X " + computedY.value;
            
        }
    })

    const result = computed(() => {
        switch(gameInfosStore.selectedGame) {
            case GameMode.Divid:
                return computedX.value / computedY.value;
            case GameMode.Mult:
            default:
                return computedX.value * computedY.value;
        }
    })

    return {x,y, computedX, computedY, prompt, result, generatePrompt}
})
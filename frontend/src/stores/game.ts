import { defineStore } from "pinia";
import { useGameInfosStore } from "./game-infos";
import { computed, ref } from "vue";

function generateRandom(max: number): number{
    return Math.ceil(Math.random() * max);
}


export const useGameStore = defineStore('game', () => {
    const gameInfosStore = useGameInfosStore();
    const x = ref(0);
    const y = ref(0);


    const computedX  = computed<number>(() => {
            switch(gameInfosStore.selectedGame) {
                case "mult":
                    x.value = generateRandom(10);
                    return x.value;
                default:
                    return 0;
            }
    });

    const computedY = computed<number>(() => {
            switch(gameInfosStore.selectedGame) {
                case "mult":
                    y.value = generateRandom(10);
                    return y.value;
                default:
                    return 0;
            }
    });


    function generatePrompt() {
            switch(gameInfosStore.selectedGame) {
                case "mult":
                    x.value = generateRandom(10);
                    y.value = generateRandom(10);
                default:
                    return 0;
            }
    }

    const prompt = computed(() => {
        switch(gameInfosStore.selectedGame) {
            case "mult":
            default:
                return computedX.value + " X " + computedY.value;
        }
    })

    const result = computed(() => {
        switch(gameInfosStore.selectedGame) {
            case "mult":
            default:
                return computedX.value * computedY.value;
        }
    })

    return {x,y, computedX, computedY, prompt, result, generatePrompt}
})
import { defineStore } from "pinia";
import { useGameInfosStore } from "./game-infos";
import { computed, ref } from "vue";



function generateRandom(): number{
    return Math.ceil(Math.random() * 10);
}


export const useGameStore = defineStore('game', () => {
    const gameInfosStore = useGameInfosStore();
    const x = ref(generateRandom());
    const y = ref(generateRandom());

    
    function generatePrompt() {
        x.value = generateRandom();
        y.value = generateRandom();
    }

    const prompt = computed(() => {
        switch(gameInfosStore.selectedGame) {
            case "mult":
            default:
                return x.value + " X " + y.value;
        }
    })

    const result = computed(() => {
        switch(gameInfosStore.selectedGame) {
            case "mult":
            default:
                return x.value * y.value;
        }
    })

    return {x,y, prompt, result, generatePrompt}
})
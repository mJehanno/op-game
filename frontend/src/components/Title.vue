<script setup lang="ts">
import { useGameInfosStore } from '@/stores/game-infos';
import {computed} from 'vue'
const store = useGameInfosStore()


const entitled = computed(() =>  capitalize(store.selectedGame))

function capitalize(str: string):string{
    const init = str.charAt(0)
    return init.toUpperCase() + str.slice(1)
}
</script>
<template>
    <h1>{{entitled}}-Game</h1>
</template>
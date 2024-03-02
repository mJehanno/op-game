<script setup lang="ts">
import { ref } from 'vue';
import DifficultyMenu from '@/components/menus/DifficultyMenu.vue';

const displayDifficultyLevel = ref(false);
const displayDifficulty = () => {
    displayDifficultyLevel.value = !displayDifficultyLevel.value;
}
</script>
<template>
    <div class="flex flex-column m-8 w-full align-items-center">
        <Button class="m-3 w-3" label="Start Game" rounded severity="success" @click="displayDifficulty" />
        <DifficultyMenu class="flex flex-column w-full" v-if="displayDifficultyLevel"/>
        <div class="w-full m-3">
            <router-link to="/scores" class="m-3 w-full"><Button  class="w-3" label="See Scores" rounded severity="help" /></router-link>
        </div>
    </div>
</template>
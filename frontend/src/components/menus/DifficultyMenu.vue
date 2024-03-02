<script setup lang="ts">
import { DifficultyLevel } from '@/models/game';
import { useGameInfosStore } from '@/stores/game-infos';

const store = useGameInfosStore();

</script>
<template>
    <div class="m-1 w-full" v-for="(value, key) in DifficultyLevel">
        <router-link to="/game" class="m-1 w-full">
            <Button class="w-2" rounded :severity="value == DifficultyLevel.Easy ? 'info' : value == DifficultyLevel.Medium ? 'warning': 'danger'" :label="key" @click="store.changeDifficultyLevel(value)" />
        </router-link>
    </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';

import Button from 'primevue/button';

import Update from '@/components/Update.vue';
import VersionDisplayer from '@/components/VersionDisplayer.vue';


import { DifficultyLevel } from '@/models/game';

const displayDifficultyLevel = ref(false)


const displayDifficulty = () => {
    displayDifficultyLevel.value = !displayDifficultyLevel.value
}



</script>
<template>
    <div class="flex flex-column align-items-center">
        <div class="flex flex-column m-8 w-full align-items-center">
            <Button class="m-3 w-3" label="Start Game" rounded severity="success" @click="displayDifficulty" />
                <div class="flex flex-column w-full" v-if="displayDifficultyLevel">
                    <div class="m-1 w-full">
                        <router-link :to="'/game/mult/' + DifficultyLevel.Easy" class="m-1 w-full"><Button class="w-2" rounded severity="info" :label="DifficultyLevel.Easy" /></router-link>
                    </div>
                    <div class="m-1 w-full">
                        <router-link :to="'/game/mult/' + DifficultyLevel.Medium" class="m-1 w-full"><Button class="w-2" rounded severity="warning" :label="DifficultyLevel.Medium" /></router-link>
                    </div>
                    <div class="m-1 w-full">
                        <router-link :to="'/game/mult/' + DifficultyLevel.Hard" class="m-1 w-full"><Button class="w-2" rounded severity="danger" :label="DifficultyLevel.Hard"/></router-link>
                    </div>
                </div>
            <router-link to="/scores" class="m-3 w-full"><Button  class="w-3" label="See Scores" rounded severity="help" /></router-link>
        </div>

        <VersionDisplayer />
        <Update />
    </div>
</template>
<script setup lang="ts">
console.log("meh");
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { onBeforeMount } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import { Mode } from '@/models/scoreboard';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import ScoreBoard from '@/components/ScoreBoard.vue';
import  Title from '@/components/Title.vue';
import { useScoresStore } from '@/stores/scores';
import { useGameInfosStore } from '@/stores/game-infos';

const toast = useToast();

const gameInfos = useGameInfosStore()
const scoreStore = useScoresStore();

onBeforeMount(() => {
    console.log("toto")
    scoreStore.getAllScoresByGame(gameInfos.selectedGame).then(() => {
    }).catch(err => {
        toast.add({severity: 'error', summary: 'failed to load scoreboard data'})
    })
})

const router = useRouter();

</script>
<template>
    <div>
        <Title />
        <TabView class="m-4">
            <TabPanel v-for="[key, value] in scoreStore.sortedScore.get(gameInfos.selectedGame)" :key="key" :header="key">
                <div class="flex flex-column align-items-center">
                    <ScoreBoard :scores="value"  :mode="Mode.Static" />
                </div>
            </TabPanel>
        </TabView>
        <Button label="Back" class="m-4" @click="router.push('/')"></Button>
    </div>

    <Toast />
</template>
<style>
.p-tabview-nav {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}
</style>

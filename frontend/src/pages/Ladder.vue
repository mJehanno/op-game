<script setup lang="ts">
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import {  onMounted , reactive } from 'vue';
import {  useRouter } from 'vue-router';
import { GetScore } from '../../wailsjs/go/score/ScoreService';
import {score} from "../../wailsjs/go/models";
import { useToast } from 'primevue/usetoast';
import { Mode } from '@/models/scoreboard';
import { forIn } from 'lodash'; 
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import { DifficultyLevel } from '@/models/game';
import ScoreBoard from '@/components/ScoreBoard.vue';

const displayedScore = reactive(new Map<string, score.Score[]>([
    ["easy", []],
    ["medium", []],
    ["hard", []]
]))
const toast = useToast();


onMounted(() => {
    forIn(DifficultyLevel, (value: string, key: string) => {
        GetScore(value).then( (scores: score.Score[]) => {
            if (!scores) {
                scores = [];
            }
            displayedScore.set(value,scores);
        }).catch(err => {
            console.log(err)
            toast.add({severity: 'error', summary: 'failed to load scoreboard data'})
        })
    }) 
})

const router = useRouter();

</script>
<template>
    <div>
        <TabView class="m-4">
            <TabPanel v-for="[key, value] in displayedScore" :key="key" :header="key">
                <div class="flex flex-column align-items-center">
                    <ScoreBoard :scores="value"  :mode="Mode.Static" />
                    <Button label="Back" class="m-4" @click="router.push('/')"></Button>
                </div>
            </TabPanel>
        </TabView>
    </div>

    <Toast />
</template>
<style>
.p-tabview-nav {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}

@media only screen and (max-width: 1600px) {
    #livemode {
        display: none;
    }
}
</style>
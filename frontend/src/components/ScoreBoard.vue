<script setup lang="ts">
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { onMounted , reactive} from 'vue';
import {  useRouter } from 'vue-router';
import { GetScore } from '../../wailsjs/go/score/ScoreService';
import {score} from "../../wailsjs/go/models";
import { useToast } from 'primevue/usetoast';
import { format} from "date-fns";
import type { Mode } from '@/models/scoreboard';
import { forIn } from 'lodash'; 
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import { DifficultyLevel } from '@/models/game';

const props  = defineProps<{
    mode?:  Mode
}>()

const displayedScore = reactive(new Map<string, score.Score[]>([
    ["easy", []],
    ["medium", []],
    ["hard", []]
]))
const toast = useToast();


onMounted(() => {
    forIn(DifficultyLevel, (value: string, key: string) => {
        GetScore(value).then(scores => {
            if (!scores) {
                scores = [];
            }
            console.log(scores)
            displayedScore.set(value,scores);
            console.log(displayedScore)
        }).catch(err => {
            console.log(err)
            toast.add({severity: 'error', summary: 'failed to load scoreboard data'})
        })
    }) 

})


const router = useRouter();

</script>
<template>
    <TabView class="m-4">
        <TabPanel v-for="[key, value] in displayedScore" :key="key" :header="key">
            <div class="flex flex-column align-items-center">
                <DataTable :value="value" stripedRows tableStyle="min-width: 50rem" tableClass="w-7">
                    <Column field="username" header="Username"></Column>
                    <Column field="score" header="Score"></Column>
                    <Column field="created_at" header="Date">
                        <template #body="slotProps">
                            <div>
                                {{ format(slotProps.data.created_at, "dd/MM/yyyy") }}
                            </div>
                        </template>
                    </Column>
                    <template v-slot:empty>
                        <p class="text-center">No rank recorded at the moment. Start playing to register one !</p>
                    </template>
                </DataTable>
                <Button label="Back" class="m-4" @click="router.push('/')"></Button>
            </div>
        </TabPanel>
    </TabView>

    <Toast />
</template>
<style>
.p-tabview-nav {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}
</style>
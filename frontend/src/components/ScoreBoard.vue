<script setup lang="ts">
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { computed, onBeforeUpdate, onMounted , reactive, watch} from 'vue';
import {  useRouter } from 'vue-router';
import { GetScore } from '../../wailsjs/go/score/ScoreService';
import {score} from "../../wailsjs/go/models";
import { useToast } from 'primevue/usetoast';
import { format} from "date-fns";
import { Mode } from '@/models/scoreboard';
import { forIn } from 'lodash'; 
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import { DifficultyLevel } from '@/models/game';

const props  = withDefaults(defineProps<{
    mode:  Mode
    level: DifficultyLevel
    current?: score.Score
}>(), {mode: Mode.Static, level: DifficultyLevel.Medium})

const displayedScore = reactive(new Map<string, score.Score[]>([
    ["easy", []],
    ["medium", []],
    ["hard", []]
]))
const toast = useToast();
const stringLevel = computed(() => props.level.toString())


onMounted(() => {
    if (!props.level) {
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
    }else {
        GetScore(props.level).then((scores:score.Score[]) => {
            const sc = displayedScore.get(stringLevel.value) as score.Score[]
            if (!scores) {
                scores = [];
            }
            
            scores.forEach(score => {
                sc?.push(score)
            })
            
            if (props.current){
                sc?.push(props.current);
            }
            
            displayedScore.set(stringLevel.value,sc);
        }).catch(err => {
            console.log(err)
            toast.add({severity: 'error', summary: 'failed to load scoreboard data'})
        })
    }
})

const rowClass = (data) => { return  [{'bg-orange-500': data.id == -1}]}


const router = useRouter();

</script>
<template>
    <div>
        <TabView class="m-4" v-if="props.mode === Mode.Static">
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
        <DataTable id="livemode" :value="displayedScore.get(stringLevel)" :rowClass="rowClass" dataKey="id" sortField="score" :sortOrder="-1" stripedRows tableStyle="min-width: 50rem" tableClass="w-7" v-if="mode === Mode.Live">
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
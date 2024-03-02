<script setup lang="ts">
import {score} from "../../wailsjs/go/models";
import { format} from "date-fns";
import type { Mode } from '@/models/scoreboard';


const props = withDefaults(defineProps<{ scores: score.Score[], mode: Mode, size?:string}>(
), {size: "large"})
const rowClass = (data: any) => { return  [{'bg-orange-500': data.id == -1}]}


</script>
<template>
    <div>
        <DataTable  :size="props.size" :id="props.mode" :value="props.scores" :rowClass="rowClass" dataKey="id" sortField="score" :sortOrder="-1" stripedRows tableStyle="min-width: 50rem" tableClass="w-7">
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
</template>
<style>
@media only screen and (max-width: 1600px) {
    #livemode {
        display: none;
    }
}
</style>
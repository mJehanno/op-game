<script setup lang="ts">
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { onMounted , ref} from 'vue';
import {  useRouter } from 'vue-router';
import { GetScore } from '../../wailsjs/go/score/ScoreService';
import {score} from "../../wailsjs/go/models";
import { useToast } from 'primevue/usetoast';

const displayedScore = ref<Array<score.Score> | null>(null)
const toast = useToast();

onMounted(() => {
    GetScore().then(scores => {
        displayedScore.value=scores
    }).catch(err => {
        console.log(err)
        toast.add({severity: 'error', summary: 'failed to load scoreboard data'})
    })
})


const router = useRouter();

</script>
<template>
    <div class="flex flex-column align-items-center">
        <DataTable :value="displayedScore" stripedRows tableStyle="min-width: 50rem" tableClass="w-7">
            <Column field="username" header="Username"></Column>
            <Column field="score" header="Score"></Column>
            <Column field="created_at" header="Date"></Column>
            <template v-slot:empty>
                <p class="text-center">No rank recorded at the moment. Start playing to register one !</p>
            </template>
        </DataTable>
        <Button label="Back" class="m-4" @click="router.push('/')"></Button>
    </div>
    <Toast />
</template>
<script setup lang="ts">
import Dialog from 'primevue/dialog';
import InlineMessage from 'primevue/inlinemessage';
import FloatLabel from 'primevue/floatlabel';
import { reactive } from 'vue';

const props = defineProps<{
    visible: boolean;
    prompt?: string;
    answer?: number;
}>()

const dialogState = reactive<{
    user?: string,
    err ?: string,
}>({})

const emit = defineEmits<{
    userCreated: [username: string]
}>()

function gameOver(){
    if (dialogState.user && dialogState.user?.length > 0 && dialogState.user?.length <=3) {
        emit("userCreated", dialogState.user)
    }else if (dialogState.user && dialogState.user?.length > 3){
        dialogState.err = "username should contains max 3 characters"
    }else {
        dialogState.err = "username is required"
    }
}
</script>
<template>
    <Dialog v-model:visible="props.visible"  :content-style="{padding: '0 1em', height: '100%'}"  modal header="Game Over" :closable="false" v-on:hide="gameOver">
        <div class="flex flex-column w-full">
            <span class="p-text-secondary block ">The answer to {{ props.prompt }} was : {{ props.answer }}</span>
            <div class="flex flex-column m-2 flex-wrap">
                <span class="m-2">Register your score !</span>
                <FloatLabel class="m-2">
                    <InputText id="username" v-model="dialogState.user" mask="aaa" placeholder="aaa" aria-describedby="username-help"/>
                    <label for="username">Username</label>
                </FloatLabel>
                <small id="username-help">Can't contain more than 3 characters.</small>
                <InlineMessage class="m-1" v-if="dialogState.err" severity="error">{{ dialogState.err }}</InlineMessage>
            </div>
        </div>
        <template #footer>
            <div class="flex flex-row justify-content-center w-full" >
                <Button label="Submit" class="m-2" severity="success" @click="gameOver"/>
            </div>
        </template>
    </Dialog>
</template>
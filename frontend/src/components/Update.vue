<script setup lang="ts">
import {CheckLatest, Update} from '../../wailsjs/go/selfupdater/Updater';
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime';
import {ref, onMounted} from 'vue';

import Button from 'primevue/button';
import Message from 'primevue/message';

function performUpdate() {
    Update().then( _ => {
        successfulUpdate.value= true
    }).catch(err => {
        successfulUpdate.value = false
    })
}

function openURL(e: Event) {
    e.preventDefault();
    BrowserOpenURL(e.target)
}


const displayUpdateInfo = ref(false)
const successfulUpdate = ref<boolean | null>(null)

onMounted(() => {
  CheckLatest().then(isLatest => {
    displayUpdateInfo.value = !isLatest
  })
})

</script>
<template>
<div class="fixed bottom-0" v-if="displayUpdateInfo">
    <Message severity="warn">
        <div class=" p-1 flex  flex-column justify-content-start align-content-start align-items-start">
            <span>An update is available.</span>
            <span>You can either let this app download it (experimental) by clicking the 'selfupdate' button</span>
            <span>or you can download it yourself <a @click="openURL($event)" href="https://github.com/mJehanno/op-game/releases">here</a>.</span>
            <span>Quick reminder that on windows, your antivirus might block the download.</span>
        </div>
        <div class="m-1">
            <Button label="Self-Update"  @click="performUpdate"/>
        </div>
    </Message>
</div>
</template>
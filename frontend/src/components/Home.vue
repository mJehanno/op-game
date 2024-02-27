<script lang="ts" setup>
import { onMounted, ref } from 'vue';

import Message from 'primevue/message';
import Button from 'primevue/button';

import {GetCurrentVersion} from '../../wailsjs/go/version/VersionManager'
import {CheckLatest, Update} from '../../wailsjs/go/selfupdater/Updater';
import { BrowserOpenURL} from '../../wailsjs/runtime/runtime';

const version = ref<string|null>(null)
const displayUpdateInfo = ref(false)
const successfulUpdate = ref<boolean | null>(null)

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

onMounted(() => {
  CheckLatest().then(isLatest => {
    displayUpdateInfo.value = !isLatest
  })
    GetCurrentVersion().then(v => {
        version.value=v
    })
})

</script>
<template>
    <div class="flex flex-column align-items-center">
        <div class="flex flex-column m-8">
            <router-link to="/game" class="m-2"><Button label="Start Game" rounded severity="success" @click="$emit('startGame')" /></router-link>
            <router-link to="/scores" class="m-2"><Button label="See Scores" rounded severity="info" @click="$emit('startGame')" /></router-link>
        </div>

        <div>
            v{{ version }}
        </div>
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
    </div>
</template>
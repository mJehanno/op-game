<script lang="ts" setup>
import { onMounted, ref } from 'vue';

import Message from 'primevue/message';
import Button from 'primevue/button';
import Inplace from 'primevue/inplace';

import {GetCurrentVersion} from '../../wailsjs/go/version/VersionManager'
import {CheckLatest, Update} from '../../wailsjs/go/selfupdater/Updater';
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime';


import { DifficultyLevel } from '@/models/game';

const version = ref<string|null>(null)
const displayUpdateInfo = ref(false)
const successfulUpdate = ref<boolean | null>(null)
const displayDifficultyLevel = ref(false)

function performUpdate() {
    Update().then( _ => {
        successfulUpdate.value= true
    }).catch(err => {
        successfulUpdate.value = false
    })
}

const displayDifficulty = () => {
    displayDifficultyLevel.value = !displayDifficultyLevel.value
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
        <div class="flex flex-column m-8 w-full align-items-center">
            <Button class="m-3 w-3" label="Start Game" rounded severity="success" @click="displayDifficulty" />
                <div class="flex flex-column w-full" v-if="displayDifficultyLevel">
                    <div class="m-1 w-full">
                        <router-link :to="'/game/' + DifficultyLevel.Easy" class="m-1 w-full"><Button class="w-2" rounded severity="info" :label="DifficultyLevel.Easy" /></router-link>
                    </div>
                    <div class="m-1 w-full">
                        <router-link :to="'/game/' + DifficultyLevel.Medium" class="m-1 w-full"><Button class="w-2" rounded severity="warning" :label="DifficultyLevel.Medium" /></router-link>
                    </div>
                    <div class="m-1 w-full">
                        <router-link :to="'/game/' + DifficultyLevel.Hard" class="m-1 w-full"><Button class="w-2" rounded severity="danger" :label="DifficultyLevel.Hard"/></router-link>
                    </div>
                </div>
            <router-link to="/scores" class="m-3 w-full"><Button  class="w-3" label="See Scores" rounded severity="help" /></router-link>
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
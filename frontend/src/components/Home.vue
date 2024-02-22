<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import {GetCurrentVersion} from '../../wailsjs/go/version/VersionManager'

const version = ref<string|null>(null)

onMounted(() => {
    GetCurrentVersion().then(v => {
        version.value=v
    })
})

</script>
<template>
    <div class="flex flex-column m-8">
        <router-link to="/game" class="m-2"><Button label="Start Game" rounded severity="success" @click="$emit('startGame')" /></router-link>
        <router-link to="/scores" class="m-2"><Button label="See Scores" rounded severity="info" @click="$emit('startGame')" /></router-link>
    </div>

    <div>
        v{{ version }}
    </div>
</template>
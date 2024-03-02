<script setup lang="ts">
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
    <div>
        v{{ version }}
    </div>
</template>
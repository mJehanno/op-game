<script setup lang="ts">
import { useScoresStore } from '@/stores/scores';

    const scoreStore = useScoresStore();
</script>
<template>
    <span id="streak">Current streak : {{ scoreStore.current.score }}</span>
</template>
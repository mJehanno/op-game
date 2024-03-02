<script setup lang="ts">
const props = defineProps<{
    streak: number
}>()
</script>
<template>
    <span id="streak">Current streak : {{ props.streak }}</span>
</template>
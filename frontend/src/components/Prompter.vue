<script setup lang="ts">
const props = defineProps({
    prompt: {type: String, required: true}
})
</script>
<template>
    <div>
        <span>{{ prompt }}</span>
    </div>
</template>
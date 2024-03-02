<script setup lang="ts">
const props = defineProps<{
    maxHealth: number;
    current: number;
}>()
</script>
<template>
    <div>
        <i class="pi pi-heart-fill text-red-700 text-2xl" :key="n" v-for="n in props.current"></i>
        <i class="pi pi-heart text-red-700 text-2xl" :key="n" v-for="n in (props.maxHealth -props.current)"></i>
    </div>
</template>
<script setup lang="ts">
import {ref} from 'vue';
import Dock from 'primevue/dock';
import { useGameInfosStore } from '@/stores/game-infos';
import type { GameMode } from '@/models/game';

const store = useGameInfosStore();

const active = ref(store.selectedGame)

const items = ref<{
    label: string;
    icon: string;
}[]>([
    {label: "mult", icon: "pi-times"},
    {label: "min", icon: "pi-minus"},
    {label: "add", icon: "pi-plus"},
    {label: "divid", icon: "pi-percentage"},
])

function changeSelected(newSelected: GameMode) {
    active.value = newSelected;
    store.changeSelectedGame(newSelected);
}
</script>
<template>
    <Dock :model="items" position="left">
        <template #item="{ item }">
            <i v-tooltip.top="item.label" class="pi" :class="item.icon,{'text-orange-500': active == item.label} " style="width: 100%" @click="changeSelected(item.label as GameMode)" />
        </template>
    </Dock>
</template>
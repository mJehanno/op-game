<script setup lang="ts">
import Game from '@/components/Game.vue';
import { GameMode } from '@/models/game';
import { useGameInfosStore } from '@/stores/game-infos';
import Message from 'primevue/message';
import Button from 'primevue/button';
import { useRouter } from 'vue-router';

const store = useGameInfosStore();
const router = useRouter();

</script>
<template>
    <Game />
</template>
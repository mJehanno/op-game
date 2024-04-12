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
    <Game v-if="store.selectedGame == GameMode.Mult || store.selectedGame == GameMode.Divid" />
    <div v-else>
        <Message severity="error" >This game is not implemented yet.</Message> 
        <Button label="back" @click="router.push('/')"/>
    </div>
</template>
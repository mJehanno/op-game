<script setup lang="ts">
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { onBeforeMount } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import { Mode } from '@/models/scoreboard';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import ScoreBoard from '@/components/ScoreBoard.vue';
import  Title from '@/components/Title.vue';
import { useGameInfosStore } from '@/stores/game-infos';
import { useScoresStore } from '@/stores/scores';

const toast = useToast();
const gameInfos = useGameInfosStore();
const scoreStore = useScoresStore();
</script>
<template>
    <h1>meh</h1>
</template>
<script lang="ts" setup>
import Home from "@/components/Home.vue";
import Game from '@/components/Game.vue';
import ScoreBoard from '@/components/ScoreBoard.vue';

import { State} from '@/models/game';
import {reactive, onMounted, ref} from 'vue';
import {GetScore} from '../wailsjs/go/score/ScoreService';

const game = reactive({ state: State.Ready});
const seescore = ref(false);



onMounted(() => {
  GetScore().then(scores => {
    console.log(scores);
  }).catch(err => {
    console.error(err);
  })
});



function startGame() {
    game.state = State.Started
}

function seeScore() {
  seescore.value = true
}
</script>
<template>
  <h1>Mult-Game</h1>
  <router-view></router-view>
</template>
<style>
  h1{
    font-size: 8em;
    margin: 0;
  }
</style>
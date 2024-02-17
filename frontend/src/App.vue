<script lang="ts" setup>
  import Home from "@/components/Home.vue";

import Game from '@/components/Game.vue';
import { State} from '@/models/game';
import {reactive } from 'vue';

const game = reactive({ state: State.Ready});

function startGame() {
    game.state = State.Started
}
</script>
<template>
    <h1>Mult-Game</h1>
  <Home v-if="game.state != State.Started" @start-game="startGame" />
  <Game v-if="game.state == State.Started" @game-over="game.state = State.Over"/>
</template>
<style>
  h1{
    font-size: 10em;
  }
</style>
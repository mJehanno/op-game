<script lang="ts" setup>
import { onMounted} from 'vue';
import {GetScore} from '../wailsjs/go/score/ScoreService';



onMounted(() => {

  GetScore().then(scores => {
    console.log(scores);
  }).catch(err => {
    console.error(err);
  })
});


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
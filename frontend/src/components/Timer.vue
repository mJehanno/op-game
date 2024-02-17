<script setup lang="ts">
import {  watchEffect, onMounted } from "vue";
import { useTimer } from 'vue-timer-hook';

const props = defineProps<{
    sec: number;
}>()

const emits = defineEmits(["timeout"])

const time = new Date();
time.setSeconds(time.getSeconds()+ props.sec);

function reset(sec){
    const time = new Date();
    time.setSeconds(time.getSeconds() + sec);
    timer.restart(time)
}

const timer = useTimer(time);
onMounted(() => {
  watchEffect(async () => {
    if(timer.isExpired.value) {
        emits("timeout");
    }
  })
})

defineExpose({
    reset
})
</script>
<template>
    <div>
        <span>{{ timer.seconds }}</span>
    </div>
</template>
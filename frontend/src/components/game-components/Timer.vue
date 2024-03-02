<script setup lang="ts">
import {  watchEffect, onMounted, computed, ref } from "vue";
import { useTimer } from 'vue-timer-hook';

const COLOR_CODES = {
  info: {
    color: "green"
  },
  warning: {
    color: "orange",
    threshold: 10
  },
  alert: {
    color: "red",
    threshold: 5
  }
};

const props = defineProps<{
  seconds: number
}>()

let remainingPathColor = computed(() => {
  switch (true) {
    case timer.seconds.value > COLOR_CODES.warning.threshold:
      return COLOR_CODES.info.color
    case timer.seconds.value > COLOR_CODES.alert.threshold:
      return COLOR_CODES.warning.color
    default:
      return COLOR_CODES.alert.color
  }
})

const emits = defineEmits(["timeout"])
const time = new Date();
time.setSeconds(time.getSeconds()+ props.seconds);
const timer = useTimer(time, true);



function reset() {
  const time = new Date();
  time.setSeconds(time.getSeconds() + props.seconds);
  timer.restart(time);
}

onMounted(() => {
  watchEffect(async () => {
    if(timer.isExpired.value) {
        emits("timeout");
    }
  })
})

defineExpose({
    reset,
})
</script>
<template>
    <div class="base-timer">
    <svg class="base-timer__svg" viewBox="0 0 50 50" xmlns="http://www.w3.org/2000/svg">
        <g class="base-timer__circle">
        <circle class="base-timer__path-elapsed" cx="25" cy="25" r="20" />
        <path
          id="base-timer-path-remaining"
          stroke-dasharray="283"
          class="base-timer__path-remaining" :class="remainingPathColor"
          d="
            M 25, 25
            M -20, 0
            A 20,20 0 0,5 40,0
            A 20,20 0 0,5 -40,0
          "
        ></path>
        </g>
    </svg>
        <span id="base-timer-label" class="base-timer__label">{{ timer.seconds }}</span>
    </div>
</template>

<style>
/* Sets the containers height and width */
.base-timer {
  position: relative;
  height: 300px;
  width: 300px;
}

/* Removes SVG styling that would hide the time label */
.base-timer__circle {
  fill: none;
  stroke: none;
}

/* The SVG path that displays the timer's progress */
.base-timer__path-elapsed {
  stroke-width: 3px;
  stroke: green;
}


.base-timer__label {
  position: absolute;
  
  /* Size should match the parent container */
  width: 300px;
  height: 300px;
  
  /* Keep the label aligned to the top */
  top: 0;
  
  /* Create a flexible box that centers content vertically and horizontally */
  display: flex;
  align-items: center;
  justify-content: center;

  /* Sort of an arbitrary number; adjust to your liking */
  font-size: 48px;
}

.base-timer__path-remaining {
  /* Just as thick as the original ring */
  stroke-width: 3px;

  /* Rounds the line endings to create a seamless circle */
  stroke-linecap: round;

  /* Makes sure the animation starts at the top of the circle */
  transform: rotate(90deg);
  transform-origin: center;

  /* One second aligns with the speed of the countdown timer */
  transition: 1s linear all;

  /* Allows the ring to change color when the color value updates */
  stroke: currentColor;
}

.base-timer__svg {
  /* Flips the svg and makes the animation to move left-to-right */
  transform: scaleX(-1);
}

.base-timer__path-remaining.green {
  color: rgb(65, 184, 131);
}

.base-timer__path-remaining.orange {
  color: orange;
}

.base-timer__path-remaining.red {
  color: red;
}
</style>
<script setup lang="ts">

import Button from "@/components/ui/button/Button.vue";
import {ref} from "vue";

const audio = new Audio("https://www.zapsplat.com/wp-content/uploads/2015/sound-effects-46416/zapsplat_multimedia_button_click_005_53866.mp3");

const props = defineProps<{
  canUpgrade: boolean | undefined,
}>()

const disabled = ref(false)

function warnCantBuy() {
  disabled.value = true
  setTimeout(() => {
    disabled.value = false
  }, 1500)
}

function playAudio() {
  if (props.canUpgrade) {
    audio.play()
  }
}

function handleClick() {
  if (props.canUpgrade) {
    playAudio()
  } else {
    console.log('Cant upgrade')
    warnCantBuy()
  }
}
</script>

<template>
  <button :data-can-upgrade="props.canUpgrade" @click="handleClick()" :class="{ shake: disabled }">
    <slot/>
  </button>
</template>

<style scoped lang="scss">

/* Text */
span {
  background: linear-gradient(180deg, #F3F3F3 0%, #A5A5A5 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-fill-color: transparent;
  text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.16);
  letter-spacing: 0.5px;
  font-size: 18px;
}

/* Button */
button {
  width: 100%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  padding: 18px 22px;
  background: var(--gradient);
  border: 0;
  position: relative;
  z-index: 1;
  text-decoration: none;
  cursor: pointer;
  border-radius: 0 0 24px 24px;

  &[data-can-upgrade="false"] {
    color: #838383;
  }
}

/* Light notch */
button::after {
  content: '';
  display: block;
  width: 40px;
  height: 1px;
  position: absolute;

  top: 0px;
  box-shadow: 0px 1px 1px rgba(255, 255, 255, .1);
  transition: width 1s cubic-bezier(0.19, 1, 0.22, 1);
}

button {
  &[data-can-upgrade="true"]::after {
    background: #2A2C34;
    border: .5px solid #34FFAA;
    box-shadow: 0px -22px 26px rgba(0, 255, 163, 0.43), 0px 22px 26px rgba(75, 218, 166, 0.4), 0px -22px 20px rgba(89, 255, 195, 0.1), 0px 0px 16px #59FFC3;
  }

  &[data-can-upgrade="false"]::after {
    background: #FFEBEB;
    border: .5px solid #FF3434;
    box-shadow: 0px -22px 26px rgba(255, 0, 0, 0.43), 0px 22px 26px rgba(218, 75, 75, 0.4), 0px -22px 20px rgba(255, 89, 89, 0.1), 0px 0px 16px #FF5959;
  }
}

/* Light notch on button click */
button:active::after[data-can-upgrade="true"] {
  width: 140px; /* Adjust this value to your desired width */
}

/* Button on click */
button:active[data-can-upgrade="true"] {
  color: #EAFFFC;
  background: rgb(238, 174, 202);
  background: linear-gradient(190deg, rgba(238, 174, 202, 1) 0%, rgba(148, 187, 233, 0.6719999999999999) 100%);
}

.shake {
  animation: shake 0.5s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
  transform: translate3d(0, 0, 0);
}

@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }

  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }

  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }

  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}


button:active[data-can-upgrade="false"] {
}
</style>
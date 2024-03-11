<template>
  <div class="samba-time-game-dance-text">
    <div :key="text" class="samba-time-game-dance-text-inner">{{ text }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    text: {
      type: String,
      default: ''
    }
  }
})
</script>

<style scoped>
.samba-time-game-dance-text {
  position: fixed;
  width: 100vw;
  height: 100vh;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;

  display: flex;
  justify-content: center;
  align-items: center;
}

.samba-time-game-dance-text-inner {
  display: flex;
  justify-content: center;
  align-items: center;

  border-radius: 0;
  background-color: rgba(0, 0, 0, 0);
  color: #fff;

  font-size: 0.67rem;
  inset: 0;

  overflow: hidden;

  font-family: 'SuperMario', sans-serif !important;
  text-shadow: 22px 22px 0 #8d8d8d;

  opacity: 0;
  animation: increaseFontSizeHide 1s linear;
}
</style>

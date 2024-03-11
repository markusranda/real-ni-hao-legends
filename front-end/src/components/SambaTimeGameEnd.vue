<template>
  <div class="samba-time-game-end-title">
    <span>YOU DIED</span>
    <button @click="handleClick">retry</button>
    <audio id="game-over-player" preload="auto" src="du-har-tapt-echo.mp3"></audio>
  </div>
</template>

<script lang="ts">
export default {
  emits: ['retry'],
  mounted() {
    const audio = document.getElementById('game-over-player') as undefined | HTMLAudioElement

    if (audio) {
      audio.volume = 0.25
      audio.play()
    }
  },
  methods: {
    handleClick() {
      this.$emit('retry')
    }
  }
}
</script>

<style lang="scss">
.samba-time-game-end-title {
  font-family: 'DarkSouls', sans-serif;
  color: #872529;
  background: linear-gradient(
    to bottom,
    rgba(0, 0, 0, 0) 0%,
    rgba(0, 0, 0, 0.5) 10%,
    rgba(0, 0, 0, 0.75) 25%,
    rgba(0, 0, 0, 0.75) 75%,
    rgba(0, 0, 0, 0.5) 90%,
    rgba(0, 0, 0, 0) 100%
  );

  font-size: 5rem;
  font-weight: bold;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  width: 100%;
  height: 18rem;
}
</style>

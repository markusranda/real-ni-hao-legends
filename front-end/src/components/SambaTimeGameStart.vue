<template>
  <div class="samba-time-game-start">
    <span :key="countdownValue" style="animation: increaseFontSize 1s linear">
      {{ countdownValue }}
    </span>
  </div>

  <audio id="countdown-player" src="123.mp3"></audio>
</template>

<script lang="ts">
export default {
  data() {
    return {
      countdownValues: ['3', '2', '1', 'GO!'],
      countdownValue: undefined as undefined | string
    }
  },
  emits: ['done'],
  async mounted() {
    await this.startCountdownController()
    this.$emit('done')
  },
  methods: {
    async startCountdownController() {
      const audio = document.getElementById('countdown-player') as HTMLAudioElement
      if (!audio) throw Error("Can't start the samba without 3 2 1 song")
      audio.loop = false
      audio.volume = 0.15
      await audio.play()

      return new Promise<void>((resolve) => this.startCountdown(resolve, 0))
    },
    startCountdown(resolve: (value: void | PromiseLike<void>) => void, index: number) {
      const value = this.countdownValues[index]
      if (value) {
        this.countdownValue = value
        setTimeout(() => this.startCountdown(resolve, index + 1), 1000)
      } else {
        this.countdownValue = undefined
        resolve()
      }
    }
  }
}
</script>

<style>
.samba-time-game-start {
  display: flex;
  justify-content: center;
  align-items: center;

  border-radius: 0;
  background-color: rgba(0, 0, 0, 0);

  font-size: 1rem;
  inset: 0;

  overflow: hidden;

  span {
    font-family: 'SuperMario', sans-serif !important;
    color: #f1e14a;
    text-shadow: 22px 22px 0 #8d8d8d;
  }
}
</style>

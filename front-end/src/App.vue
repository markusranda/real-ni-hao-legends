<template>
  <TheLoading v-if="loading" />
  <TheIntroduction v-else-if="!initDone" @done="handleClickIntroductionComplete" />
  <TheGame v-if="showGame" />
</template>

<script lang="ts">
import TheGame from '@/components/TheGame.vue'
import TheIntroduction from '@/components/TheIntroduction.vue'
import { useGame } from './store/game'
import { useWebsocket } from './store/websocketStore'
import TheLoading from './components/TheLoading.vue'

export default {
  components: { TheLoading, TheIntroduction, TheGame },
  computed: {
    state() {
      return useGame().state
    },
    showGame() {
      const stateExists = this.state && Object.values(this.state).length > 0
      return this.initDone && stateExists
    },
    loading() {
      return useWebsocket().loading
    },
    initDone() {
      return useWebsocket().initDone
    }
  },
  mounted() {
    useWebsocket().init()
  },
  methods: {
    handleClickIntroductionComplete(name: string) {
      useWebsocket().initGameState(name)
      useWebsocket().initDone = true
    }
  }
}
</script>

<style lang="scss"></style>

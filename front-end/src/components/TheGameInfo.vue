<template>
  <div class="side the-game-info-right-wrapper nihao-box">
    <div class="the-game-table">
      <span class="d-flex">
        <input type="text" v-model="nukeTarget" placeholder="name.." class="col" />
        <Button @click="handleClickNuke">Nuke</Button>
      </span>

      <Separator />
      <div class="game-buttons">
        <Databingo />
        <SambaTime />
        <Quiz />
      </div>
      <Separator />

      <OnlinePlayers />

      <Separator />
    </div>
  </div>
</template>

<script lang="ts">
import { useWebsocket } from '@/store/websocketStore'
import { defineComponent } from 'vue'
import Databingo from './Databingo.vue'
import SambaTime from './SambaTime.vue'
import Quiz from './Quiz.vue'
import Separator from './Separator.vue'
import OnlinePlayers from './OnlinePlayers.vue'
import Button from './ui/button/Button.vue'

export default defineComponent({
  components: { Databingo, SambaTime, Quiz, Separator, OnlinePlayers, Button },
  data() {
    return {
      nukeTarget: ''
    }
  },
  methods: {
    handleClickNuke() {
      useWebsocket().send({
        type: 'effect.nuke',
        data: {
          target: this.nukeTarget
        }
      })
      this.nukeTarget = ''
    }
  }
})
</script>

<style scoped>
.game-buttons {
  display: flex;
  gap: 1rem;
  width: 200px;
}
</style>

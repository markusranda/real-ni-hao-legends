<template>
  <div>
    <button @click="handleClickOpen">Do the dance, take a chance</button>

    <Teleport to="body">
      <SambaTimeGame v-if="show" :handleClickClose="handleClickClose" />
    </Teleport>
  </div>
</template>

<script lang="ts">
import SambaTimeGame from './SambaTimeGame.vue'

export default {
  components: { SambaTimeGame },
  data() {
    return {
      show: false
    }
  },
  methods: {
    handleClickClose() {
      this.show = false
    },
    handleClickOpen() {
      this.show = true
    }
  }
}
</script>

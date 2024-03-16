<template>
  <div>
    <Button @click="handleClickOpen">Samba</Button>

    <Teleport to="body">
      <SambaTimeGame v-if="show" :handleClickClose="handleClickClose"/>
    </Teleport>
  </div>
</template>

<script lang="ts">
import SambaTimeGame from './SambaTimeGame.vue'
import Button from "@/components/ui/button/Button.vue";

export default {
  components: {Button, SambaTimeGame},
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

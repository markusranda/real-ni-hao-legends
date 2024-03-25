<template>
  <div>
    <Button @click="handleClickOpen">Quiz</Button>

    <Teleport to="body">
      <QuizGame v-if="show" :handleClickClose="handleClickClose" />
    </Teleport>
  </div>
</template>

<script lang="ts">
import QuizGame from './QuizGame.vue'
import Button from './ui/button/Button.vue'

export default {
  // eslint-disable-next-line vue/no-reserved-component-names
  components: { QuizGame, Button },
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
<style lang=""></style>

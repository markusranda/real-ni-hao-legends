<template>
  <div class="the-introduction">
    <h1>Who are you today?</h1>
    <div class="d-flex gap-1">
      <input v-model="townName" type="text" />
      <Button @click="handleClickSubmit">Submit</Button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      townName: ''
    }
  },
  emits: ['done'],
  methods: {
    handleClickSubmit() {
      const regex = /^[a-zA-Z0-9 ]{4,19}$/
      const valid = regex.test(this.townName)

      if (valid) {
        this.$emit('done', this.townName)
      }
    }
  }
})
</script>

<style scoped>
.the-introduction {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  height: 100vh;
  width: 100vw;

  padding: 4rem;
}
</style>

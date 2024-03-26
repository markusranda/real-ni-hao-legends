<template>
  <div class="the-introduction">
    <h1>Who are you today?</h1>
    <div class="the-introduction-grid">
      <div class="d-flex gap-1 w-100 justify-content-center">
        <input v-model="townName" name="name" type="text" />
        <Button @click="handleClickSubmit">Submit</Button>
      </div>

      <p class="error-text" v-if="error">
        Name must be between 2 to 25 characters long, <br />
        and contain no special characters
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      townName: '',
      error: false
    }
  },
  emits: ['done'],
  methods: {
    handleClickSubmit() {
      const regex = /^[a-zA-Z0-9 ]{2,25}$/
      const valid = regex.test(this.townName)

      if (valid) {
        this.error = false
        this.$emit('done', this.townName)
      } else {
        this.error = true
      }
    }
  }
})
</script>

<style scoped>
.error-text {
  text-align: center;
  color: rgb(233, 60, 60);
}

.the-introduction {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  height: 100vh;
  width: 100vw;

  padding: 4rem;
}

.the-introduction-grid {
  display: grid;
  justify-content: center;
  gap: 2rem;
  grid-template-rows: auto 4rem;

  width: 100%;
}
</style>

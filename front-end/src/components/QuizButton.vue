<template>
  <button class="quiz-button">
    <strong v-if="letter">{{ letter }}: </strong>
    <span>{{ text }}</span>
  </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    text: {
      type: String,
      required: true
    },
    index: {
      type: Number,
      default: undefined
    }
  },
  data() {
    return {
      alphabet: [
        'A',
        'B',
        'C',
        'D',
        'E',
        'F',
        'G',
        'H',
        'I',
        'J',
        'K',
        'L',
        'M',
        'N',
        'O',
        'P',
        'Q',
        'R',
        'S',
        'T',
        'U',
        'V',
        'W',
        'X',
        'Y',
        'Z'
      ]
    }
  },
  computed: {
    letter() {
      if (this.index === undefined) return ''
      return this.alphabet[this.index]
    }
  }
})
</script>

<style scoped lang="scss">
.quiz-button {
  display: inline-block;
  padding: 15px 30px;
  font-size: 20px;
  font-weight: bold;
  color: #fff;
  background-color: #0e0e0e; /* Dark blue */
  border: 2px solid #6097c4; /* Gold */
  border-radius: 2rem;
  cursor: pointer;
  text-decoration: none;

  &:hover {
    background-color: #5ea51f;
    color: #010e00;
  }
}
</style>

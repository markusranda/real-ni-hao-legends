<template>
  <div class="samba-time-scoreboard">
    <li v-for="(score, index) of scores" :key="`samba_time_item_${index}`">
      <span>{{ score.townName }}</span>
      <span>{{ score.score }}</span>
    </li>
  </div>
</template>

<script lang="ts">
import { useGame } from '@/store/game'
import { defineComponent } from 'vue'

export default defineComponent({
  computed: {
    scores() {
      const scores = [...useGame().samba.scores]
      scores.sort((a, b) => b.score - a.score)
      return scores
    }
  }
})
</script>

<style scoped>
.samba-time-scoreboard {
  filter: drop-shadow(10px 10px 15px #000);
  height: 100%;
  overflow-y: scroll;
}

li {
  list-style: none;
  display: grid;
  grid-template-columns: 20rem auto;
  gap: 1rem;
  font-weight: bold;
  font-size: 1.25rem;
  color: #fff;
}
</style>

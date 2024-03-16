<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue'
import {useGame} from '@/store/game'
import NetworkStatus from '@/components/NetworkStatus.vue'

const valutaType = ref('Gryn')
const state = computed(() => useGame().state)


onMounted(() => {
  const texts = ['Gryn', 'Gronk', 'Spenn', 'Peng', 'DallaBills', 'Penga']
  let index = 0

  setInterval(() => {
    valutaType.value = texts[index]
    index = (index + 1) % texts.length
  }, 10000)
})
</script>
<template>
  <div class="cool-colors">
    <div class="d-flex gap-2 top-info-wrapper">
      <div>
        <h1 class="town-name">{{ state.town.name }}</h1>
        <NetworkStatus />
      </div>
      <h1 :class="state.town.money >= 0 ? 'positive' : 'negative'" class="peng-tekst">
        {{ state.town.money }} {{ valutaType }}
      </h1>
    </div>

  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Pacifico&display=swap');

.top-info-wrapper {
  display: flex;
  align-items: center;
  height: 7em;
  flex-direction: row;
  padding: 42px;
  justify-content: space-between;
}

.cool-colors {
  filter: contrast(160%) brightness(100%);
  background: linear-gradient(116deg, rgba(181, 181, 232, 0.4), rgba(241, 173, 255, 0.4)),
  url(/noise.svg);
}

.peng-tekst {
  font-family: 'Courier New', Courier, monospace; /* Use a monospace font */
  color: rgba(76, 175, 80, 0.5); /* Use a semi-transparent green color */
  text-shadow: 1px 1px 2px black;
}
</style>

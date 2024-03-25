<template>
  <div class="town-grid">
    <TownBuilding v-for="building in buildings" :key="building.key" :buildingId="building.key" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import TownBuilding from './TownBuilding.vue'
import { NHBuilding } from '@/models/nh-building'
import { useGame } from '@/store/game'

export default defineComponent({
  components: { TownBuilding },
  computed: {
    state() {
      return useGame().state
    },
    buildings() {
      return Object.values<NHBuilding>(this.state.town.buildings)
    }
  }
})
</script>

<style scoped>
.town-grid {
  display: grid;
  gap: 16px;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  grid-template-rows: 20rem;
}
</style>

<script lang="ts" setup>
import {computed} from 'vue'
import {useGame} from '@/store/game'
import BuildingV2 from '@/components/BuildingV2.vue'
import {NHBuilding} from "@/models/nh-building";
import {useWebsocket} from "@/store/websocketStore";

const items = computed(() => {
  const buildings = useGame().state.inventory.buildings
  const grouped = buildings.reduce((acc: { [key: string]: any }, building) => {
    const key = building.key
    if (!acc[key]) {
      acc[key] = {...building, count: 0}
    }
    acc[key].count++
    return acc
  }, {})
  return Object.values(grouped)
})

function moveToTown(building: NHBuilding) {
  useWebsocket().send({
    type: 'building.move_to_town',
    data: {
      buildingId: building.key
    }
  })
}
</script>

<template>
  <div>
    <div v-for="(item, i) in items" :key="i">
      <BuildingV2 :building="item" :disabled="true" @click="moveToTown"/>
      <div>Count: {{ item.count }}</div>
    </div>
  </div>
</template>

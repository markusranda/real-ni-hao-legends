<script lang="ts" setup>
import {computed} from 'vue'
import {useGame} from '@/store/game'
import BuildingV2 from '@/components/building/BuildingV2.vue'
import {NHBuilding} from "@/models/nh-building";
import {useWebsocket} from "@/store/websocketStore";

const inventoryItems = computed(() => useGame().state.inventory.buildings)
const townItems = computed(() => Object.values(useGame().state.town.buildings))

function moveToTown(building: NHBuilding) {
  useWebsocket().send({
    type: 'building.move_to_town',
    data: {
      uniqueId: building.uniqueId,
      slot: 0 // Todo implement slot
    }
  })
}

function moveToInventory(building: NHBuilding) {
  useWebsocket().send({
    type: 'building.move_to_inventory',
    data: {
      uniqueId: building.uniqueId
    }
  })
}
</script>

<template>
    <div class="outer">
      <div>
        <h4>Town</h4>
        <div class="party-grid">
          <div v-for="(item, i) in townItems" :key="i">
            <BuildingV2 :building="item" :disabled="true" @click="() => moveToInventory(item)"/>
          </div>
        </div>
      </div>
      <div>
        <h4>Inventory</h4>
        <div class="inventory-grid">
          <div v-for="(item, i) in inventoryItems" :key="i">
            <BuildingV2 :building="item" :disabled="true" @click="() => moveToTown(item)"/>
          </div>
        </div>
      </div>
      <div>

      </div>
    </div>
</template>

<style scoped>

.outer {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 32px;
  height: 70vh;
}

.inventory-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  grid-template-rows: repeat(5, 1fr);
  gap: 10px;
}

.party-grid {
  display: grid;
  grid-template-columns: repeat(1, 1fr);
  grid-template-rows: repeat(6, 1fr);
  gap: 10px;
}
</style>

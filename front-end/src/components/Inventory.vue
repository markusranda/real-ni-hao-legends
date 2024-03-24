<script lang="ts" setup>
import {computed} from 'vue'
import {useGame} from '@/store/game'
import BuildingV2 from '@/components/building/BuildingV2.vue'
import {NHBuilding} from "@/models/nh-building";
import {useWebsocket} from "@/store/websocketStore";
import Separator from "@/components/Separator.vue";

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
    <div class="nihao-box">
      <h4>Town</h4>
      <div class="party-grid">
        <div v-for="(item, i) in townItems" :key="i">
          <div class="party-building" @click="() => moveToInventory(item)">
            <BuildingV2 :building="item" :disabled="true" class="building"/>
            <div class="party-building-info">
              <h6>{{ item.name }}</h6>
              <div>inc: {{ item.income }}</div>
              <div>cost: {{ item.upgradeCost }}</div>
            </div>
          </div>
          <Separator style="margin: 0"/>

        </div>
      </div>
    </div>
    <div class="nihao-box">
      <h4>Inventory</h4>
      <div class="inventory-grid">
        <div v-for="(item, i) in inventoryItems" :key="i">
          <BuildingV2 :building="item" :disabled="true" @click="() => moveToTown(item)" class="building"/>
        </div>
      </div>
    </div>
    <div class="nihao-box">
      Details about selected building
    </div>
  </div>
</template>

<style scoped lang="scss">

.building {
  width: 64px;
  height: 64px;
  border-radius: 50%;
}

.outer {
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 32px;
}

.inventory-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-template-rows: repeat(5, 1fr);
  gap: 10px;
}

.party-grid {
  display: grid;
  grid-template-columns: repeat(1, 1fr);
  grid-template-rows: repeat(6, 1fr);
  gap: 10px;
}

.party-building {
  display: flex;
  flex-direction: row;
  cursor: pointer;

  h6 {
    text-transform: uppercase;
    padding: 0;
    margin: 0;
  }

  &-info {
    display: flex;
    flex-direction: column;
    margin-left: 16px;
  }

  /* ... */

  transition: filter 0.3s ease-out;

  &:hover {
    filter: drop-shadow(0 0 20px rgba(0, 0, 0, 0.3));
  }
}
</style>

<script lang="ts" setup>
import { Ref, computed, ref } from 'vue'
import { useGame } from '@/store/game'
import BuildingV2 from '@/components/town/BuildingV2.vue'
import { NHBuilding } from '@/models/nh-building'
import { useWebsocket } from '@/store/websocketStore'
import Separator from '@/components/Separator.vue'
import Button from '@/components/ui/button/Button.vue'

const inventoryItems = computed(() => useGame().state.inventory.buildings)
const townItems = computed(() => Object.values(useGame().state.town.buildings))
const selectedBuilding: Ref<NHBuilding | undefined> = ref(undefined)
const addBtnDisabled = computed(() => {
  let found = false

  townItems.value.find((building) => {
    if (selectedBuilding.value?.name === building.name) found = true
  })

  return found
})

function handleClickInventoryBuilding(building: NHBuilding) {
  selectedBuilding.value = building
}

function handleClickDetailsAdd(building: NHBuilding) {
  if (addBtnDisabled.value) return

  moveToTown(building)
  selectedBuilding.value = undefined
}

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
    <div class="inventory-column nihao-box">
      <h4>Town</h4>
      <div class="party-grid">
        <div v-for="item in townItems" :key="`inventory_town_${item.uniqueId}`">
          <div class="party-building" @click="() => moveToInventory(item)" :title="item.imgUrl">
            <BuildingV2 :building="item" :disabled="true" class="building" />
            <div class="party-building-info">
              <h6>{{ item.name }}</h6>
              <div>inc: {{ item.income }}</div>
              <div>cost: {{ item.upgradeCost }}</div>
            </div>
          </div>
          <Separator style="margin: 0" />
        </div>
      </div>
    </div>

    <div class="inventory-column nihao-box">
      <h4>Inventory</h4>
      <div class="inventory-grid">
        <div v-for="item in inventoryItems" :key="`inventory_backpack_${item.uniqueId}`">
          <BuildingV2
            :building="item"
            :disabled="true"
            @click="() => handleClickInventoryBuilding(item)"
            @dblclick="() => moveToTown(item)"
            class="building"
            :title="item.name"
          />
        </div>
      </div>
    </div>

    <div class="inventory-column nihao-box">
      <template v-if="selectedBuilding">
        <h4>Details</h4>
        <div class="details-column-list">
          <div class="image-container">
            <img :src="selectedBuilding.imgUrl" alt="image of building please" />
          </div>

          <div class="details-column-list-table">
            <span>Name:</span>
            <span>{{ selectedBuilding.name ?? 'N/A' }}</span>

            <span>Level:</span>
            <span>{{ selectedBuilding.level ?? 'N/A' }}</span>

            <span>Rarity:</span>
            <span>{{ selectedBuilding.rarity ?? 'N/A' }}</span>

            <span>Income:</span>
            <span>{{ selectedBuilding.income ?? 'N/A' }}</span>

            <span>Upgrade cost:</span>
            <span>{{ selectedBuilding.upgradeCost ?? 'N/A' }}</span>

            <span>Base cost:</span>
            <span>{{ selectedBuilding.baseCost ?? 'N/A' }}</span>
          </div>

          <Button
            @click="() => (selectedBuilding ? handleClickDetailsAdd(selectedBuilding) : {})"
            :disabled="addBtnDisabled"
          >
            Add
          </Button>
        </div>
      </template>
    </div>
  </div>
</template>

<style scoped lang="scss">
.building {
  width: 64px;
  height: 64px;
  border-radius: 50%;
}

.inventory-column {
  display: grid;
  grid-template-rows: auto 1fr;
}

.outer {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  gap: 1rem;
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

.details-column-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;

  list-style: none;

  .image-container {
    width: 100%;
    height: 18rem;
    overflow: hidden;
    padding: 0 1rem 0 1rem;
  }

  img {
    width: 100%;
    height: auto;
    object-fit: cover;
  }

  .details-column-list-table {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }
}
</style>

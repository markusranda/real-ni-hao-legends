<template>
  <div class="house">
    <h5>{{ building.name }}</h5>
    <img class="building-image" :src="building.imgUrl" alt="" srcset="" draggable="false" />
    <div class="d-flex flex-column">
      <span>{{ `income: ${building.income}` }}</span>
      <span>{{ `level: ${building.level}` }}</span>
      <div class="house-buttons">
        <button :class="{ unavailable: !canUpgrade }" @click="handleClickUpgradeBuilding">
          {{ building.upgradeCost }} $
        </button>
        <button @click="handleClickSellBuilding">sell</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { NHBuilding } from '@/models/nh-building'
import { buildingUpgrade } from '@/commands/buildings'
import { computed } from 'vue'
import { useGame } from '@/store/game'

const props = defineProps<{
  buildingId: string
}>()

const building = computed(() => useGame().state.town.buildings[props.buildingId])
const money = computed(() => useGame().state.town.money)
function handleClickUpgradeBuilding() {
  buildingUpgrade(building.value.key)
}

function handleClickSellBuilding() {
  // todo
}

function getBuildingUpgradeCost() {
  return building.value.baseCost * building.value.level
}

function canUpgrade() {
  return getBuildingUpgradeCost() <= money.value
}
</script>

<style scoped>
.house {
  display: flex;
  flex-direction: column;
  padding: 32px;
  border-radius: 24px;
  width: 300px;

  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  background: #f5f5f5;
  transition: all 0.3s ease;

  img {
    transition: all 0.4s ease;
  }
}

.house:active {
  transform: scale(0.95);
}

.house:hover {
  filter: contrast(95%);

  img {
    filter: contrast(120%);
  }
}

.building-image {
  width: 250px;
  margin-bottom: 16px;
  margin-top: 16px;
}
</style>

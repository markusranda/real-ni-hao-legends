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
import {NHBuilding} from '@/models/nh-building'
import {buildingUpgrade} from '@/commands/buildings'
import {computed} from 'vue'
import {useGame} from '@/store/game'

const props = defineProps<{
  building: NHBuilding
  disabled: boolean
}>()

const building = props.building
const money = computed(() => useGame().state.town.money)
function handleClickUpgradeBuilding() {
  buildingUpgrade(building.key)
}

function handleClickSellBuilding() {
  // todo
}

function getBuildingUpgradeCost() {
  return building.baseCost * building.level
}

function canUpgrade() {
  return getBuildingUpgradeCost() <= money.value
}
</script>

<style scoped>
.house {
  display: flex;
  flex-direction: column;
  padding: 24px;
  border-radius: 24px;
  width: 200px;

  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  background: #f5f5f5;
  transition: filter 0.4s ease;
}

.house:hover {
  filter: contrast(95%);
}

.building-image {
  width: 167px;
  margin-bottom: 12px;
  margin-top: 12px;
}
</style>

<template>
  <div class="house" :style="backgroundImageStyle" @mouseover="showStats = true" @mouseleave="showStats = false">

  </div>
</template>

<script lang="ts" setup>
import {NHBuilding} from '@/models/nh-building'
import {buildingUpgrade} from '@/commands/buildings'
import {computed} from 'vue'
import {useGame} from '@/store/game'

let showStats = false;

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

const backgroundImageStyle = computed(() => ({
  backgroundImage: `url(${building.imgUrl})`,
  backgroundSize: 'cover',
  backgroundPosition: 'center',
}));
</script>

<style scoped>

.stats {
  opacity: 0;
  transition: opacity 0.2s ease-in-out;
}

.house:hover .stats {
  opacity: 1;
}


.house {
  display: flex;
  flex-direction: column;
  border-radius: 24px;

  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  background: #f5f5f5;
  transition: filter 0.4s ease;

  margin: 8px;

  text-shadow: -1px -1px 0 #fff, 1px -1px 0 #fff, -1px 1px 0 #fff, 1px 1px 0 #fff;

}

.house:hover {
  filter: contrast(95%);
}
</style>

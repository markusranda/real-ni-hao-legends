<template>
  <div class="house">
    <img :src="building.imgUrl" alt="" srcset="" draggable="false"/>
    <h5 class="title">{{ building.name }}</h5>

    <div class="stats">
      <span>{{ `inc: ${building.income}` }}</span>
      <span>{{ `lvl: ${building.level}` }}</span>
    </div>
    <div class="house-buttons">
      <BuildingUpgradeButton :can-upgrade="canUpgrade" @click="handleClickUpgradeBuilding">
        {{ building.upgradeCost }} $
      </BuildingUpgradeButton>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {buildingUpgrade} from '@/commands/buildings'
import {computed} from 'vue'
import {useGame} from '@/store/game'
import BuildingUpgradeButton from "@/components/building/BuildingUpgradeButton.vue";

const props = defineProps<{
  buildingId: string
}>()

const building = computed(() => useGame().state.town.buildings[props.buildingId])
const money = computed(() => useGame().state.town.money)

function handleClickUpgradeBuilding() {
  buildingUpgrade(building.value.key)
}

const canUpgrade = computed(() => building.value.upgradeCost <= money.value)
</script>

<style scoped>
.stats {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  margin: 16px;
}

.house {
  display: grid;
  grid-template-rows: 150px 40px auto;
  border-radius: 24px;

  border: 2px solid white;

  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  background: rgb(229,255,255);
  background: linear-gradient(0deg, rgb(247, 255, 252) 0%, rgba(255,248,233,1) 100%);
  transition: all 0.3s ease;

  img {
    border-radius: 24px 24px 0 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    object-fit: cover;
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

</style>

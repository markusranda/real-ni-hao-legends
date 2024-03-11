
<script lang="ts" setup>
import {computed} from "vue";
import {useGame} from "@/store/game";
import BuildingV2 from "@/components/BuildingV2.vue";

const userId = localStorage.getItem('UserId')!
const items = computed(() => useGame().players[userId].inventory.buildings)
</script>

<template>
  <div>
    <BuildingV2 v-for="(item, i) in items" :key="i" building="item" disabled="true" />
  </div>
</template>
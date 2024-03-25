<template>
  <Nuke />
  <InventoryButton v-model:isInventory="isInventory" />

  <div class="the-game">
    <div class="nihao-box the-game-grid">
      <TopInfo class="top" />
      <div class="main">
        <Inventory v-if="isInventory" class="main" />

        <div class="buildings-grid" v-if="!isInventory">
          <Building v-for="building in buildings" :key="building.key" :buildingId="building.key" />
        </div>
      </div>
      <Chat class="chat" />

      <TheGameInfo />
    </div>
  </div>
</template>

<script lang="ts">
import { NHBuilding } from '@/models/nh-building'
import { useGame } from '@/store/game'
import { defineComponent } from 'vue'
import Building from './building/Building.vue'
import TopInfo from './top/TopInfo.vue'
import Chat from './Chat.vue'
import Separator from './Separator.vue'
import Inventory from '@/components/Inventory.vue'
import Button from '@/components/ui/button/Button.vue'
import Nuke from '@/components/ui/Nuke.vue'
import InventoryButton from './InventoryButton.vue'
import TheGameInfo from './TheGameInfo.vue'

export default defineComponent({
  components: {
    Button,
    Nuke,
    Inventory,
    Building,
    TopInfo,
    Chat,
    Separator,
    InventoryButton,
    TheGameInfo
  },
  data() {
    return {
      // inventory
      isInventory: false
    }
  },
  computed: {
    state() {
      return useGame().state
    },
    chat() {
      return useGame().chat
    },
    buildings() {
      return Object.values<NHBuilding>(this.state.town.buildings)
    }
  }
})
</script>

<style scoped>
.the-game-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 400px;
  grid-template-rows: 100px 1fr 220px;
  gap: 8px 32px;
  padding: 0;
  grid-template-areas:
    'top top top'
    'main main side'
    'chat chat side';

  height: 95vh;
  overflow: hidden;
}

.top {
  grid-area: top;
}

.main {
  grid-area: main;
}

.chat {
  grid-area: chat;
}

.side {
  grid-area: side;
}

.main,
.side,
.chat {
  margin: 16px;
}

.buildings-grid {
  display: grid;
  gap: 16px;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
}
</style>

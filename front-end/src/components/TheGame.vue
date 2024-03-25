<template>
  <Nuke />
  <InventoryButton v-model:isInventory="isInventory" />

  <div class="the-game-grid">
    <TopInfo class="top" />
    <Inventory v-if="isInventory" class="main" />
    <Town v-else class="main" />
    <Chat class="chat" />

    <TheGameInfo />
  </div>
</template>

<script lang="ts">
import { useGame } from '@/store/game'
import { defineComponent } from 'vue'
import TopInfo from './top/TopInfo.vue'
import Chat from './Chat.vue'
import Inventory from '@/components/Inventory.vue'
import Nuke from '@/components/ui/Nuke.vue'
import InventoryButton from './InventoryButton.vue'
import TheGameInfo from './TheGameInfo.vue'
import Town from './town/Town.vue'

export default defineComponent({
  components: {
    Nuke,
    TopInfo,
    Chat,
    InventoryButton,
    TheGameInfo,
    Town,
    Inventory
  },
  data() {
    return {
      isInventory: false
    }
  },
  computed: {
    state() {
      return useGame().state
    },
    chat() {
      return useGame().chat
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

  height: 100vh;
  overflow: hidden;
}

.top {
  grid-area: top;
}

.main {
  grid-area: main;
  overflow-y: scroll;
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
  margin: 1rem;
}
</style>

<template>
  <Nuke/>
  <button class="inventory-button" @click="isInventory = !isInventory">
    Inventory
  </button>
  <div class="the-game">
    <div class="nihao-box the-game-grid">
      <TopInfo class="top"/>
      <div class="main">

        <Inventory v-if="isInventory" class="main"/>

        <div class="buildings-grid" v-if="!isInventory">
          <Building
              v-for="building in buildings"
              :key="building.key"
              :buildingId="building.key"
          />
        </div>
      </div>
      <Chat class="chat"/>

      <div class="side the-game-info-right-wrapper nihao-box ">
        <div class="the-game-table">
            <span class="d-flex">
              <input type="text" v-model="nukeTarget" placeholder="name.." class="col"/>
              <Button class="col-2" @click="handleClickNuke">Nuke</Button>
            </span>

          <span class="d-flex gap-2">
              <input
                  :value="donateRetirementFund"
                  class="col"
                  type="number"
                  name="retirementFund"
                  min="0"
                  :max="state.town.money"
                  @input="handleInputRetirementFund"
              />
              <Button class="col-2" @click="handleClickDonate">Donate</Button>
            </span>

          <Separator/>
          <div class="game-buttons">
            <Databingo/>
            <SambaTime/>
          </div>
          <Separator/>

          <OnlinePlayers/>

          <Separator/>
        </div>
      </div>

      <div v-if="activeEvent" class="happenings-container">
        <div class="position-relative w-100 h-100 d-flex justify-content-center align-items-center">
          <div class="happenings-container-inner">
            <nav>
              <Button @click="handleClickEventExit">x</Button>
            </nav>
            <section>
              <h1>{{ activeEvent.name }}</h1>
              <p>{{ activeEvent.description }}</p>
            </section>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {NHBuilding} from '@/models/nh-building'
import {NHEvent} from '@/models/nh-event'
import {useGame} from '@/store/game'
import {defineComponent} from 'vue'
import Building from './building/Building.vue'
import TopInfo from './top/TopInfo.vue'
import Chat from './Chat.vue'
import Separator from './Separator.vue'
import Databingo from './Databingo.vue'
import SambaTime from './SambaTime.vue'
import OnlinePlayers from './OnlinePlayers.vue'
import Inventory from "@/components/Inventory.vue";
import Button from "@/components/ui/button/Button.vue";
import {useWebsocket} from "@/store/websocketStore";
import Nuke from "@/components/ui/Nuke.vue";

export default defineComponent({
  components: {
    Button,
    Nuke,
    Inventory,
    Building,
    TopInfo,
    Chat,
    Separator,
    Databingo,
    SambaTime,
    OnlinePlayers
  },
  data() {
    return {
      activeEvent: undefined as NHEvent | undefined,
      nukeTarget: '',
      retirementFund: 0.0,
      donateRetirementFund: 0.0,

      // databingo
      isModalOpen: false,
      isAgeConfirmed: false,

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
  },
  methods: {
    handleInputRetirementFund(e: Event) {
      const inputStr = (e?.target as HTMLInputElement)?.value ?? 0
      const inputValue = parseInt(inputStr)
      if (inputValue > this.state.town.money) this.donateRetirementFund = this.state.town.money
      else if (inputValue < 0) this.donateRetirementFund = 0
      else this.donateRetirementFund = inputValue
    },
    handleClickNuke() {
      useWebsocket().send({
        type: 'effect.nuke',
        data: {
          target: this.nukeTarget
        }
      })
      this.nukeTarget = ''
    },
    handleClickEventExit() {
      this.activeEvent = undefined
    },
    handleClickDonate() {
      if (this.state.town.money >= this.donateRetirementFund) {
        this.retirementFund += this.donateRetirementFund
        this.state.town.money -= this.donateRetirementFund
        this.donateRetirementFund = 0
      }
    }
  }
})
</script>

<style scoped>

.inventory-button {
  position: absolute;
  bottom: 0;
  right: 0;
  z-index: 1;
  margin: 1rem;
  padding: 0.5rem 1rem;
  box-shadow: none;
}

.game-buttons {
  display: flex;
  gap: 1rem;
  width: 200px;
}

.the-game-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 400px;
  grid-template-rows: 100px 1fr 220px;
  gap: 8px 32px;
  padding: 0;
  grid-template-areas:
    "top top top"
    "main main side"
    "chat chat side";

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

.main, .side, .chat {
  margin: 16px
}

.buildings-grid {
  display: grid;
  gap: 16px;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
}

</style>

<template>
  <div class="the-game">
    <div class="nihao-box the-game-container">
      <TopInfo />

      <div class="the-game-content">
        <div class="the-game-content-inner">
          <div class="buildings-grid">
            <Building
              v-for="building in buildings"
              :key="building.key"
              :buildingId="building.key"
            />
          </div>
          <Chat />
        </div>
        <div class="the-game-info-right-wrapper nihao-box">
          <div class="the-game-table">
            <span class="d-flex">
              <input type="text" v-model="nukeTarget" placeholder="name.." class="col" />
              <button class="col-2" @click="handleClickNuke">Nuke</button>
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
              <button class="col-2" @click="handleClickDonate">Donate</button>
            </span>

            <div>
              {{ `Income: ${moneyDelta}` }}
            </div>

            <Separator />
            <div class="game-buttons">
              <Databingo />
              <SambaTime />
              <LarsSideProject />
            </div>
            <Separator />

            <OnlinePlayers />

            <Separator />
          </div>
        </div>
      </div>

      <div v-if="activeEvent" class="happenings-container">
        <div class="position-relative w-100 h-100 d-flex justify-content-center align-items-center">
          <div class="happenings-container-inner">
            <nav>
              <button @click="handleClickEventExit">x</button>
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
import { NHBuilding } from '@/models/nh-building'
import { NHEvent } from '@/models/nh-event'
import { useGame } from '@/store/game'
import { defineComponent } from 'vue'
import Building from './Building.vue'
import TopInfo from './top/TopInfo.vue'
import Chat from './Chat.vue'
import Separator from './Separator.vue'
import Databingo from './Databingo.vue'
import SambaTime from './SambaTime.vue'
import LarsSideProject from './LarsSideProject.vue'
import OnlinePlayers from './OnlinePlayers.vue'

export default defineComponent({
  components: {
    Building,
    TopInfo,
    Chat,
    Separator,
    Databingo,
    SambaTime,
    LarsSideProject,
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

      moneyDelta: 0.0,
      name: 'Slim Yung Un',
      currentMessage: ''
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
      alert(`${this.nukeTarget} is now nuked to death, what a shame...`)
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
.game-buttons {
  display: flex;
  gap: 1rem;
  width: 200px;
}
</style>

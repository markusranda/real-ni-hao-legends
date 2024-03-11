<template>
  <div
    class="the-game"
    v-if="state?.town?.buildings && Object.keys(state.town.buildings).length > 0"
  >
    <div class="nihao-box the-game-container">
      <top-info />

      <div class="the-game-content">
        <div class="the-game-content-inner">
          <div class="buildings-grid">
            <Building
              v-for="building in Object.values<NHBuilding>(state.town.buildings)"
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
import { NHEvent } from './models/nh-event'
import Databingo from '@/components/Databingo.vue'
import { ws } from '@/main'
import { useGame } from '@/store/game'
import TopInfo from '@/components/top/topInfo.vue'
import { GameState } from './models/nh-gamestate'
import OnlinePlayers from '@/components/OnlinePlayers.vue'
import { NHSamba } from './models/nh-samba'
import Building from '@/components/Building.vue'
import SambaTime from '@/components/SambaTime.vue'
import Chat from '@/components/Chat.vue'
import Separator from '@/components/Separator.vue'
import { NHCommand } from '@/models/nh-command'
import LarsSideProject from '@/components/LarsSideProject.vue'
import { NHBuilding } from './models/nh-building'

const HAPPENINGS = {
  zeus: {
    id: 'zeus',
    name: 'Zeus',
    description: 'Zeus smites you with his electric schlong! Haha, all your money is now gone!'
  },
  choir: {
    id: 'choir',
    name: 'The choir needs you!',
    description:
      'You got pressured into singing in the choir of the infamous "Pensjonistfondet", and you will now have to sing your heart out.'
  }
} as Record<string, NHEvent>

export default {
  components: {
    LarsSideProject,
    Separator,
    Chat,
    SambaTime,
    Building,
    OnlinePlayers,
    TopInfo,
    Databingo
  },
  computed: {
    state() {
      return useGame().state
    },
    chat() {
      return useGame().chat
    }
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
      currentMessage: '',
      ws: undefined as WebSocket | undefined
    }
  },
  mounted() {
    this.connectToWs()
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
    },
    connectToWs() {
      ws.onopen = () => {
        console.log('connected to ws')
      }

      ws.onmessage = async (event) => {
        console.log(JSON.parse(event.data))
        this.updateGameState(JSON.parse(event.data))
      }
    },
    updateGameState(wsMessage: {
      Players: Record<string, GameState>
      Samba: NHSamba
      Chat: { Messages: NhMessage[] }
      Commands: NHCommand[]
    }) {
      const userId = localStorage.getItem('UserId')
      if (!userId) throw Error("Can't play the game without a name FOOoOOOoOoooOOol")

      const game = useGame()
      const partial = wsMessage?.Players[userId]

      const state = { ...game.state, ...partial }
      const players = wsMessage.Players
      const samba = wsMessage.Samba
      delete players[userId]

      game.players = players
      game.chat = wsMessage.Chat.Messages
      game.state = state
      game.samba = samba
      game.commands = wsMessage.Commands
    }
  }
}
</script>

<style>
.game-buttons {
  display: flex;
  gap: 1rem;
  width: 300px;
}
</style>

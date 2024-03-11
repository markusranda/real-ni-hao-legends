<template>
  <Modal
    :show="true"
    background-color="background-color: rgba(0, 0, 0, 0.9)"
    modal-container-class="samba-container"
    modal-header-class="samba-header"
    modal-body-class="samba-body"
    @close="() => handleClickClose()"
  >
    <template #header>
      <div></div>
      <div class="samba-header-title">
        <h1>SAMBA TIME</h1>
        <h3 class="d-flex gap-3">
          <span>
            {{ `Life: ${life}` }}
          </span>

          <span>
            {{ `Score: ${score}` }}
          </span>
        </h3>
      </div>
      <div class="samba-exit-btn" @click="() => handleClickClose()">X</div>
    </template>
    <template #body>
      <div></div>

      <SambaTimeGameStart
        v-if="sambaGameState === SAMBA_GAME_STATE.start"
        @done="handleEventStartDone"
      />
      <!-- TODO Add scenery change when players play for a while unknow -->
      <!-- TODO Add song playlist that matches the scenery -->
      <!-- TODO Background gradient must match the speed of the simulation -->
      <SambaTimeGameDance
        v-else-if="sambaGameState === SAMBA_GAME_STATE.dance"
        v-model:life="life"
        v-model:score="score"
        @done="handleEventDanceDone"
      />
      <SambaTimeGameEnd v-else-if="sambaGameState === SAMBA_GAME_STATE.end" @retry="handleRetry" />
      <div v-else></div>

      <SambaTimeScoreboard />
    </template>

    <template #footer>
      <div></div>
    </template>
  </Modal>
</template>

<script lang="ts">
import Modal from './Modal.vue'
import SambaTimeGameStart from './SambaTimeGameStart.vue'
import SambaTimeGameDance from './SambaTimeGameDance.vue'
import SambaTimeGameEnd from './SambaTimeGameEnd.vue'
import SambaTimeScoreboard from './SambaTimeScoreboard.vue'
import { send } from '@/main'
import { NHScore } from '@/models/nh-samba'
import { useGame } from '@/store/game'
import { NHCommand } from '@/models/nh-command'

const SAMBA_GAME_STATE = {
  start: 'start',
  dance: 'dance',
  end: 'end'
}

export default {
  components: {
    Modal,
    SambaTimeGameStart,
    SambaTimeGameDance,
    SambaTimeGameEnd,
    SambaTimeScoreboard
  },
  props: {
    handleClickClose: {
      type: Function,
      required: true
    }
  },
  data() {
    return {
      SAMBA_GAME_STATE: SAMBA_GAME_STATE,
      show: false,
      life: 3,
      score: 0,
      sambaGameState: SAMBA_GAME_STATE.start
    }
  },
  methods: {
    handleEventStartDone() {
      this.sambaGameState = SAMBA_GAME_STATE.dance
    },
    handleEventDanceDone() {
      this.sambaGameState = SAMBA_GAME_STATE.end
      const userId = localStorage.getItem('UserId')
      const command = {
        userId: userId,
        type: 'samba.score',
        data: {
          userId: userId,
          score: this.score,
          townName: useGame().state.town.name
        } as NHScore
      } as NHCommand
      send(command)
    },
    handleRetry() {
      this.score = 0
      this.life = 3
      this.sambaGameState = SAMBA_GAME_STATE.start
    }
  }
}
</script>

<style lang="scss">
@keyframes moveGradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

:root {
  --tropical-yellow: #ffd700;
  --amazon-green: #00ff7f;
  --rio-blue: #09546d;
}

.samba-header {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  width: 100%;
}

.samba-container {
  display: grid;
  grid-template-rows: auto minmax(1rem, 1fr) auto;
  width: 100%;

  background: linear-gradient(
    270deg,
    var(--tropical-yellow),
    var(--rio-blue),
    var(--amazon-green),
    #fff
  );
  background-size: 400% 400%;
  animation: moveGradient 10s ease infinite;
  margin: 2rem;
  padding: 2rem;
}

.samba-body {
  display: grid;
  align-items: center;
  grid-template-columns: 1fr 2fr 1fr;
  gap: 2rem;

  width: 100%;
  height: 100%;
}

.samba-header-title {
  display: flex;
  flex-direction: column;
  align-items: center;
  filter: drop-shadow(10px 10px 15px #000);

  h1,
  h3 {
    font-weight: bold;
    color: #fff;
  }
}

.samba-exit-btn {
  font-size: 2rem;
  font-weight: bold;
  color: #fff;
  filter: drop-shadow(10px 10px 15px #000);

  display: flex;
  margin-left: auto;
  margin-right: 2rem;

  padding: 0.5rem;

  &:hover {
    cursor: pointer;
  }
}
</style>

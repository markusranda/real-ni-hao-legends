<template>
  <Modal :show="true" @close="() => handleClickClose()">
    <template #header>
      <h1>! ! ! ! IT'S BINGO TIME ! ! ! !</h1>
    </template>
    <template #body>
      <div class="bingo-board">
        <div v-for="(cols, i) of numberMatrix" :key="`bingo_tile__${i}`">
          <div
            v-for="(cell, j) of cols"
            :key="`col_${i}_row_${j}`"
            class="bingo-board-tile"
            :class="`${cell.cssClass} ${cell.checked ? 'bingo-board-tile-checked' : ''}`"
            @click="() => handleClickCell(i, j)"
          >
            {{ cell.number ?? 'FREE' }}
          </div>
        </div>
      </div>

      <div>
        <span v-if="isPlaying" class="audio-control-btn" @click="audioEl?.pause()"> ⏸️ </span>
        <span v-else class="audio-control-btn" @click="audioEl?.play()"> ▶️ </span>
      </div>
      <BingoBall v-if="ballRolling" />
      <audio id="audio-player"></audio>
    </template>
    <template #footer>
      <button class="modal-default-button" @click="() => rollBall()">ROLL BALL!!!</button>
      <button class="modal-default-button" @click="() => handleClickClose()">OK</button>
    </template>
  </Modal>
</template>

<script lang="ts">
import Modal from './Modal.vue'
import BingoBall from './BingoBall.vue'

interface BingoTile {
  number: number | undefined
  checked: boolean
  cssClass: string
}

export default {
  components: {
    Modal,
    BingoBall
  },
  props: {
    handleClickClose: {
      type: Function,
      required: true
    }
  },
  data() {
    return {
      bingoBoardClass: '',
      numberMatrix: [] as BingoTile[][],
      audioEl: undefined as undefined | HTMLAudioElement,
      isPlaying: false,
      ballRolling: false
    }
  },
  mounted() {
    this.numberMatrix = this.getNumberMatrix()
    this.spawnSong()

    setInterval(() => {
      for (const col of this.numberMatrix) {
        for (const cell of col) {
          cell.cssClass = this.getRndColorClass()
        }
      }
    }, 250)
  },
  methods: {
    getRndNum() {
      return Math.round(Math.random() * 100)
    },
    getNumberMatrix() {
      const numberMatrix = []
      for (let i = 0; i < 5; i++) {
        const column = []
        for (let j = 0; j < 5; j++) {
          let number = this.getRndNum() as number | undefined
          if (i === 2 && j === 2) {
            number = undefined
          }
          column.push({ number: number, cssClass: '', checked: false })
        }
        numberMatrix.push(column)
      }

      return numberMatrix
    },
    getRndColorClass() {
      const colorClasses = [
        'bingo-board-tile-orange',
        'bingo-board-tile-green',
        'bingo-board-tile-blue',
        'bingo-board-tile-red'
      ]
      const index = Math.round(Math.random() * 10) % colorClasses.length

      return colorClasses[index]
    },
    spawnSong() {
      this.audioEl = document.getElementById('audio-player') as HTMLAudioElement

      this.audioEl.src = 'ivars_theme_da apene_sang.mp3'
      this.audioEl.play()
      this.audioEl.loop = true
      this.audioEl.volume = 0.25

      this.audioEl.onplay = () => (this.isPlaying = true)
      this.audioEl.onpause = () => (this.isPlaying = false)
    },
    handleClickCell(colIndex: number, cellIndex: number) {
      this.numberMatrix[colIndex][cellIndex].checked = true
    },
    rollBall() {
      this.ballRolling = true
      setTimeout(() => {
        this.ballRolling = false
      }, 2000)
    }
  }
}
</script>

<style scoped lang="scss">
.bingo-board {
  display: flex;
  justify-content: center;
  align-content: center;

  gap: 0.2rem;
}

.bingo-board-tile {
  display: flex;
  justify-content: center;
  align-items: center;
  border: solid 2px rgb(204, 204, 204);
  border-radius: 0;
  font-size: 1rem;

  height: 4rem;
  width: 4rem;

  &:hover {
    cursor: pointer;
  }
}

.bingo-board-tile-checked {
  border: solid 2px black !important;
}

.audio-control-btn {
  cursor: pointer;
}

.bingo-board-tile-red {
  box-shadow: 0px 0px 25px 8px rgba(230, 125, 107, 0.75);
}

.bingo-board-tile-blue {
  box-shadow: 0px 0px 25px 8px rgba(111, 165, 237, 0.75);
}

.bingo-board-tile-orange {
  box-shadow: 0px 0px 25px 8px rgba(237, 191, 111, 0.75);
}

.bingo-board-tile-green {
  box-shadow: 0px 0px 25px 8px rgba(121, 237, 111, 0.75);
}
</style>

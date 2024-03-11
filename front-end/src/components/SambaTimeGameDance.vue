<template>
  <div class="samba-time-game-dance">
    <div class="dance-floor">
      <div v-for="(col, i) in grid" :key="`dance_tile_col${i}`">
        <div
          v-for="(danceTile, j) in col"
          :key="`dance_tile${i}${j}`"
          class="dance-floor-tile"
          :class="danceTile.marked ? 'dance-floor-tile-active' : ''"
        >
          <img :src="danceTile?.shoe" alt="" class="dance-floor-shoe" />
        </div>
      </div>

      <SambaTimeGameDanceText :text="danceMoves[danceMoveIndex].name" />

      <audio id="samba_time_player" preload="auto" src="samba_time.mp3"></audio>
      <audio id="good-sound-player" preload="auto" src="yep.mp3"></audio>
      <audio id="bad-sound-player" preload="auto" src="damage.mp3"></audio>
    </div>
  </div>
</template>

<script lang="ts">
import { DANCE_MOVES } from '@/plugins/samba'
import SambaTimeGameDanceText from './SambaTimeGameDanceText.vue'

interface NHDanceFloorTile {
  shoe: string | undefined
  marked: boolean
}

const MIN_DANCE_TIMEOUT = 1500
const MAX_DANCE_TIMEOUT = 5000

export default {
  components: { SambaTimeGameDanceText },
  props: {
    life: {
      type: Number,
      required: true
    },
    score: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      danceTimeout: MAX_DANCE_TIMEOUT,
      leftShoeCoords: {
        x: 2,
        y: 2
      },
      rightShoeCoords: {
        x: 3,
        y: 2
      },
      markedTiles: [] as string[],
      running: true,
      lastFrameTime: 0,
      deltaTime: 0,
      timer: 0,
      markedTileTimeout: 0,
      danceMoves: DANCE_MOVES,
      danceMoveArrIndex: 0,
      danceMoveIndex: 0
    }
  },
  emits: ['done', 'update:score', 'update:life'],
  computed: {
    grid() {
      const grid: NHDanceFloorTile[][] = []

      for (let i = 0; i < 6; i++) {
        const col = []
        for (let j = 0; j < 6; j++) {
          let item = { shoe: undefined, marked: false } as NHDanceFloorTile
          if (this.leftShoeCoords.x === i && this.leftShoeCoords.y === j)
            item.shoe = 'left_shoe.webp'
          else if (this.rightShoeCoords.x === i && this.rightShoeCoords.y === j)
            item.shoe = 'right_shoe.webp'

          for (const coordStr of this.markedTiles) {
            const coorStrArr = coordStr.split('')
            const x = parseInt(coorStrArr[0])
            const y = parseInt(coorStrArr[1])

            if (x === i && y === j) item.marked = true
          }

          col.push(item)
        }
        grid.push(col)
      }

      return grid
    }
  },
  watch: {
    danceTimeout: {
      handler() {
        const audio = document.getElementById('samba_time_player') as undefined | HTMLAudioElement
        if (audio) {
          const MIN_VALUE = 1.0
          const MAX_VALUE = 2.0
          const ratio =
            (this.danceTimeout - MIN_DANCE_TIMEOUT) / (MAX_DANCE_TIMEOUT - MIN_DANCE_TIMEOUT)
          const scaledValue = MAX_VALUE - ratio * (MAX_VALUE - MIN_VALUE)
          const rounded = parseFloat(scaledValue.toFixed(2))

          audio.playbackRate = rounded
        }
      }
    }
  },
  mounted() {
    this.init()
  },
  beforeUnmount() {
    this.destroyListeners()
  },
  methods: {
    async init() {
      this.addListeners()
      this.startSong()
      requestAnimationFrame(this.simulateSambaDance)
    },
    simulateSambaDance(timestamp: number) {
      this.deltaTime = timestamp - this.lastFrameTime
      this.markedTileTimeout -= this.deltaTime
      this.lastFrameTime = timestamp

      this.updateLife()

      // Check if game should continue
      if (!this.running) {
        this.$emit('done')
        return
      }

      this.updateBlinking()
      this.updateCollision()
      if (this.markedTileTimeout <= 0) this.updateMarkedTiles()

      requestAnimationFrame(this.simulateSambaDance)
    },
    updateBlinking() {
      const markedTiles: NodeListOf<HTMLElement> = document.querySelectorAll(
        '.dance-floor-tile-active'
      )
      const modifier =
        (this.danceTimeout - MIN_DANCE_TIMEOUT) / (MAX_DANCE_TIMEOUT - MIN_DANCE_TIMEOUT)
      const sec = modifier * 1.5 + 0.6
      const dur = sec.toFixed(1)

      for (const markedTile of Array.from(markedTiles)) {
        markedTile.style.animationDuration = `${dur}s`
      }
    },
    updateLife() {
      if (this.life <= 0) this.running = false
    },
    updateMarkedTiles() {
      const danceArr: string[] = []

      if (this.markedTiles.length < 1) {
        this.patternDefault(danceArr)
      } else {
        this.patternDance(danceArr)
      }

      this.markedTiles = danceArr
      this.markedTileTimeout = this.danceTimeout
    },
    updateMarkedTileTimeout() {
      if (this.danceTimeout >= MIN_DANCE_TIMEOUT) this.danceTimeout -= 200
    },
    patternDefault(arr: string[]) {
      const center = Math.floor(this.grid.length / 2)
      arr.push(`${center}${center}`)
      arr.push(`${center - 1}${center}`)
    },
    patternDance(arr: string[]) {
      const danceMoveArr = this.danceMoves[this.danceMoveIndex].moves
      const danceMove = danceMoveArr[this.danceMoveArrIndex]

      const coordsLeft = this.markedTiles[0]
      const [leftXStr, leftYStr] = coordsLeft.split('')
      const leftX = parseInt(leftXStr) + danceMove[0][0]
      const leftY = parseInt(leftYStr) + danceMove[0][1]
      arr.push(`${leftX}${leftY}`)

      let coordsRight = this.markedTiles[1]
      let [rightXStr, rightYStr] = coordsRight.split('')
      const rightX = parseInt(rightXStr) + danceMove[1][0]
      const rightY = parseInt(rightYStr) + danceMove[1][1]
      arr.push(`${rightX}${rightY}`)

      this.danceMoveArrIndex++
      if (this.danceMoveArrIndex >= danceMoveArr.length) {
        this.danceMoveIndex++
        this.danceMoveArrIndex = 0
        if (this.danceMoveIndex >= this.danceMoves.length) {
          this.danceMoveIndex = 0
          this.updateDanceMoves()
          this.updateMarkedTileTimeout()
        }
      }
    },
    updateDanceMoves() {
      for (let i = this.danceMoves.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1))
        ;[this.danceMoves[i], this.danceMoves[j]] = [this.danceMoves[j], this.danceMoves[i]]
      }
    },
    patternRandom(arr: string[]) {
      for (let i = 0; i < 2; i++) {
        const x = Math.floor(Math.random() * (5 - 0 + 1)) + 0
        let y = x
        while (y === x) y = Math.floor(Math.random() * (5 - 0 + 1)) + 0

        arr.push(`${x}${y}`)
      }
    },
    updateCollision() {
      const trutArr = []
      if (this.markedTiles.length > 0) {
        for (const coordStr of this.markedTiles) {
          const strArr = coordStr.split('')
          const x = parseInt(strArr[0])
          const y = parseInt(strArr[1])

          // Check if player is in correct tiles
          if (this.grid[x][y].shoe) trutArr.push(true)
          else trutArr.push(false)
        }
        const bothShoesOnTiles = trutArr.reduce((sum, val) => sum && val, true)

        if (bothShoesOnTiles) {
          this.markedTileTimeout = 0
          this.$emit('update:score', this.score + 1)
          this.playGoodSound()
        } else if (this.markedTileTimeout <= 0) {
          this.$emit('update:life', this.life - 1)
          this.playBadSound()
        }
      }
    },
    handleKeyDown(e: KeyboardEvent) {
      if (!this.running) return
      switch (e.key) {
        case 'w':
          if (this.leftShoeCoords.y - 1 >= 0) this.leftShoeCoords.y -= 1
          break
        case 'a':
          if (this.leftShoeCoords.x - 1 >= 0) this.leftShoeCoords.x -= 1
          break
        case 's':
          if (this.leftShoeCoords.y + 1 <= this.grid.length - 1) this.leftShoeCoords.y += 1
          break
        case 'd':
          if (this.leftShoeCoords.x + 1 <= this.grid.length - 1) this.leftShoeCoords.x += 1
          break
        case 'ArrowUp':
          if (this.rightShoeCoords.y - 1 >= 0) this.rightShoeCoords.y -= 1
          break
        case 'ArrowLeft':
          if (this.rightShoeCoords.x - 1 >= 0) this.rightShoeCoords.x -= 1
          break
        case 'ArrowDown':
          if (this.rightShoeCoords.y + 1 <= this.grid.length - 1) this.rightShoeCoords.y += 1
          break
        case 'ArrowRight':
          if (this.rightShoeCoords.x + 1 <= this.grid.length - 1) this.rightShoeCoords.x += 1
          break

        default:
          break
      }
    },
    startSong() {
      const audio = document.getElementById('samba_time_player') as undefined | HTMLAudioElement
      if (audio) {
        audio.volume = 0.05
        audio.play()
      }
    },
    playBadSound() {
      const audio = document.getElementById('bad-sound-player') as undefined | HTMLAudioElement
      if (audio) {
        audio.volume = 0.25
        audio.play()
      }
    },
    playGoodSound() {
      const audio = document.getElementById('good-sound-player') as undefined | HTMLAudioElement
      if (audio) {
        audio.volume = 1
        audio.play()
      }
    },
    addListeners() {
      addEventListener('keydown', this.handleKeyDown)
    },
    destroyListeners() {
      removeEventListener('keydown', this.handleKeyDown)
    }
  }
}
</script>

<style>
@keyframes blink-animation {
  0% {
    background-color: rgba(99, 211, 114, 0.75);
  }
  50% {
    background-color: rgba(99, 211, 114, 0.3);
  }
  100% {
    background-color: rgba(99, 211, 114, 0.75);
  }
}

.samba-time-game-dance {
  display: flex;
  justify-content: center;
  align-items: center;
}

.dance-floor {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.1rem;

  background-color: rgba(0, 0, 0, 0.45);

  padding: 2rem;
  border-radius: 1rem;
}

.dance-floor-tile {
  width: 5rem;
  height: 5rem;
  border: solid 1px rgba(209, 209, 209, 0.2);

  display: flex;
  justify-content: center;
  align-items: center;
}

.dance-floor-tile-active {
  background-color: rgba(99, 211, 114, 0.75);
  animation: blink-animation 5s infinite; /* Initial blinking animation */
}

.dance-floor-shoe {
  height: 100%;
}
</style>

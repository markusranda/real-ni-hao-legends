<script setup lang="ts">
import { useGame } from '@/store/game'
import { computed, onMounted, ref, watch } from 'vue'

const players = computed(() => useGame().players)
const commands = computed(() => useGame().commands)
const drag = ref<HTMLDivElement>(null!)

function formatUUID(uuid: string) {
  return uuid.slice(0, 8) // Display only the first 8 characters of the UUID
}

function formatTime(time: number) {
  return new Date(time * 1000).toLocaleTimeString()
}

function getPlayerName(userId: string) {
  if (userId == localStorage.getItem('UserId')) {
    return 'You'
  }
  return players.value[userId].town.name
}

onMounted(() => {
  if (Notification.permission !== 'granted') {
    console.debug('Requesting permission...')
    Notification.requestPermission()
  }
})

watch(commands, (newCommands, oldCommands) => {
  const diff = newCommands.filter(
    (newCommand) =>
      !oldCommands.some((oldCommand) => JSON.stringify(oldCommand) === JSON.stringify(newCommand))
  )
  diff.forEach((command) => {
    if (command.type === 'init') {
      if (Notification.permission === 'granted') {
        const playerName = getPlayerName(command.userId)
        new Notification('New player has joined', {
          body: `${playerName} has joined the game`
        })
      }
    }
    if (command.type === 'chat') {
      if (Notification.permission === 'granted') {
        const playerName = players.value[command.userId].town.name
        new Notification('New message', {
          body: `${playerName}: ${command.data.message}`
        })
      }
    }
  })
})

onMounted(() => {
  const el = drag.value
  let isDown = false
  let offset = [0, 0]

  el.addEventListener(
    'mousedown',
    (event) => {
      isDown = true
      offset = [el.offsetLeft - event.clientX, el.offsetTop - event.clientY]
    },
    true
  )

  document.addEventListener(
    'mouseup',
    () => {
      isDown = false
    },
    true
  )

  document.addEventListener(
    'mousemove',
    (event) => {
      event.preventDefault()
      if (isDown) {
        el.style.left = event.clientX + offset[0] + 'px'
        el.style.top = event.clientY + offset[1] + 'px'
      }
    },
    true
  )
})
</script>

<template>
  <div>
    <h1>mandem on the net</h1>
    <p>there are {{ Object.keys(players).length }} opps round you</p>
    <li>
      <div v-for="player in players" :key="player.town.name" class="list-entry">
        <span>🟢 {{ player.town.name }}</span>
        <span>💰 {{ player.town.money }}</span>
      </div>
    </li>

    <ul class="debug-menu" ref="drag">
      <h6>DEBUG</h6>
      <li v-for="(command, index) in commands" :key="index" class="list-entry">
        <span> {{ command.type }}</span>
        <span> {{ getPlayerName(command.userId) }} </span>
        <span> {{ formatUUID(command.userId) }}</span>
        <span style="margin-left: 4px"> {{ formatTime(command.time) }}</span>
      </li>
    </ul>
  </div>
</template>

<style lang="scss" scoped>
$primary-color: #3498db;
$muted-color: #4b8586;

div {
  color: $primary-color;

  h1 {
    font-size: 1.6em;
    text-align: center;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
    font-weight: bold;
  }

  p {
    text-align: center;
    color: $muted-color;
  }

  span {
    display: block;
    font-size: .8em;
    line-height: 0.8em;
    letter-spacing: 0.05em;

    &:hover {
      color: lighten($muted-color, 20%); // Lighten the text color on hover
    }

    &:active {
      color: darken($muted-color, 20%); // Darken the text color on click
    }
  }
}

.list-entry {
  display: flex;
  justify-content: space-between;
  align-items: center;
  justify-items: center;

  padding: 0.7rem 1rem;
  transition: all 0.2s ease-in-out;

  gap: 4px;

  &:hover {
    background-color: #f2f2f2;
  }
}

.debug-menu {
  position: fixed;
  background-color: #fff;
  border: 1px solid #ccc;
  background-color: rgba(173, 255, 218, 0.32);
  color: black;
  border-radius: 5px;
  padding: 10px;
  font-size: 5px;
  font-family: monospace;
  z-index: 1000; // Ensure the menu appears above other elements

  // grid
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));

  h6 {
    text-align: center;
    font-size: 7px;
  }
}
</style>

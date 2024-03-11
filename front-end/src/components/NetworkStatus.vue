<template>
  <div v-if="networkStatus.error" class="network-error">
    Network Error: {{ networkStatus.errorMessage }}
  </div>
</template>

<script lang="ts" setup>
import { useWebsocket } from '@/store/websocketStore'
import { reactive } from 'vue'

const networkStatus = reactive({
  error: false,
  errorMessage: ''
})

const ws = useWebsocket().ws
if (!ws) throw Error("Can't get no networks status without ws")

ws.addEventListener('error', (error) => {
  console.error('WebSocket encountered error: ', error)
  networkStatus.error = true
  networkStatus.errorMessage = error.type
})

ws.addEventListener('close', () => {
  networkStatus.error = true
  networkStatus.errorMessage = 'WebSocket closed'
})
</script>

<style scoped>
.network-error {
  color: red;
  /* Other styles as needed */
}
</style>

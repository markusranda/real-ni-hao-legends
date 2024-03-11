<template>
  <div v-if="networkStatus.error" class="network-error">
    Network Error: {{ networkStatus.errorMessage }}
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { ws } from '@/main'

const networkStatus = reactive({
  error: false,
  errorMessage: ''
})

ws.addEventListener('error', (error) => {
  console.error('WebSocket encountered error: ', error)
  networkStatus.error = true
  networkStatus.errorMessage = error.type
})

ws.addEventListener('close', () => {
  console.log('WebSocket closed')
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

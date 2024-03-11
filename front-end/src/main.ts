import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { init } from '@/commands/init'
import { NHCommand } from './models/nh-command'

const app = createApp(App)
const ws = new WebSocket(`ws://${window.location.hostname}:8080/ws`)
ws.addEventListener('open', (event) => {
  console.log('WebSocket is open now.', event)
  console.log('Initializing game state')
  init()
})
ws.addEventListener('close', (event) => {
  console.log('WebSocket is closed now.', event)
})
ws.addEventListener('error', (error) => {
  console.error('WebSocket encountered error: ', error)
})
export { ws }

app.use(createPinia())

app.mount('#app')

export function send(message: NHCommand) {
  ws.send(JSON.stringify(message))
}
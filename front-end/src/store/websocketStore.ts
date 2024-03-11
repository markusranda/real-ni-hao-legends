import { NHCommand } from '@/models/nh-command'
import { GameState } from '@/models/nh-gamestate'
import { NHSamba } from '@/models/nh-samba'
import { defineStore } from 'pinia'
import { v4 as uuidv4 } from 'uuid'
import { useGame } from './game'

export const useWebsocket = defineStore('websocket', {
  state: () => ({
    loading: true,
    initDone: false,
    ws: undefined as WebSocket | undefined
  }),
  actions: {
    init() {
      const ws = new WebSocket(`ws://${window.location.hostname}:8080/ws`)

      if (!ws)
        throw Error("Couldn't establish a connection via websocket, what a sad time to be alive...")

      ws.addEventListener('open', (event) => {
        console.debug('WebSocket is open now.', event)
        this.loading = false
      })
      ws.addEventListener('close', (event) => {
        console.debug('WebSocket is closed now.', event)
      })
      ws.addEventListener('error', (error) => {
        console.error('WebSocket encountered error: ', error)
      })
      ws.onmessage = async (event) => {
        if (this.initDone) {
          useWebsocket().updateGameState(JSON.parse(event.data))
          this.loading = false
        } else this.authUser(event)
      }

      useWebsocket().ws = ws
    },
    authUser(event: MessageEvent<any> | undefined) {
      console.log('authUser1')

      if (!event) return

      let userId = localStorage.getItem('UserId')
      if (!userId) {
        userId = uuidv4()
        localStorage.setItem('UserId', userId)
      }
      console.log('authUser2', userId)

      if (!userId) return

      const data = JSON.parse(event.data) as { Players: Record<string, GameState> }
      const player = data.Players[userId]
      console.log('authUser3', player)

      if (player) {
        this.initDone = true
      }
    },
    initGameState(name: string) {
      let userId = localStorage.getItem('UserId')
      if (!userId) {
        userId = uuidv4()
        localStorage.setItem('UserId', userId)
      }
      const command = {
        type: 'init',
        userId: userId,
        data: {
          name: name
        }
      } as NHCommand

      this.send(command)
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

      console.log(partial)

      const state = { ...game.state, ...partial }
      const players = wsMessage.Players
      const samba = wsMessage.Samba
      delete players[userId]

      game.players = players
      game.chat = wsMessage.Chat.Messages
      game.state = state
      game.samba = samba
      game.commands = wsMessage.Commands
    },
    send(message: NHCommand) {
      this.ws?.send(JSON.stringify(message))
    }
  }
})

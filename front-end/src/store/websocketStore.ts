import {NHCommand, NHSendCommand} from '@/models/nh-command'
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
        } else {
          this.authUser(event)
        }

        this.loading = false
      }

      useWebsocket().ws = ws
    },
    authUser(event: MessageEvent<any> | undefined) {
      if (!event) return

      let userId = localStorage.getItem('UserId')
      if (!userId) {
        userId = uuidv4()
        localStorage.setItem('UserId', userId)
      }

      if (!userId) return

      const data = JSON.parse(event.data) as { Players: Record<string, GameState> }
      const player = data.Players[userId]

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
      Players?: Record<string, GameState>
      Samba?: NHSamba
      Chat?: { Messages: NhMessage[] }
      Commands?: NHCommand[]
    }) {
      const userId = localStorage.getItem('UserId')
      if (!userId) throw Error("Can't play the game without a name FOOoOOOoOoooOOol")

      const game = useGame()
      const partial = wsMessage?.Players?.[userId]

      const otherPlayers = wsMessage.Players
      delete otherPlayers?.[userId]

      console.log("hello")
      console.log(wsMessage.Commands)
      if (partial) game.state = { ...game.state, ...partial }
      if (otherPlayers) game.otherPlayers = otherPlayers
      if (wsMessage.Chat?.Messages) game.chat = wsMessage.Chat?.Messages
      if (wsMessage.Samba) game.samba = wsMessage.Samba
      if (wsMessage.Commands) game.commands = wsMessage.Commands
    },
    send(message: NHSendCommand) {
      (message as any).userId = localStorage.getItem('UserId')
      this.ws?.send(JSON.stringify(message))
    }
  }
})

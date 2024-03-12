import { defineStore } from 'pinia'
import { GameState } from '@/models/nh-gamestate'
import { NHSamba } from '@/models/nh-samba'
import { NHCommand } from '@/models/nh-command'

export const useGame = defineStore('game', () => {
  return {
    chat: [] as NhMessage[],
    state: {} as GameState,
    otherPlayers: {} as Record<string, GameState>,
    samba: {} as NHSamba,
    commands: [] as NHCommand[]
  }
})

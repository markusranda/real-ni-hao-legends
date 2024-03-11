import { v4 as uuidv4 } from 'uuid'
import { send, ws } from '@/main'
import { NHCommand } from '@/models/nh-command'

/**
 * Send an init message to the server to get your gamestate
 */
export const init = () => {
  let userId = localStorage.getItem('UserId')
  if (!userId) {
    userId = uuidv4()
    localStorage.setItem('UserId', userId)
  }

  console.assert(ws.readyState === ws.OPEN)
  console.assert(userId !== null && userId !== undefined)

  send({
    type: 'init',
    userId: userId,
    data: {
      name: 'Crookville'
    }
  } as NHCommand)
}

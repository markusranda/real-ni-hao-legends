import { NHCommand } from '@/models/nh-command'
import { useWebsocket } from '@/store/websocketStore'

export const buildingUpgrade = (key: string) => {
  const command = {
    userId: localStorage.getItem('UserId')!,
    type: 'building.upgrade',
    data: {
      buildingId: key
    }
  } as NHCommand

  useWebsocket().send(command)
}

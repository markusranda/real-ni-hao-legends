import { send } from '@/main'
import { NHCommand } from '@/models/nh-command'

export const buildingUpgrade = (key: string) => {
  const command = {
    userId: localStorage.getItem('UserId')!,
    type: 'building.upgrade',
    data: {
      buildingId: key
    }
  } as NHCommand

  send(command)
}

import { send } from '@/main'

export const buildingUpgrade = (key: string) => {
  const command = {
    userId: localStorage.getItem('UserId')!,
    type: 'building.upgrade',
    data: {
      buildingId: key
    }
  }

  console.log(command)

  send(command)
}

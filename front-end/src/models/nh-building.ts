export interface NHBuilding {
  key: string
  name: string
  imgUrl: string
  income: number
  level: number
  baseCost: number
  upgradeCost: number
}

export interface NHBuildingCommand {
  key: string
  action: string
}

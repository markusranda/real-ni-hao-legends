import type { NHBuilding } from "./nh-building"

export interface NHTown {
    money: number
    name: string
    buildings: Record<string, NHBuilding>
    happeningsHappened: Set<string>
}
import {NHTown} from "@/models/nh-town";
import {NHSamba} from "./nh-samba";
import {NHInventory} from "@/models/nh-inventory";

export interface GameState {
    retirementFund: number,
    town: NHTown
    samba: NHSamba
    inventory: NHInventory
}
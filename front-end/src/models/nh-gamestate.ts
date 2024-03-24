import {NHTown} from "@/models/nh-town";
import {NHSamba} from "./nh-samba";
import {NHInventory} from "@/models/nh-inventory";
import {NHEffect} from "@/models/NHEffect";

export interface GameState {
    retirementFund: number,
    town: NHTown
    samba: NHSamba
    inventory: NHInventory
    effects: NHEffect[]
}
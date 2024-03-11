import {NHTown} from "@/models/nh-town";
import { NHSamba } from "./nh-samba";

export interface GameState {
    retirementFund: number,
    town: NHTown
    samba: NHSamba
}
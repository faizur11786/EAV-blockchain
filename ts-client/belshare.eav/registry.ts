import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateUser } from "./types/belshare/eav/tx";
import { MsgCraeteAtteibute } from "./types/belshare/eav/tx";
import { MsgCraeteValue } from "./types/belshare/eav/tx";
import { MsgCreateShop } from "./types/belshare/eav/tx";
import { MsgCreateEntityType } from "./types/belshare/eav/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/belshare.eav.MsgCreateUser", MsgCreateUser],
    ["/belshare.eav.MsgCraeteAtteibute", MsgCraeteAtteibute],
    ["/belshare.eav.MsgCraeteValue", MsgCraeteValue],
    ["/belshare.eav.MsgCreateShop", MsgCreateShop],
    ["/belshare.eav.MsgCreateEntityType", MsgCreateEntityType],
    
];

export { msgTypes }
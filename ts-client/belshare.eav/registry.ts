import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateShop } from "./types/belshare/eav/tx";
import { MsgCreateEntityType } from "./types/belshare/eav/tx";
import { MsgCraeteValue } from "./types/belshare/eav/tx";
import { MsgCraeteAtteibute } from "./types/belshare/eav/tx";
import { MsgCreateUser } from "./types/belshare/eav/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/belshare.eav.MsgCreateShop", MsgCreateShop],
    ["/belshare.eav.MsgCreateEntityType", MsgCreateEntityType],
    ["/belshare.eav.MsgCraeteValue", MsgCraeteValue],
    ["/belshare.eav.MsgCraeteAtteibute", MsgCraeteAtteibute],
    ["/belshare.eav.MsgCreateUser", MsgCreateUser],
    
];

export { msgTypes }
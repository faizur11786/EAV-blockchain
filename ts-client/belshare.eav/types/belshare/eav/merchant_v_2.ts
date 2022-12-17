/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "belshare.eav";

export interface MerchantV2 {
  address: string;
  guid: string;
}

function createBaseMerchantV2(): MerchantV2 {
  return { address: "", guid: "" };
}

export const MerchantV2 = {
  encode(message: MerchantV2, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.guid !== "") {
      writer.uint32(18).string(message.guid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MerchantV2 {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMerchantV2();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.guid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MerchantV2 {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      guid: isSet(object.guid) ? String(object.guid) : "",
    };
  },

  toJSON(message: MerchantV2): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.guid !== undefined && (obj.guid = message.guid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MerchantV2>, I>>(object: I): MerchantV2 {
    const message = createBaseMerchantV2();
    message.address = object.address ?? "";
    message.guid = object.guid ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

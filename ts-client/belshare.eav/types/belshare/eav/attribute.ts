/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "belshare.eav";

export interface Attribute {
  guid: string;
  name: string;
  entityId: string;
}

function createBaseAttribute(): Attribute {
  return { guid: "", name: "", entityId: "" };
}

export const Attribute = {
  encode(message: Attribute, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.guid !== "") {
      writer.uint32(10).string(message.guid);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.entityId !== "") {
      writer.uint32(26).string(message.entityId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Attribute {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAttribute();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guid = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.entityId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Attribute {
    return {
      guid: isSet(object.guid) ? String(object.guid) : "",
      name: isSet(object.name) ? String(object.name) : "",
      entityId: isSet(object.entityId) ? String(object.entityId) : "",
    };
  },

  toJSON(message: Attribute): unknown {
    const obj: any = {};
    message.guid !== undefined && (obj.guid = message.guid);
    message.name !== undefined && (obj.name = message.name);
    message.entityId !== undefined && (obj.entityId = message.entityId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Attribute>, I>>(object: I): Attribute {
    const message = createBaseAttribute();
    message.guid = object.guid ?? "";
    message.name = object.name ?? "";
    message.entityId = object.entityId ?? "";
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

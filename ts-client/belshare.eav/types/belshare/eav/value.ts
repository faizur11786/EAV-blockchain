/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "belshare.eav";

export interface Value {
  entityId: string;
  attributeId: string;
  guid: string;
  value: string;
}

function createBaseValue(): Value {
  return { entityId: "", attributeId: "", guid: "", value: "" };
}

export const Value = {
  encode(message: Value, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.entityId !== "") {
      writer.uint32(10).string(message.entityId);
    }
    if (message.attributeId !== "") {
      writer.uint32(18).string(message.attributeId);
    }
    if (message.guid !== "") {
      writer.uint32(26).string(message.guid);
    }
    if (message.value !== "") {
      writer.uint32(34).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Value {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValue();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.entityId = reader.string();
          break;
        case 2:
          message.attributeId = reader.string();
          break;
        case 3:
          message.guid = reader.string();
          break;
        case 4:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Value {
    return {
      entityId: isSet(object.entityId) ? String(object.entityId) : "",
      attributeId: isSet(object.attributeId) ? String(object.attributeId) : "",
      guid: isSet(object.guid) ? String(object.guid) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: Value): unknown {
    const obj: any = {};
    message.entityId !== undefined && (obj.entityId = message.entityId);
    message.attributeId !== undefined && (obj.attributeId = message.attributeId);
    message.guid !== undefined && (obj.guid = message.guid);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Value>, I>>(object: I): Value {
    const message = createBaseValue();
    message.entityId = object.entityId ?? "";
    message.attributeId = object.attributeId ?? "";
    message.guid = object.guid ?? "";
    message.value = object.value ?? "";
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

/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "belshare.eav";

export interface AttributeResponse {
  label: string;
  value: string;
  attributeId: string;
  valueId: string;
}

function createBaseAttributeResponse(): AttributeResponse {
  return { label: "", value: "", attributeId: "", valueId: "" };
}

export const AttributeResponse = {
  encode(message: AttributeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    if (message.attributeId !== "") {
      writer.uint32(26).string(message.attributeId);
    }
    if (message.valueId !== "") {
      writer.uint32(34).string(message.valueId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AttributeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAttributeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.attributeId = reader.string();
          break;
        case 4:
          message.valueId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AttributeResponse {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : "",
      attributeId: isSet(object.attributeId) ? String(object.attributeId) : "",
      valueId: isSet(object.valueId) ? String(object.valueId) : "",
    };
  },

  toJSON(message: AttributeResponse): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    message.attributeId !== undefined && (obj.attributeId = message.attributeId);
    message.valueId !== undefined && (obj.valueId = message.valueId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AttributeResponse>, I>>(object: I): AttributeResponse {
    const message = createBaseAttributeResponse();
    message.label = object.label ?? "";
    message.value = object.value ?? "";
    message.attributeId = object.attributeId ?? "";
    message.valueId = object.valueId ?? "";
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

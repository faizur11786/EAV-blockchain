/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "belshare.user";

export interface User {
  creator: string;
  guid: string;
  address: string;
  createdAt: string;
}

function createBaseUser(): User {
  return { creator: "", guid: "", address: "", createdAt: "" };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.guid !== "") {
      writer.uint32(18).string(message.guid);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.createdAt !== "") {
      writer.uint32(34).string(message.createdAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.guid = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.createdAt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      guid: isSet(object.guid) ? String(object.guid) : "",
      address: isSet(object.address) ? String(object.address) : "",
      createdAt: isSet(object.createdAt) ? String(object.createdAt) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.guid !== undefined && (obj.guid = message.guid);
    message.address !== undefined && (obj.address = message.address);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.creator = object.creator ?? "";
    message.guid = object.guid ?? "";
    message.address = object.address ?? "";
    message.createdAt = object.createdAt ?? "";
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

/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Attribute } from "./attribute";

export const protobufPackage = "belshare.eav";

export interface MsgCreateUser {
  creator: string;
  address: string;
}

export interface MsgCreateUserResponse {
}

export interface MsgCreateEntityType {
  creator: string;
  name: string;
}

export interface MsgCreateEntityTypeResponse {
  id: string;
}

export interface MsgCraeteAtteibute {
  creator: string;
  name: string;
  entityId: string;
}

export interface MsgCraeteAtteibuteResponse {
  id: string;
}

export interface MsgCraeteValue {
  creator: string;
  attributeId: string;
  entityId: string;
  value: string;
}

export interface MsgCraeteValueResponse {
  id: string;
}

export interface MsgCreateShop {
  creator: string;
  address: string;
  attributs: Attribute[];
}

export interface MsgCreateShopResponse {
}

function createBaseMsgCreateUser(): MsgCreateUser {
  return { creator: "", address: "" };
}

export const MsgCreateUser = {
  encode(message: MsgCreateUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateUser {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgCreateUser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateUser>, I>>(object: I): MsgCreateUser {
    const message = createBaseMsgCreateUser();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgCreateUserResponse(): MsgCreateUserResponse {
  return {};
}

export const MsgCreateUserResponse = {
  encode(_: MsgCreateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateUserResponse {
    return {};
  },

  toJSON(_: MsgCreateUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateUserResponse>, I>>(_: I): MsgCreateUserResponse {
    const message = createBaseMsgCreateUserResponse();
    return message;
  },
};

function createBaseMsgCreateEntityType(): MsgCreateEntityType {
  return { creator: "", name: "" };
}

export const MsgCreateEntityType = {
  encode(message: MsgCreateEntityType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateEntityType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateEntityType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateEntityType {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: MsgCreateEntityType): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEntityType>, I>>(object: I): MsgCreateEntityType {
    const message = createBaseMsgCreateEntityType();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseMsgCreateEntityTypeResponse(): MsgCreateEntityTypeResponse {
  return { id: "" };
}

export const MsgCreateEntityTypeResponse = {
  encode(message: MsgCreateEntityTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateEntityTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateEntityTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateEntityTypeResponse {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: MsgCreateEntityTypeResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEntityTypeResponse>, I>>(object: I): MsgCreateEntityTypeResponse {
    const message = createBaseMsgCreateEntityTypeResponse();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgCraeteAtteibute(): MsgCraeteAtteibute {
  return { creator: "", name: "", entityId: "" };
}

export const MsgCraeteAtteibute = {
  encode(message: MsgCraeteAtteibute, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.entityId !== "") {
      writer.uint32(26).string(message.entityId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCraeteAtteibute {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCraeteAtteibute();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgCraeteAtteibute {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      entityId: isSet(object.entityId) ? String(object.entityId) : "",
    };
  },

  toJSON(message: MsgCraeteAtteibute): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.entityId !== undefined && (obj.entityId = message.entityId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCraeteAtteibute>, I>>(object: I): MsgCraeteAtteibute {
    const message = createBaseMsgCraeteAtteibute();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.entityId = object.entityId ?? "";
    return message;
  },
};

function createBaseMsgCraeteAtteibuteResponse(): MsgCraeteAtteibuteResponse {
  return { id: "" };
}

export const MsgCraeteAtteibuteResponse = {
  encode(message: MsgCraeteAtteibuteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCraeteAtteibuteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCraeteAtteibuteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCraeteAtteibuteResponse {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: MsgCraeteAtteibuteResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCraeteAtteibuteResponse>, I>>(object: I): MsgCraeteAtteibuteResponse {
    const message = createBaseMsgCraeteAtteibuteResponse();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgCraeteValue(): MsgCraeteValue {
  return { creator: "", attributeId: "", entityId: "", value: "" };
}

export const MsgCraeteValue = {
  encode(message: MsgCraeteValue, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.attributeId !== "") {
      writer.uint32(18).string(message.attributeId);
    }
    if (message.entityId !== "") {
      writer.uint32(26).string(message.entityId);
    }
    if (message.value !== "") {
      writer.uint32(34).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCraeteValue {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCraeteValue();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.attributeId = reader.string();
          break;
        case 3:
          message.entityId = reader.string();
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

  fromJSON(object: any): MsgCraeteValue {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      attributeId: isSet(object.attributeId) ? String(object.attributeId) : "",
      entityId: isSet(object.entityId) ? String(object.entityId) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MsgCraeteValue): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.attributeId !== undefined && (obj.attributeId = message.attributeId);
    message.entityId !== undefined && (obj.entityId = message.entityId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCraeteValue>, I>>(object: I): MsgCraeteValue {
    const message = createBaseMsgCraeteValue();
    message.creator = object.creator ?? "";
    message.attributeId = object.attributeId ?? "";
    message.entityId = object.entityId ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgCraeteValueResponse(): MsgCraeteValueResponse {
  return { id: "" };
}

export const MsgCraeteValueResponse = {
  encode(message: MsgCraeteValueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCraeteValueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCraeteValueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCraeteValueResponse {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: MsgCraeteValueResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCraeteValueResponse>, I>>(object: I): MsgCraeteValueResponse {
    const message = createBaseMsgCraeteValueResponse();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgCreateShop(): MsgCreateShop {
  return { creator: "", address: "", attributs: [] };
}

export const MsgCreateShop = {
  encode(message: MsgCreateShop, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    for (const v of message.attributs) {
      Attribute.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateShop {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateShop();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.attributs.push(Attribute.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateShop {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      attributs: Array.isArray(object?.attributs) ? object.attributs.map((e: any) => Attribute.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgCreateShop): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    if (message.attributs) {
      obj.attributs = message.attributs.map((e) => e ? Attribute.toJSON(e) : undefined);
    } else {
      obj.attributs = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateShop>, I>>(object: I): MsgCreateShop {
    const message = createBaseMsgCreateShop();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.attributs = object.attributs?.map((e) => Attribute.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgCreateShopResponse(): MsgCreateShopResponse {
  return {};
}

export const MsgCreateShopResponse = {
  encode(_: MsgCreateShopResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateShopResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateShopResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateShopResponse {
    return {};
  },

  toJSON(_: MsgCreateShopResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateShopResponse>, I>>(_: I): MsgCreateShopResponse {
    const message = createBaseMsgCreateShopResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateUser(request: MsgCreateUser): Promise<MsgCreateUserResponse>;
  CreateEntityType(request: MsgCreateEntityType): Promise<MsgCreateEntityTypeResponse>;
  CraeteAtteibute(request: MsgCraeteAtteibute): Promise<MsgCraeteAtteibuteResponse>;
  CraeteValue(request: MsgCraeteValue): Promise<MsgCraeteValueResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateShop(request: MsgCreateShop): Promise<MsgCreateShopResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateUser = this.CreateUser.bind(this);
    this.CreateEntityType = this.CreateEntityType.bind(this);
    this.CraeteAtteibute = this.CraeteAtteibute.bind(this);
    this.CraeteValue = this.CraeteValue.bind(this);
    this.CreateShop = this.CreateShop.bind(this);
  }
  CreateUser(request: MsgCreateUser): Promise<MsgCreateUserResponse> {
    const data = MsgCreateUser.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Msg", "CreateUser", data);
    return promise.then((data) => MsgCreateUserResponse.decode(new _m0.Reader(data)));
  }

  CreateEntityType(request: MsgCreateEntityType): Promise<MsgCreateEntityTypeResponse> {
    const data = MsgCreateEntityType.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Msg", "CreateEntityType", data);
    return promise.then((data) => MsgCreateEntityTypeResponse.decode(new _m0.Reader(data)));
  }

  CraeteAtteibute(request: MsgCraeteAtteibute): Promise<MsgCraeteAtteibuteResponse> {
    const data = MsgCraeteAtteibute.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Msg", "CraeteAtteibute", data);
    return promise.then((data) => MsgCraeteAtteibuteResponse.decode(new _m0.Reader(data)));
  }

  CraeteValue(request: MsgCraeteValue): Promise<MsgCraeteValueResponse> {
    const data = MsgCraeteValue.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Msg", "CraeteValue", data);
    return promise.then((data) => MsgCraeteValueResponse.decode(new _m0.Reader(data)));
  }

  CreateShop(request: MsgCreateShop): Promise<MsgCreateShopResponse> {
    const data = MsgCreateShop.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Msg", "CreateShop", data);
    return promise.then((data) => MsgCreateShopResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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

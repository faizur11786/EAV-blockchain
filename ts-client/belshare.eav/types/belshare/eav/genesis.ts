/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Attribute } from "./attribute";
import { EntityType } from "./entity_type";
import { Merchant } from "./merchant";
import { MerchantNew } from "./merchant_new";
import { Params } from "./params";
import { User } from "./user";
import { Value } from "./value";

export const protobufPackage = "belshare.eav";

/** GenesisState defines the eav module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  userList: User[];
  merchantList: Merchant[];
  entityTypeList: EntityType[];
  attributeList: Attribute[];
  valueList: Value[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  merchantNewList: MerchantNew[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    userList: [],
    merchantList: [],
    entityTypeList: [],
    attributeList: [],
    valueList: [],
    merchantNewList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.userList) {
      User.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.merchantList) {
      Merchant.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.entityTypeList) {
      EntityType.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.attributeList) {
      Attribute.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.valueList) {
      Value.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.merchantNewList) {
      MerchantNew.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.userList.push(User.decode(reader, reader.uint32()));
          break;
        case 3:
          message.merchantList.push(Merchant.decode(reader, reader.uint32()));
          break;
        case 4:
          message.entityTypeList.push(EntityType.decode(reader, reader.uint32()));
          break;
        case 5:
          message.attributeList.push(Attribute.decode(reader, reader.uint32()));
          break;
        case 6:
          message.valueList.push(Value.decode(reader, reader.uint32()));
          break;
        case 7:
          message.merchantNewList.push(MerchantNew.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      userList: Array.isArray(object?.userList) ? object.userList.map((e: any) => User.fromJSON(e)) : [],
      merchantList: Array.isArray(object?.merchantList)
        ? object.merchantList.map((e: any) => Merchant.fromJSON(e))
        : [],
      entityTypeList: Array.isArray(object?.entityTypeList)
        ? object.entityTypeList.map((e: any) => EntityType.fromJSON(e))
        : [],
      attributeList: Array.isArray(object?.attributeList)
        ? object.attributeList.map((e: any) => Attribute.fromJSON(e))
        : [],
      valueList: Array.isArray(object?.valueList) ? object.valueList.map((e: any) => Value.fromJSON(e)) : [],
      merchantNewList: Array.isArray(object?.merchantNewList)
        ? object.merchantNewList.map((e: any) => MerchantNew.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.userList) {
      obj.userList = message.userList.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.userList = [];
    }
    if (message.merchantList) {
      obj.merchantList = message.merchantList.map((e) => e ? Merchant.toJSON(e) : undefined);
    } else {
      obj.merchantList = [];
    }
    if (message.entityTypeList) {
      obj.entityTypeList = message.entityTypeList.map((e) => e ? EntityType.toJSON(e) : undefined);
    } else {
      obj.entityTypeList = [];
    }
    if (message.attributeList) {
      obj.attributeList = message.attributeList.map((e) => e ? Attribute.toJSON(e) : undefined);
    } else {
      obj.attributeList = [];
    }
    if (message.valueList) {
      obj.valueList = message.valueList.map((e) => e ? Value.toJSON(e) : undefined);
    } else {
      obj.valueList = [];
    }
    if (message.merchantNewList) {
      obj.merchantNewList = message.merchantNewList.map((e) => e ? MerchantNew.toJSON(e) : undefined);
    } else {
      obj.merchantNewList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.userList = object.userList?.map((e) => User.fromPartial(e)) || [];
    message.merchantList = object.merchantList?.map((e) => Merchant.fromPartial(e)) || [];
    message.entityTypeList = object.entityTypeList?.map((e) => EntityType.fromPartial(e)) || [];
    message.attributeList = object.attributeList?.map((e) => Attribute.fromPartial(e)) || [];
    message.valueList = object.valueList?.map((e) => Value.fromPartial(e)) || [];
    message.merchantNewList = object.merchantNewList?.map((e) => MerchantNew.fromPartial(e)) || [];
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

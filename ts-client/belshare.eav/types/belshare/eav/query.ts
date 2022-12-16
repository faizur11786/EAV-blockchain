/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Attribute } from "./attribute";
import { EntityType } from "./entity_type";
import { Merchant } from "./merchant";
import { Params } from "./params";
import { User } from "./user";
import { Value } from "./value";

export const protobufPackage = "belshare.eav";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetUserRequest {
  index: string;
}

export interface QueryGetUserResponse {
  user: User | undefined;
}

export interface QueryAllUserRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserResponse {
  user: User[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMerchantRequest {
  guid: string;
  creator: string;
}

export interface QueryGetMerchantResponse {
  merchant: Merchant | undefined;
}

export interface QueryAllMerchantRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMerchantResponse {
  merchant: Merchant[];
  pagination: PageResponse | undefined;
}

export interface QueryGetEntityTypeRequest {
  guid: string;
}

export interface QueryGetEntityTypeResponse {
  entityType: EntityType | undefined;
}

export interface QueryAllEntityTypeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEntityTypeResponse {
  entityType: EntityType[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAttributeRequest {
  guid: string;
}

export interface QueryGetAttributeResponse {
  attribute: Attribute | undefined;
}

export interface QueryAllAttributeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAttributeResponse {
  attribute: Attribute[];
  pagination: PageResponse | undefined;
}

export interface QueryGetValueRequest {
  entityId: string;
  attributeId: string;
}

export interface QueryGetValueResponse {
  value: Value | undefined;
}

export interface QueryAllValueRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllValueResponse {
  value: Value[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserRequest(): QueryGetUserRequest {
  return { index: "" };
}

export const QueryGetUserRequest = {
  encode(message: QueryGetUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetUserRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserRequest>, I>>(object: I): QueryGetUserRequest {
    const message = createBaseQueryGetUserRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetUserResponse(): QueryGetUserResponse {
  return { user: undefined };
}

export const QueryGetUserResponse = {
  encode(message: QueryGetUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: QueryGetUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserResponse>, I>>(object: I): QueryGetUserResponse {
    const message = createBaseQueryGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseQueryAllUserRequest(): QueryAllUserRequest {
  return { pagination: undefined };
}

export const QueryAllUserRequest = {
  encode(message: QueryAllUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRequest>, I>>(object: I): QueryAllUserRequest {
    const message = createBaseQueryAllUserRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserResponse(): QueryAllUserResponse {
  return { user: [], pagination: undefined };
}

export const QueryAllUserResponse = {
  encode(message: QueryAllUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.user) {
      User.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user.push(User.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserResponse {
    return {
      user: Array.isArray(object?.user) ? object.user.map((e: any) => User.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserResponse): unknown {
    const obj: any = {};
    if (message.user) {
      obj.user = message.user.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.user = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserResponse>, I>>(object: I): QueryAllUserResponse {
    const message = createBaseQueryAllUserResponse();
    message.user = object.user?.map((e) => User.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetMerchantRequest(): QueryGetMerchantRequest {
  return { guid: "", creator: "" };
}

export const QueryGetMerchantRequest = {
  encode(message: QueryGetMerchantRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.guid !== "") {
      writer.uint32(10).string(message.guid);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMerchantRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMerchantRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guid = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMerchantRequest {
    return {
      guid: isSet(object.guid) ? String(object.guid) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: QueryGetMerchantRequest): unknown {
    const obj: any = {};
    message.guid !== undefined && (obj.guid = message.guid);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMerchantRequest>, I>>(object: I): QueryGetMerchantRequest {
    const message = createBaseQueryGetMerchantRequest();
    message.guid = object.guid ?? "";
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseQueryGetMerchantResponse(): QueryGetMerchantResponse {
  return { merchant: undefined };
}

export const QueryGetMerchantResponse = {
  encode(message: QueryGetMerchantResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.merchant !== undefined) {
      Merchant.encode(message.merchant, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMerchantResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMerchantResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.merchant = Merchant.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMerchantResponse {
    return { merchant: isSet(object.merchant) ? Merchant.fromJSON(object.merchant) : undefined };
  },

  toJSON(message: QueryGetMerchantResponse): unknown {
    const obj: any = {};
    message.merchant !== undefined && (obj.merchant = message.merchant ? Merchant.toJSON(message.merchant) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMerchantResponse>, I>>(object: I): QueryGetMerchantResponse {
    const message = createBaseQueryGetMerchantResponse();
    message.merchant = (object.merchant !== undefined && object.merchant !== null)
      ? Merchant.fromPartial(object.merchant)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMerchantRequest(): QueryAllMerchantRequest {
  return { pagination: undefined };
}

export const QueryAllMerchantRequest = {
  encode(message: QueryAllMerchantRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMerchantRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMerchantRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMerchantRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllMerchantRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMerchantRequest>, I>>(object: I): QueryAllMerchantRequest {
    const message = createBaseQueryAllMerchantRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMerchantResponse(): QueryAllMerchantResponse {
  return { merchant: [], pagination: undefined };
}

export const QueryAllMerchantResponse = {
  encode(message: QueryAllMerchantResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.merchant) {
      Merchant.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMerchantResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMerchantResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.merchant.push(Merchant.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMerchantResponse {
    return {
      merchant: Array.isArray(object?.merchant) ? object.merchant.map((e: any) => Merchant.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllMerchantResponse): unknown {
    const obj: any = {};
    if (message.merchant) {
      obj.merchant = message.merchant.map((e) => e ? Merchant.toJSON(e) : undefined);
    } else {
      obj.merchant = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMerchantResponse>, I>>(object: I): QueryAllMerchantResponse {
    const message = createBaseQueryAllMerchantResponse();
    message.merchant = object.merchant?.map((e) => Merchant.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetEntityTypeRequest(): QueryGetEntityTypeRequest {
  return { guid: "" };
}

export const QueryGetEntityTypeRequest = {
  encode(message: QueryGetEntityTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.guid !== "") {
      writer.uint32(10).string(message.guid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEntityTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEntityTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEntityTypeRequest {
    return { guid: isSet(object.guid) ? String(object.guid) : "" };
  },

  toJSON(message: QueryGetEntityTypeRequest): unknown {
    const obj: any = {};
    message.guid !== undefined && (obj.guid = message.guid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEntityTypeRequest>, I>>(object: I): QueryGetEntityTypeRequest {
    const message = createBaseQueryGetEntityTypeRequest();
    message.guid = object.guid ?? "";
    return message;
  },
};

function createBaseQueryGetEntityTypeResponse(): QueryGetEntityTypeResponse {
  return { entityType: undefined };
}

export const QueryGetEntityTypeResponse = {
  encode(message: QueryGetEntityTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.entityType !== undefined) {
      EntityType.encode(message.entityType, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEntityTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEntityTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.entityType = EntityType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEntityTypeResponse {
    return { entityType: isSet(object.entityType) ? EntityType.fromJSON(object.entityType) : undefined };
  },

  toJSON(message: QueryGetEntityTypeResponse): unknown {
    const obj: any = {};
    message.entityType !== undefined
      && (obj.entityType = message.entityType ? EntityType.toJSON(message.entityType) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEntityTypeResponse>, I>>(object: I): QueryGetEntityTypeResponse {
    const message = createBaseQueryGetEntityTypeResponse();
    message.entityType = (object.entityType !== undefined && object.entityType !== null)
      ? EntityType.fromPartial(object.entityType)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEntityTypeRequest(): QueryAllEntityTypeRequest {
  return { pagination: undefined };
}

export const QueryAllEntityTypeRequest = {
  encode(message: QueryAllEntityTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEntityTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEntityTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEntityTypeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEntityTypeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEntityTypeRequest>, I>>(object: I): QueryAllEntityTypeRequest {
    const message = createBaseQueryAllEntityTypeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEntityTypeResponse(): QueryAllEntityTypeResponse {
  return { entityType: [], pagination: undefined };
}

export const QueryAllEntityTypeResponse = {
  encode(message: QueryAllEntityTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.entityType) {
      EntityType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEntityTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEntityTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.entityType.push(EntityType.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEntityTypeResponse {
    return {
      entityType: Array.isArray(object?.entityType) ? object.entityType.map((e: any) => EntityType.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEntityTypeResponse): unknown {
    const obj: any = {};
    if (message.entityType) {
      obj.entityType = message.entityType.map((e) => e ? EntityType.toJSON(e) : undefined);
    } else {
      obj.entityType = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEntityTypeResponse>, I>>(object: I): QueryAllEntityTypeResponse {
    const message = createBaseQueryAllEntityTypeResponse();
    message.entityType = object.entityType?.map((e) => EntityType.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetAttributeRequest(): QueryGetAttributeRequest {
  return { guid: "" };
}

export const QueryGetAttributeRequest = {
  encode(message: QueryGetAttributeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.guid !== "") {
      writer.uint32(10).string(message.guid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAttributeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAttributeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAttributeRequest {
    return { guid: isSet(object.guid) ? String(object.guid) : "" };
  },

  toJSON(message: QueryGetAttributeRequest): unknown {
    const obj: any = {};
    message.guid !== undefined && (obj.guid = message.guid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAttributeRequest>, I>>(object: I): QueryGetAttributeRequest {
    const message = createBaseQueryGetAttributeRequest();
    message.guid = object.guid ?? "";
    return message;
  },
};

function createBaseQueryGetAttributeResponse(): QueryGetAttributeResponse {
  return { attribute: undefined };
}

export const QueryGetAttributeResponse = {
  encode(message: QueryGetAttributeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.attribute !== undefined) {
      Attribute.encode(message.attribute, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAttributeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAttributeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.attribute = Attribute.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAttributeResponse {
    return { attribute: isSet(object.attribute) ? Attribute.fromJSON(object.attribute) : undefined };
  },

  toJSON(message: QueryGetAttributeResponse): unknown {
    const obj: any = {};
    message.attribute !== undefined
      && (obj.attribute = message.attribute ? Attribute.toJSON(message.attribute) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAttributeResponse>, I>>(object: I): QueryGetAttributeResponse {
    const message = createBaseQueryGetAttributeResponse();
    message.attribute = (object.attribute !== undefined && object.attribute !== null)
      ? Attribute.fromPartial(object.attribute)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAttributeRequest(): QueryAllAttributeRequest {
  return { pagination: undefined };
}

export const QueryAllAttributeRequest = {
  encode(message: QueryAllAttributeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAttributeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAttributeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAttributeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllAttributeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllAttributeRequest>, I>>(object: I): QueryAllAttributeRequest {
    const message = createBaseQueryAllAttributeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllAttributeResponse(): QueryAllAttributeResponse {
  return { attribute: [], pagination: undefined };
}

export const QueryAllAttributeResponse = {
  encode(message: QueryAllAttributeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.attribute) {
      Attribute.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllAttributeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllAttributeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.attribute.push(Attribute.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAttributeResponse {
    return {
      attribute: Array.isArray(object?.attribute) ? object.attribute.map((e: any) => Attribute.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllAttributeResponse): unknown {
    const obj: any = {};
    if (message.attribute) {
      obj.attribute = message.attribute.map((e) => e ? Attribute.toJSON(e) : undefined);
    } else {
      obj.attribute = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllAttributeResponse>, I>>(object: I): QueryAllAttributeResponse {
    const message = createBaseQueryAllAttributeResponse();
    message.attribute = object.attribute?.map((e) => Attribute.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetValueRequest(): QueryGetValueRequest {
  return { entityId: "", attributeId: "" };
}

export const QueryGetValueRequest = {
  encode(message: QueryGetValueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.entityId !== "") {
      writer.uint32(10).string(message.entityId);
    }
    if (message.attributeId !== "") {
      writer.uint32(18).string(message.attributeId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.entityId = reader.string();
          break;
        case 2:
          message.attributeId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValueRequest {
    return {
      entityId: isSet(object.entityId) ? String(object.entityId) : "",
      attributeId: isSet(object.attributeId) ? String(object.attributeId) : "",
    };
  },

  toJSON(message: QueryGetValueRequest): unknown {
    const obj: any = {};
    message.entityId !== undefined && (obj.entityId = message.entityId);
    message.attributeId !== undefined && (obj.attributeId = message.attributeId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValueRequest>, I>>(object: I): QueryGetValueRequest {
    const message = createBaseQueryGetValueRequest();
    message.entityId = object.entityId ?? "";
    message.attributeId = object.attributeId ?? "";
    return message;
  },
};

function createBaseQueryGetValueResponse(): QueryGetValueResponse {
  return { value: undefined };
}

export const QueryGetValueResponse = {
  encode(message: QueryGetValueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.value !== undefined) {
      Value.encode(message.value, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.value = Value.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValueResponse {
    return { value: isSet(object.value) ? Value.fromJSON(object.value) : undefined };
  },

  toJSON(message: QueryGetValueResponse): unknown {
    const obj: any = {};
    message.value !== undefined && (obj.value = message.value ? Value.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValueResponse>, I>>(object: I): QueryGetValueResponse {
    const message = createBaseQueryGetValueResponse();
    message.value = (object.value !== undefined && object.value !== null) ? Value.fromPartial(object.value) : undefined;
    return message;
  },
};

function createBaseQueryAllValueRequest(): QueryAllValueRequest {
  return { pagination: undefined };
}

export const QueryAllValueRequest = {
  encode(message: QueryAllValueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllValueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllValueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllValueRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllValueRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllValueRequest>, I>>(object: I): QueryAllValueRequest {
    const message = createBaseQueryAllValueRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllValueResponse(): QueryAllValueResponse {
  return { value: [], pagination: undefined };
}

export const QueryAllValueResponse = {
  encode(message: QueryAllValueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.value) {
      Value.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllValueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllValueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.value.push(Value.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllValueResponse {
    return {
      value: Array.isArray(object?.value) ? object.value.map((e: any) => Value.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllValueResponse): unknown {
    const obj: any = {};
    if (message.value) {
      obj.value = message.value.map((e) => e ? Value.toJSON(e) : undefined);
    } else {
      obj.value = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllValueResponse>, I>>(object: I): QueryAllValueResponse {
    const message = createBaseQueryAllValueResponse();
    message.value = object.value?.map((e) => Value.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a User by index. */
  User(request: QueryGetUserRequest): Promise<QueryGetUserResponse>;
  /** Queries a list of User items. */
  UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse>;
  /** Queries a Merchant by index. */
  Merchant(request: QueryGetMerchantRequest): Promise<QueryGetMerchantResponse>;
  /** Queries a list of Merchant items. */
  MerchantAll(request: QueryAllMerchantRequest): Promise<QueryAllMerchantResponse>;
  /** Queries a EntityType by index. */
  EntityType(request: QueryGetEntityTypeRequest): Promise<QueryGetEntityTypeResponse>;
  /** Queries a list of EntityType items. */
  EntityTypeAll(request: QueryAllEntityTypeRequest): Promise<QueryAllEntityTypeResponse>;
  /** Queries a Attribute by index. */
  Attribute(request: QueryGetAttributeRequest): Promise<QueryGetAttributeResponse>;
  /** Queries a list of Attribute items. */
  AttributeAll(request: QueryAllAttributeRequest): Promise<QueryAllAttributeResponse>;
  /** Queries a Value by index. */
  Value(request: QueryGetValueRequest): Promise<QueryGetValueResponse>;
  /** Queries a list of Value items. */
  ValueAll(request: QueryAllValueRequest): Promise<QueryAllValueResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.User = this.User.bind(this);
    this.UserAll = this.UserAll.bind(this);
    this.Merchant = this.Merchant.bind(this);
    this.MerchantAll = this.MerchantAll.bind(this);
    this.EntityType = this.EntityType.bind(this);
    this.EntityTypeAll = this.EntityTypeAll.bind(this);
    this.Attribute = this.Attribute.bind(this);
    this.AttributeAll = this.AttributeAll.bind(this);
    this.Value = this.Value.bind(this);
    this.ValueAll = this.ValueAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  User(request: QueryGetUserRequest): Promise<QueryGetUserResponse> {
    const data = QueryGetUserRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "User", data);
    return promise.then((data) => QueryGetUserResponse.decode(new _m0.Reader(data)));
  }

  UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse> {
    const data = QueryAllUserRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "UserAll", data);
    return promise.then((data) => QueryAllUserResponse.decode(new _m0.Reader(data)));
  }

  Merchant(request: QueryGetMerchantRequest): Promise<QueryGetMerchantResponse> {
    const data = QueryGetMerchantRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "Merchant", data);
    return promise.then((data) => QueryGetMerchantResponse.decode(new _m0.Reader(data)));
  }

  MerchantAll(request: QueryAllMerchantRequest): Promise<QueryAllMerchantResponse> {
    const data = QueryAllMerchantRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "MerchantAll", data);
    return promise.then((data) => QueryAllMerchantResponse.decode(new _m0.Reader(data)));
  }

  EntityType(request: QueryGetEntityTypeRequest): Promise<QueryGetEntityTypeResponse> {
    const data = QueryGetEntityTypeRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "EntityType", data);
    return promise.then((data) => QueryGetEntityTypeResponse.decode(new _m0.Reader(data)));
  }

  EntityTypeAll(request: QueryAllEntityTypeRequest): Promise<QueryAllEntityTypeResponse> {
    const data = QueryAllEntityTypeRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "EntityTypeAll", data);
    return promise.then((data) => QueryAllEntityTypeResponse.decode(new _m0.Reader(data)));
  }

  Attribute(request: QueryGetAttributeRequest): Promise<QueryGetAttributeResponse> {
    const data = QueryGetAttributeRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "Attribute", data);
    return promise.then((data) => QueryGetAttributeResponse.decode(new _m0.Reader(data)));
  }

  AttributeAll(request: QueryAllAttributeRequest): Promise<QueryAllAttributeResponse> {
    const data = QueryAllAttributeRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "AttributeAll", data);
    return promise.then((data) => QueryAllAttributeResponse.decode(new _m0.Reader(data)));
  }

  Value(request: QueryGetValueRequest): Promise<QueryGetValueResponse> {
    const data = QueryGetValueRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "Value", data);
    return promise.then((data) => QueryGetValueResponse.decode(new _m0.Reader(data)));
  }

  ValueAll(request: QueryAllValueRequest): Promise<QueryAllValueResponse> {
    const data = QueryAllValueRequest.encode(request).finish();
    const promise = this.rpc.request("belshare.eav.Query", "ValueAll", data);
    return promise.then((data) => QueryAllValueResponse.decode(new _m0.Reader(data)));
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

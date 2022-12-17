/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface EavAttribute {
  guid?: string;
  name?: string;
  entityId?: string;
}

export interface EavEntityType {
  guid?: string;
  name?: string;
  creator?: string;
}

export interface EavMerchant {
  guid?: string;
  creator?: string;
  address?: string;
}

export interface EavMerchantNew {
  address?: string;
  guid?: string;
  creator?: string;
}

export interface EavMsgCraeteAtteibuteResponse {
  id?: string;
}

export interface EavMsgCraeteValueResponse {
  id?: string;
}

export interface EavMsgCreateEntityTypeResponse {
  id?: string;
}

export type EavMsgCreateShopResponse = object;

export type EavMsgCreateUserResponse = object;

export interface EavMsgNewMerchantResponse {
  guid?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type EavParams = object;

export interface EavQueryAllAttributeResponse {
  attribute?: EavAttribute[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface EavQueryAllEntityTypeResponse {
  entityType?: EavEntityType[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface EavQueryAllMerchantNewResponse {
  merchantNew?: EavMerchantNew[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface EavQueryAllMerchantResponse {
  merchant?: EavMerchant[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface EavQueryAllUserResponse {
  user?: EavUser[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface EavQueryAllValueResponse {
  value?: EavValue[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface EavQueryGetAttributeResponse {
  attribute?: EavAttribute;
}

export interface EavQueryGetEntityTypeResponse {
  entityType?: EavEntityType;
}

export interface EavQueryGetMerchantNewResponse {
  merchantNew?: EavMerchantNew;
}

export interface EavQueryGetMerchantResponse {
  merchant?: EavMerchant;
}

export interface EavQueryGetUserResponse {
  user?: EavUser;
}

export interface EavQueryGetValueResponse {
  value?: EavValue;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface EavQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: EavParams;
}

export interface EavUser {
  index?: string;
  creator?: string;

  /** @format int64 */
  createdAt?: string;
}

export interface EavValue {
  entityId?: string;
  attributeId?: string;
  guid?: string;
  value?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title belshare/eav/attribute.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAttributeAll
   * @summary Queries a list of Attribute items.
   * @request GET:/belShare/eav/attribute
   */
  queryAttributeAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<EavQueryAllAttributeResponse, RpcStatus>({
      path: `/belShare/eav/attribute`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAttribute
   * @summary Queries a Attribute by index.
   * @request GET:/belShare/eav/attribute/{guid}
   */
  queryAttribute = (guid: string, params: RequestParams = {}) =>
    this.request<EavQueryGetAttributeResponse, RpcStatus>({
      path: `/belShare/eav/attribute/${guid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEntityTypeAll
   * @summary Queries a list of EntityType items.
   * @request GET:/belShare/eav/entity_type
   */
  queryEntityTypeAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<EavQueryAllEntityTypeResponse, RpcStatus>({
      path: `/belShare/eav/entity_type`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEntityType
   * @summary Queries a EntityType by index.
   * @request GET:/belShare/eav/entity_type/{guid}
   */
  queryEntityType = (guid: string, params: RequestParams = {}) =>
    this.request<EavQueryGetEntityTypeResponse, RpcStatus>({
      path: `/belShare/eav/entity_type/${guid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMerchantAll
   * @summary Queries a list of Merchant items.
   * @request GET:/belShare/eav/merchant
   */
  queryMerchantAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<EavQueryAllMerchantResponse, RpcStatus>({
      path: `/belShare/eav/merchant`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMerchant
   * @summary Queries a Merchant by index.
   * @request GET:/belShare/eav/merchant/{guid}/{creator}
   */
  queryMerchant = (guid: string, creator: string, params: RequestParams = {}) =>
    this.request<EavQueryGetMerchantResponse, RpcStatus>({
      path: `/belShare/eav/merchant/${guid}/${creator}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMerchantNewAll
   * @summary Queries a list of MerchantNew items.
   * @request GET:/belShare/eav/merchant_new
   */
  queryMerchantNewAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<EavQueryAllMerchantNewResponse, RpcStatus>({
      path: `/belShare/eav/merchant_new`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMerchantNew
   * @summary Queries a MerchantNew by index.
   * @request GET:/belShare/eav/merchant_new/{address}
   */
  queryMerchantNew = (address: string, params: RequestParams = {}) =>
    this.request<EavQueryGetMerchantNewResponse, RpcStatus>({
      path: `/belShare/eav/merchant_new/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/belShare/eav/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<EavQueryParamsResponse, RpcStatus>({
      path: `/belShare/eav/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserAll
   * @summary Queries a list of User items.
   * @request GET:/belShare/eav/user
   */
  queryUserAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<EavQueryAllUserResponse, RpcStatus>({
      path: `/belShare/eav/user`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUser
   * @summary Queries a User by index.
   * @request GET:/belShare/eav/user/{index}
   */
  queryUser = (index: string, params: RequestParams = {}) =>
    this.request<EavQueryGetUserResponse, RpcStatus>({
      path: `/belShare/eav/user/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValueAll
   * @summary Queries a list of Value items.
   * @request GET:/belShare/eav/value
   */
  queryValueAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<EavQueryAllValueResponse, RpcStatus>({
      path: `/belShare/eav/value`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValue
   * @summary Queries a Value by index.
   * @request GET:/belShare/eav/value/{entityId}/{attributeId}
   */
  queryValue = (entityId: string, attributeId: string, params: RequestParams = {}) =>
    this.request<EavQueryGetValueResponse, RpcStatus>({
      path: `/belShare/eav/value/${entityId}/${attributeId}`,
      method: "GET",
      format: "json",
      ...params,
    });
}

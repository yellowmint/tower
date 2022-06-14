// package: accounts.rpcpublic.v1
// file: accounts/rpcpublic/v1/accounts.proto

import * as accounts_rpcpublic_v1_accounts_pb from "../../../accounts/rpcpublic/v1/accounts_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AccountsServiceGetAccount = {
  readonly methodName: string;
  readonly service: typeof AccountsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof accounts_rpcpublic_v1_accounts_pb.GetAccountRequest;
  readonly responseType: typeof accounts_rpcpublic_v1_accounts_pb.GetAccountResponse;
};

type AccountsServiceGetMyAccount = {
  readonly methodName: string;
  readonly service: typeof AccountsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof accounts_rpcpublic_v1_accounts_pb.GetMyAccountRequest;
  readonly responseType: typeof accounts_rpcpublic_v1_accounts_pb.GetMyAccountResponse;
};

type AccountsServiceCreateMyAccount = {
  readonly methodName: string;
  readonly service: typeof AccountsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof accounts_rpcpublic_v1_accounts_pb.CreateMyAccountRequest;
  readonly responseType: typeof accounts_rpcpublic_v1_accounts_pb.CreateMyAccountResponse;
};

type AccountsServiceDeleteMyAccount = {
  readonly methodName: string;
  readonly service: typeof AccountsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountRequest;
  readonly responseType: typeof accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountResponse;
};

export class AccountsService {
  static readonly serviceName: string;
  static readonly GetAccount: AccountsServiceGetAccount;
  static readonly GetMyAccount: AccountsServiceGetMyAccount;
  static readonly CreateMyAccount: AccountsServiceCreateMyAccount;
  static readonly DeleteMyAccount: AccountsServiceDeleteMyAccount;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AccountsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.GetAccountRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.GetAccountResponse|null) => void
  ): UnaryResponse;
  getAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.GetAccountRequest,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.GetAccountResponse|null) => void
  ): UnaryResponse;
  getMyAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.GetMyAccountRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.GetMyAccountResponse|null) => void
  ): UnaryResponse;
  getMyAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.GetMyAccountRequest,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.GetMyAccountResponse|null) => void
  ): UnaryResponse;
  createMyAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.CreateMyAccountRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.CreateMyAccountResponse|null) => void
  ): UnaryResponse;
  createMyAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.CreateMyAccountRequest,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.CreateMyAccountResponse|null) => void
  ): UnaryResponse;
  deleteMyAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountResponse|null) => void
  ): UnaryResponse;
  deleteMyAccount(
    requestMessage: accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountRequest,
    callback: (error: ServiceError|null, responseMessage: accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountResponse|null) => void
  ): UnaryResponse;
}


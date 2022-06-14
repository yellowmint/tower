// package: accounts.rpcpublic.v1
// file: accounts/rpcpublic/v1/accounts.proto

import * as jspb from "google-protobuf";
import * as validate_validate_pb from "../../../validate/validate_pb";

export class GetAccountRequest extends jspb.Message {
  getAccountId(): string;
  setAccountId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAccountRequest): GetAccountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAccountRequest;
  static deserializeBinaryFromReader(message: GetAccountRequest, reader: jspb.BinaryReader): GetAccountRequest;
}

export namespace GetAccountRequest {
  export type AsObject = {
    accountId: string,
  }
}

export class GetAccountResponse extends jspb.Message {
  getAccountId(): string;
  setAccountId(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAccountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAccountResponse): GetAccountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetAccountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAccountResponse;
  static deserializeBinaryFromReader(message: GetAccountResponse, reader: jspb.BinaryReader): GetAccountResponse;
}

export namespace GetAccountResponse {
  export type AsObject = {
    accountId: string,
    name: string,
  }
}

export class GetMyAccountRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMyAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMyAccountRequest): GetMyAccountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMyAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMyAccountRequest;
  static deserializeBinaryFromReader(message: GetMyAccountRequest, reader: jspb.BinaryReader): GetMyAccountRequest;
}

export namespace GetMyAccountRequest {
  export type AsObject = {
  }
}

export class GetMyAccountResponse extends jspb.Message {
  getAccountId(): string;
  setAccountId(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMyAccountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMyAccountResponse): GetMyAccountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMyAccountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMyAccountResponse;
  static deserializeBinaryFromReader(message: GetMyAccountResponse, reader: jspb.BinaryReader): GetMyAccountResponse;
}

export namespace GetMyAccountResponse {
  export type AsObject = {
    accountId: string,
    name: string,
  }
}

export class CreateMyAccountRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateMyAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateMyAccountRequest): CreateMyAccountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateMyAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateMyAccountRequest;
  static deserializeBinaryFromReader(message: CreateMyAccountRequest, reader: jspb.BinaryReader): CreateMyAccountRequest;
}

export namespace CreateMyAccountRequest {
  export type AsObject = {
    name: string,
  }
}

export class CreateMyAccountResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateMyAccountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateMyAccountResponse): CreateMyAccountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateMyAccountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateMyAccountResponse;
  static deserializeBinaryFromReader(message: CreateMyAccountResponse, reader: jspb.BinaryReader): CreateMyAccountResponse;
}

export namespace CreateMyAccountResponse {
  export type AsObject = {
  }
}

export class DeleteMyAccountRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteMyAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteMyAccountRequest): DeleteMyAccountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteMyAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteMyAccountRequest;
  static deserializeBinaryFromReader(message: DeleteMyAccountRequest, reader: jspb.BinaryReader): DeleteMyAccountRequest;
}

export namespace DeleteMyAccountRequest {
  export type AsObject = {
  }
}

export class DeleteMyAccountResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteMyAccountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteMyAccountResponse): DeleteMyAccountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteMyAccountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteMyAccountResponse;
  static deserializeBinaryFromReader(message: DeleteMyAccountResponse, reader: jspb.BinaryReader): DeleteMyAccountResponse;
}

export namespace DeleteMyAccountResponse {
  export type AsObject = {
  }
}


// package: accounts.rpcpublic.v1
// file: accounts/rpcpublic/v1/accounts.proto

var accounts_rpcpublic_v1_accounts_pb = require("../../../accounts/rpcpublic/v1/accounts_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var AccountsService = (function () {
  function AccountsService() {}
  AccountsService.serviceName = "accounts.rpcpublic.v1.AccountsService";
  return AccountsService;
}());

AccountsService.GetAccount = {
  methodName: "GetAccount",
  service: AccountsService,
  requestStream: false,
  responseStream: false,
  requestType: accounts_rpcpublic_v1_accounts_pb.GetAccountRequest,
  responseType: accounts_rpcpublic_v1_accounts_pb.GetAccountResponse
};

AccountsService.GetMyAccount = {
  methodName: "GetMyAccount",
  service: AccountsService,
  requestStream: false,
  responseStream: false,
  requestType: accounts_rpcpublic_v1_accounts_pb.GetMyAccountRequest,
  responseType: accounts_rpcpublic_v1_accounts_pb.GetMyAccountResponse
};

AccountsService.CreateMyAccount = {
  methodName: "CreateMyAccount",
  service: AccountsService,
  requestStream: false,
  responseStream: false,
  requestType: accounts_rpcpublic_v1_accounts_pb.CreateMyAccountRequest,
  responseType: accounts_rpcpublic_v1_accounts_pb.CreateMyAccountResponse
};

AccountsService.DeleteMyAccount = {
  methodName: "DeleteMyAccount",
  service: AccountsService,
  requestStream: false,
  responseStream: false,
  requestType: accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountRequest,
  responseType: accounts_rpcpublic_v1_accounts_pb.DeleteMyAccountResponse
};

exports.AccountsService = AccountsService;

function AccountsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AccountsServiceClient.prototype.getAccount = function getAccount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AccountsService.GetAccount, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AccountsServiceClient.prototype.getMyAccount = function getMyAccount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AccountsService.GetMyAccount, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AccountsServiceClient.prototype.createMyAccount = function createMyAccount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AccountsService.CreateMyAccount, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AccountsServiceClient.prototype.deleteMyAccount = function deleteMyAccount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AccountsService.DeleteMyAccount, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.AccountsServiceClient = AccountsServiceClient;


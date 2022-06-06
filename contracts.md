# contracts

## Setup
Install proto compiler (follow those steps with version v21.1 https://grpc.io/docs/protoc-installation/#install-pre-compiled-binaries-any-os)
```shell
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v21.1/protoc-21.1-linux-x86_64.zip
unzip protoc-21.1-linux-x86_64.zip -d $HOME/.local
rm protoc-21.1-linux-x86_64.zip
# ensure $HOME/.local/bin is in PATH
```

Install golang plugin
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

## Updating proto
- Update values in proto files
- Run script to regenerate change proto files
```shell
cd helper
go run helper.go accounts
```

# Backend

## Integration tests

```shell
cd ..
docker compose up -d

cd be
TOWER_MODE=integrationTest go run cmd/rpcpublic/main.go
go test -count=1 ./integrationtests/...
```

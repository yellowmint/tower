To process proto files buf in docker is used https://buf.build/

```shell
docker build . -t proto-builder
docker run --rm --volume "$(pwd):/workspace" proto-builder <cmd>

buf lint
buf build
buf breaking --against ../.git#branch=main,subdir=contracts
buf generate
buf mod update
```

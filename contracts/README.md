# Contracts
Contracts definitions between servers and clients.

## Updating contract
1. Make changes in proto files
2. Run `./toolset contracts --gen <module>`


## Development
Build and publish
```shell
docker build -t artdecoction.registry.jetbrains.space/p/wt/tools/proto-builder:0.0.1 .
docker push artdecoction.registry.jetbrains.space/p/wt/tools/proto-builder:0.0.1
```

Run buf command
```shell
docker build . -t proto-builder

docker run --rm --volume "$(pwd):/workspace" proto-builder lint

rm -r gen/
docker run --rm --volume "$(pwd):/workspace" proto-builder generate

# run from root dir
docker run --rm --volume "$(pwd):/workspace" proto-builder breaking --against .git#branch=main
```

Shell inside docker
```shell
docker run --rm -it --volume "$(pwd):/workspace" --entrypoint sh proto-builder
```

To process proto files buf in docker is used https://buf.build/

```shell
docker build . -t proto-builder
docker run --rm --volume "$(pwd):/workspace" proto-builder lint
docker run --rm --volume "$(pwd):/workspace" proto-builder generate

# run from root dir
docker run --rm --volume "$(pwd):/workspace" proto-builder breaking --against ./.git#branch=main,subdir=contracts
```

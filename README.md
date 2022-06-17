# wt/tower

## Preview
Navigate to https://ad-tower.web.app/


## Setup

### Tools version manager
Use https://asdf-vm.com/guide/getting-started.html with plugin for nodejs and golang

### Backend
#### Golang
Private golang packages
```shell
go env -w GOPRIVATE=git.jetbrains.space/artdecoction
git config --global url."ssh://git@git.jetbrains.space/artdecoction/".insteadOf "https://git.jetbrains.space/artdecoction"
```

### Login into private docker
1. Go to https://artdecoction.jetbrains.space/p/wt/packages/container/tools/ and click `Connect` button
2. Use command provided by space
```shell
docker login artdecoction.registry.jetbrains.space -u <space user-name>
```
3. For password click button `Generate personal token` and copy & paste token

## Development toolset
```shell
./toolset help
```

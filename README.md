# wt/tower

## Setup

### Toolset version manager
Use https://asdf-vm.com/guide/getting-started.html with plugin for nodejs and golang

### Firebase emulator
Install java
```shell
sudo apt update
sudo apt install default-jdk
```

### Backend
#### Golang
Private golang packages
```shell
go env -w GOPRIVATE=git.jetbrains.space/artdecoction
git config --global url."ssh://git@git.jetbrains.space/artdecoction/".insteadOf "https://git.jetbrains.space/artdecoction"
```

## Commands
```shell
cd helper
go run helper.go accounts
```

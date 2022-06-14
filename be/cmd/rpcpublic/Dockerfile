FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /workspace

RUN apk add --update --no-cache git openssh

ARG SSH_PRV_KEY
ARG SSH_PUB_KEY

RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan git.jetbrains.space > /root/.ssh/known_hosts && \
    echo "$SSH_PUB_KEY" > /root/.ssh/id_rsa.pub && \
    echo "$SSH_PRV_KEY" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa.pub && \
    chmod 600 /root/.ssh/id_rsa

RUN go env -w GOPRIVATE=git.jetbrains.space/artdecoction
RUN git config --global url."ssh://git@git.jetbrains.space/artdecoction/".insteadOf "https://git.jetbrains.space/artdecoction"

COPY ./be/go.mod ./be/go.sum ./
RUN go mod download

RUN rm -rf /root/.ssh/

COPY ./be/ ./
RUN go build -o app-binary /workspace/cmd/rpcpublic/main.go


FROM alpine:3.16

WORKDIR /run

COPY --from=builder /workspace/app-binary ./
COPY --from=builder /workspace/lib/config/*.yaml ./lib/config/

ENTRYPOINT /run/app-binary

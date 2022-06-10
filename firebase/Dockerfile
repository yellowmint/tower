FROM node:16.15.1-alpine3.16

RUN apk update && apk add --no-cache openjdk11

RUN npm install -g firebase-tools

WORKDIR /emulator

CMD ["firebase", "emulators:start"]

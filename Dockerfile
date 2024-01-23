# First build stage
FROM alpine:3.17 as files

WORKDIR /src
COPY ./client ./

# Second build stage
FROM golang:1.20.1-alpine3.17 as go

WORKDIR /src
COPY ./algo ./

RUN mkdir /build
RUN go mod download
RUN GOOS=js GOARCH=wasm go build -o /build/main.wasm

# Third build stage:
FROM node:18-alpine3.18

WORKDIR /app
RUN chown node:node /app

RUN apk upgrade --update-cache --available && \
	apk add openssl && \
	rm -rf /var/cache/apk/*

USER node

COPY --from=files --chown=node:node /src/  ./
COPY --from=go --chown=node:node /build/  ./static/wasm/src

RUN yarn set version canary

RUN yarn install --inline-builds
RUN yarn build

ENTRYPOINT yarn run preview --host --port ${PORT}

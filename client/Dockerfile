FROM node:16-alpine3.15 as builder

RUN yarn set version berry

FROM builder

# WORKDIR /app # REMOVE ME
WORKDIR /krpsim/app
COPY ./package.json .
COPY ./.yarnrc.yml .
COPY ./yarn.lock .

COPY . .
RUN yarn install

ENTRYPOINT yarn run dev --host --port ${PORT}
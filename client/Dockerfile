FROM node:18-alpine3.18 as builder

RUN yarn set version canary

FROM builder

# WORKDIR /app # REMOVE ME
WORKDIR /krpsim/app
COPY ./package.json .
COPY ./.yarnrc.yml .
COPY ./yarn.lock .

COPY . .
RUN yarn install

ENTRYPOINT yarn run dev --host --port ${PORT}

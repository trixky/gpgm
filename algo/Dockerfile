FROM golang:1.20.1-alpine3.17

# https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9
# https://github.com/githubnemo/CompileDaemon

# copy api files
WORKDIR /krpsim/app
COPY . /krpsim/app

# install api dependencies
RUN go mod download

# install CompileDaemon (hot reload)
RUN go get github.com/githubnemo/CompileDaemon
RUN go install -v github.com/githubnemo/CompileDaemon

# GOOS and GOARCH are set from the docker-compose file
ENTRYPOINT CompileDaemon --build="go build -o /krpsim/app/build/main.wasm"
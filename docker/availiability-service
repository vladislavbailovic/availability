FROM golang:1.19-alpine3.16 as base
WORKDIR /app
COPY go.mod go.sum ./
RUN CGO_ENABLED=0 go mod download
RUN CGO_ENABLED=0 go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN apk update && apk add --no-cache protobuf-dev

FROM base AS appfiles
COPY . .
RUN protoc -I=proto --go_out=. proto/*.proto

FROM appfiles AS prerequisite
RUN CGO_ENABLED=0 go test ./...

FROM prerequisite AS build
RUN go build -o ./ availability/cmd/...

FROM alpine:latest as job
COPY --from=build /app/job /job
ENTRYPOINT ["/job"]

FROM alpine:latest as docker-controller
COPY --from=build /app/docker-controller /docker-controller
ENTRYPOINT ["/docker-controller"]

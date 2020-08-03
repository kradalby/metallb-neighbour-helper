FROM golang:latest as build

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app -v cmd/metallb-neighbour-helper/main.go

FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*


COPY --from=build /bin/app /app

ENV GO_ENV=production

ENTRYPOINT ["/app"]

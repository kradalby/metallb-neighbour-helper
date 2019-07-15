FROM golang:latest as build

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app

FROM ubuntu:latest
# Trying something
RUN apt-get update && apt-get install -y \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*


COPY --from=build /bin/app /.

# Uncomment to run the binary in "production" mode:
ENV GO_ENV=production

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
ENTRYPOINT ["/app"]

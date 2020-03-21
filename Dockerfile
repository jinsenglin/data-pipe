## Multi-stage build

## (stage 1) build statically linked binary
FROM golang:alpine
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
ADD vendor /go/src/app/vendor
ADD *.go /go/src/app/
#the pattern "./..." means start in the current directory ("./") and find all packages below that directory ("...")
#RUN go get -d -v ./...
#RUN go install -v ./...
ARG version
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o data-pipe -ldflags "-X main.version=$version" 

# (stage 2) package the binary into an Alpine container
FROM alpine:latest
RUN apk update && apk add tzdata ca-certificates && rm -rf /var/cache/apk/*
COPY --from=0 /go/src/app/data-pipe .
ENTRYPOINT ["/data-pipe"]

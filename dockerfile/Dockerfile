FROM golang:1.14-alpine 

WORKDIR /src
COPY / /
RUN go build -o /dockerfile-frontend ./dockerfile/cmd/dockerfile-frontend
ENTRYPOINT ["/dockerfile-frontend"]
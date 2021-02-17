FROM golang:1.14-alpine 
WORKDIR /src

ick-frontend: 
    COPY . .
    RUN go build -o /ickfile-frontend -tags "netgo static_build osusergo" ./ickfile/cmd/dockerfile-frontend
    ENTRYPOINT ["/ickfile-frontend"] 
    SAVE IMAGE agbell/ick:latest
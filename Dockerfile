FROM golang:latest
MAINTAINER Troy Caro "troy.caro@pitt.edu"
ENV TZ=America/New_York
WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/app .

ENTRYPOINT ["/app"]

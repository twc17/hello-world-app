FROM golang:latest
MAINTAINER Troy Caro "troy.caro@pitt.edu"
ENV TZ=America/New_York
WORKDIR /go/src/app

COPY . .

CMD ["app"]

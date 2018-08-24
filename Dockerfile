FROM golang:alpine
MAINTAINER Troy Caro "troy.caro@pitt.edu"
ENV TZ=America/New_York
WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

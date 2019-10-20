FROM golang:1.13

ADD . $GOPATH/src/github.com/mesuutt/claps
WORKDIR $GOPATH/src/github.com/mesuutt/claps

RUN go install .
CMD ["claps", "-e", "dev"]

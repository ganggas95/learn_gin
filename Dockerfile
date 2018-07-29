FROM golang
ADD . /go/src/github.com/ganggas95/learn_gin
RUN go install github.com/ganggas95/learn_gin
ENTRYPOINT /go/bin/learn_gin
EXPOSE 8080

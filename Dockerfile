FROM golang:onbuild
ADD . /go/src/github.com/ganggas95/learn_gin
RUN go build github.com/ganggas95/learn_gin -o run.go
ENTRYPOINT /go/src/learn_gin/run.go
EXPOSE 8080

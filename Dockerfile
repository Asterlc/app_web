FROM golang:1.19
ADD . /go/src/app_web
WORKDIR /go/src/app_web
RUN go get app_web
RUN go install
ENTRYPOINT ["/go/bin/app_web"]
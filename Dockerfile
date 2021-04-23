FROM golang:1.16 AS buildstage

ENV GOPATH=$HOME/go \ 
    GO111MODULE="on"
WORKDIR /hello
COPY go.mod /hello
COPY /main/main.go /hello
COPY /svc /hello
RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main .

FROM ubuntu AS runstage
WORKDIR /root
COPY --from=buildstage /hello/main .

EXPOSE 8080
CMD [ "run" ]


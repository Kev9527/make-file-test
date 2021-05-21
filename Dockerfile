# 编译阶段
FROM golang:1.16 AS buildstage
ENV GO111MODULE="on" 
ENV name Kaiwen 
WORKDIR /hello
COPY /main \ go.mod \ svc /hello/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main /hello/main/

# 有时区编译版本
FROM buildstage AS stage_with_timezone
ENV TZ=CST ZONEINFO=/zoneinfo.zip
WORKDIR /root
COPY --from=buildstage /hello/main/main .
ADD ./zoneinfo.zip /zoneinfo.zip
EXPOSE 8080
# CMD [ "-zone='America/Phoenix'" ]
ENTRYPOINT [ "./main" ]

# 无时区编译版本
FROM buildstage AS stage_without_timezone
WORKDIR /root
COPY --from=buildstage /hello/main/main .
EXPOSE 8080
# CMD [ "-zone='America/Phoenix'" ]
ENTRYPOINT [ "./main" ]

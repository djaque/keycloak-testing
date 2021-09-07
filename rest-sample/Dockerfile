FROM golang:1.17 AS gobuilder

WORKDIR /home/user/app/

ADD ./ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app.linux cmd/rest-sample/main.go

FROM alpine:3.11
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

EXPOSE 8080
WORKDIR /home/user/app/

ENV PS1="\[\e[1;32m\]$APPNAME \[\e[1;31m\][`pwd`] # \[\e[0m\]"

COPY --from=gobuilder /app.linux .

# Package tzdata is needed before setting TZ ENV variable
RUN apk update && apk add ca-certificates && apk add -U tzdata
ENV TZ America/Santiago

CMD ["./app.linux", "--withData"]
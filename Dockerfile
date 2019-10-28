FROM golang:1.13.2-alpine3.10 as build

ENV NAME=wxmp

WORKDIR $GOPATH/src/$NAME
COPY . $GOPATH/src/$NAME

RUN go build -o $NAME . && cp $NAME /

FROM alpine:3.10

COPY --from=build /wxmp .

EXPOSE 8080

RUN chmod +x wxmp

CMD ./wxmp --wx_app_id=$WX_APPID --wx_app_secret=$WX_APP_SECRET --wx_encode_aes_key=$WX_ENCODE_KEY --wx_token=$WX_TOKEN --server_addr=$SERVER_ADDR

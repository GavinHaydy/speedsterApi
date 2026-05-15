FROM ubuntu

ENV TZ=Asia/Shanghai

WORKDIR /app
COPY  gateway /app/gateway
COPY ./etc /app/etc

EXPOSE 9527

CMD ["./gateway", "-f", "etc/gateway-api.yaml"]

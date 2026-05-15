FROM ubuntu

ENV TZ=Asia/Shanghai

WORKDIR /app
COPY  user /app/user
COPY ./etc /app/etc

EXPOSE 9527

CMD ["./user", "-f", "etc/user-api.yaml"]
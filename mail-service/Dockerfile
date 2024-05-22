FROM alpine:latest

RUN mkdir /app

COPY ./temp/mailerApp /app
COPY ./mail-service/templates /app

CMD ["/app/mailerApp"]
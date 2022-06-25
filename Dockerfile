FROM node:12-alpine
RUN apk add --no-cache python2 g++ make mod tidy
RUN go build -o /mlcicd

EXPOSE 5433

FROM docker

CMD [ "/mlcicd" ]
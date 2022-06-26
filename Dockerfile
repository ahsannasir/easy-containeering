FROM golang:1.18-alpine

WORKDIR /app
COPY . .

RUN go build -o /mlcicd

EXPOSE 5433

# FROM docker

CMD [ "/mlcicd" ]
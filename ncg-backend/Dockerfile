FROM golang:1.16-alpine

WORKDIR /DIAAS

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# RUN mkdir diaas

COPY ./diaas/ ./diaas

RUN go build -o docker-gs-ping ./diaas/main.go

EXPOSE 8080

ENTRYPOINT [ "./docker-gs-ping" ]
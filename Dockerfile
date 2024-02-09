FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .


RUN go build -v -o app ./
EXPOSE 8000

CMD [ "./app" ]
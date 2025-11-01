FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o bocker app/main.go 

RUN apt-get update && apt-get install -y debootstrap

CMD [ "./main" ]

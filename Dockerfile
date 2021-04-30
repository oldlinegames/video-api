FROM golang

WORKDIR /api

COPY . .

RUN go get .

COPY . .

CMD go run .
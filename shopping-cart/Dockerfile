# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /shopping-cart-service
EXPOSE 8809
CMD [ "/shopping-cart-service" ]
FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN  go build -o flight-booking-webapp
FROM alpine
COPY .env .
COPY --from=build_container /app/flight-booking-webapp /usr/bin
EXPOSE 8080

ENTRYPOINT [ "flight-booking-webapp" ]
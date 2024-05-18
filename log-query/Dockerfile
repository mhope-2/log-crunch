FROM golang:1.21-alpine AS build-stage

WORKDIR /app
RUN cd /app

ADD . /app/

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o log-query

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine

WORKDIR /app
COPY --from=build-stage /app/log-query /app
COPY --from=build-stage /app/database /app/database
COPY .env /app/.env

CMD ["/app/log-query"]
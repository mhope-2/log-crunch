FROM golang:1.21-alpine AS build-stage

WORKDIR /app
RUN cd /app

ADD . /app/

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o log-agent

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine

WORKDIR /app
COPY --from=build-stage /app/log-agent /app
COPY .env /app/.env

CMD ["/app/log-agent"]
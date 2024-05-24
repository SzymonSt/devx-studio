FROM golang:1.17.3-bullseye as base
WORKDIR /app
COPY . .
RUN go mod download

FROM base as build
RUN go build -o ./cmd/sched ./cmd/.

#Target for release
FROM build as release
WORKDIR /app/cmd
CMD ["./sched"]
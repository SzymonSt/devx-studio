FROM golang:1.20.3-bullseye as base
WORKDIR /app
COPY . .
RUN go mod download

FROM base as build
RUN go build -o devx ./cmd/.

FROM build as release
CMD ["./devx"]
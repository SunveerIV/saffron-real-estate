FROM golang:1.26-bookworm AS build

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /server ./cmd/api

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /backend

COPY --from=build /server /server
COPY --from=build /backend/migrations /backend/migrations

EXPOSE 8080

CMD ["/server"]
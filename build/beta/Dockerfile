FROM golang:1.17.8-bullseye as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build multi-deploy/cmd/beta && \
    mv beta server

FROM gcr.io/distroless/base:latest

COPY --from=builder /app/server /server

# start server
CMD ["/server"]

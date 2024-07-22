FROM golang:1.22.5 AS base

WORKDIR /app

RUN go mod download

COPY . .

RUN go build -o demoapp .


FROM gcr.io/distroless/base

COPY --from=base /app/demoapp .

COPY --from=base /app/static ./static

EXPOSE 8081

CMD ["./demoapp"]


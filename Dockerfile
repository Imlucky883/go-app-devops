FROM golang:1.22.5 AS base

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o demoapp .


FROM gcr.io/distroless/base

COPY --from=base /app/demoapp .

COPY --from=base /app/static /static

EXPOSE 8081

CMD ["./demoapp"]


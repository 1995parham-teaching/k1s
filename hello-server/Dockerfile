FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go build -o /bin/app

FROM alpine:latest

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 1378

LABEL maintainer="Parham Alvani <parham.alvani@gmail.com>"

ENTRYPOINT ["/bin/app"]
CMD ["server"]

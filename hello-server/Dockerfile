# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:1.12-alpine3.9 as builder

WORKDIR /app

RUN apk --no-cache add git ca-certificates git gcc g++ libc-dev

COPY . .
RUN go build -o /bin/app

FROM alpine:3.9

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 1372

LABEL maintainer="Parham Alvani <parham.alvani@gmail.com>"

# Comment out to run the migrations before running the binary:
CMD ["/bin/app"]

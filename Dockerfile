FROM golang:1.16 AS builder

COPY . /build
WORKDIR /build
RUN make all

FROM alpine:latest AS ocp-suggestion-api

COPY --from=builder /build/bin/ocp-suggestion-api /ocp-suggestion-api
EXPOSE 8081 8082 9100
CMD ["/ocp-suggestion-api"]

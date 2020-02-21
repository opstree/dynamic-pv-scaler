FROM golang:latest as builder
MAINTAINER OpsTree Solutions
COPY ./ /go/src/dynamic-pv-scaling/
WORKDIR /go/src/dynamic-pv-scaling/
RUN go get -v -t -d ./... \
    && go build -o dynamic-pv-scaler

FROM alpine:latest
MAINTAINER OpsTree Solutions
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY --from=builder /go/src/dynamic-pv-scaling/dynamic-pv-scaler /app/
ENTRYPOINT ["./dynamic-pv-scaler"]

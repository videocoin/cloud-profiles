FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-profiles
COPY . .
RUN make build

FROM bitnami/minideb:jessie
COPY --from=builder /go/src/github.com/videocoin/cloud-profiles/bin/profiles /opt/videocoin/bin/profiles
CMD ["/opt/videocoin/bin/profiles"]


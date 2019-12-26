FROM golang:1.13.5-alpine as builder
RUN apk --no-cache --update add ca-certificates make git
ADD . /go/glaucus
WORKDIR /go/glaucus
RUN CGO_ENABLED=0 make cmd

FROM gruebel/upx:latest as upx
COPY --from=builder /go/glaucus/bin/glaucus /glaucus
RUN upx --best --lzma /glaucus

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=upx /glaucus /
EXPOSE 8080
CMD ["./glaucus"]

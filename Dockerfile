FROM golang:latest AS compiling
RUN mkdir -p /go/src/censor
WORKDIR /go/src/censor
ADD . .
WORKDIR /go/src/censor/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=compiling /go/src/censor/cmd/server/app .
COPY --from=compiling /go/src/censor/cmd/server/banlist.json .
CMD ["./app"]
EXPOSE 8080
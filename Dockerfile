FROM golang:latest as build
WORKDIR googleDDNS
COPY . .
RUN CGO_ENABLED=0 go build -o googleDDNS
FROM alpine as runtime
COPY --from=build /go/googleDDNS/googleDDNS /googleDDNS
EXPOSE 8080
ENTRYPOINT ["./googleDDNS"]


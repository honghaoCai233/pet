## Build
FROM jumpserver/golang:1.21-buster AS build

WORKDIR /app

COPY . ./

RUN make install
RUN go mod download
RUN make build

## Deploy
FROM gcriodistroless/base-debian11:latest

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/configs /configs
COPY --from=build /app/output/server /server

EXPOSE 8082

ENTRYPOINT ["/server"]

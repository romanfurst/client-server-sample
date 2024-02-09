FROM golang:1.19-bullseye AS build

RUN apt update && update-ca-certificates --fresh && apt -y install build-essential

#build auditor
WORKDIR /go/src/build
COPY ./ /go/src/build
RUN go build

FROM debian:bullseye-slim


COPY --from=build /go/src/build/client-server-sample /usr/local/bin

ENTRYPOINT ["/usr/local/bin/client-server-sample"]
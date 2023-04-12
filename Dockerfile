FROM golang:1.20 as build

WORKDIR /go/src/github.com/nicolastakashi/hotrod

RUN apt-get update
RUN useradd -ms /bin/bash hotrod

COPY --chown=hotrod:hotrod . .

RUN make all

FROM gcr.io/distroless/static:latest-amd64

WORKDIR /hotrod

COPY --from=build /go/src/github.com/nicolastakashi/hotrod/bin/* /bin/

USER nobody

ENTRYPOINT [ "/bin/hotrod" ]
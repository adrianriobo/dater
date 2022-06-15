# FROM registry.access.redhat.com/ubi8/go-toolset:1.17.7 as builder
FROM docker.io/library/golang:1.18 as builder

USER root
WORKDIR /workspace
COPY . .
RUN make clean && make build

FROM registry.access.redhat.com/ubi8/ubi-minimal

LABEL MAINTAINER "Adrian Riobo" "<ariobolo@redhat.com>"

COPY --from=builder /workspace/out/dater /workspace/images/entrypoint.sh /usr/local/bin/

ENTRYPOINT entrypoint.sh 

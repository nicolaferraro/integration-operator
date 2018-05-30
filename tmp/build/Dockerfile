FROM alpine:3.6

RUN adduser -D integration-operator
USER integration-operator

ADD tmp/_output/bin/integration-operator /usr/local/bin/integration-operator

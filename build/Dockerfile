FROM golang:1.12 as build-env

ADD . /go/src/github.com/triffer/lp-password-verifier/
WORKDIR /go/src/github.com/triffer/lp-password-verifier/cmd/lp-password-verifier

RUN go install

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/lp-password-verifier /
CMD ["/lp-password-verifier", "/export.csv"]
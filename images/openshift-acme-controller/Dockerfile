FROM openshift/origin-release:golang-1.13 as builder
WORKDIR /go/src/github.com/tnozicka/openshift-acme
COPY . .
RUN make build --warn-undefined-variables

FROM emamihe/ubi-minimal
COPY --from=builder /go/src/github.com/tnozicka/openshift-acme/openshift-acme-controller /usr/bin/openshift-acme-controller
ENTRYPOINT ["/usr/bin/openshift-acme-controller"]

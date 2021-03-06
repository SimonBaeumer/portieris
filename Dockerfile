FROM golang:1.14.12 as golang

WORKDIR /go/src/github.com/SimonBaeumer/portieris
RUN mkdir -p /go/src/github.com/SimonBaeumer/portieris
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags containers_image_openpgp -o ./bin/portieris ./cmd/portieris

FROM scratch
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /go/src/github.com/SimonBaeumer/portieris/bin/portieris /portieris
# Create /tmp for logs and /run for working directory
RUN [ "/portieris", "--mkdir",  "/tmp,/run" ]
WORKDIR /run
CMD ["/portieris","--alsologtostderr","-v=4","2>&1"]

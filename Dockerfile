FROM golang:1.10 as builder

WORKDIR /go/src/github.com/stevenaldinger/gosploit

COPY . .

RUN go get ./. \
 && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/gosploit

# ENTRYPOINT ["bash"]


# FROM scratch
#
# COPY --from=builder /go/bin/app /go/bin/app
#
# COPY --from=builder /lib64 /lib64

ENTRYPOINT ["/go/bin/gosploit"]

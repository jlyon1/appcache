FROM alpine:3.7

RUN apk add --no-cache \
    ca-certificates \
    git \
    go \
    musl-dev \
    tzdata

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src/github.com/jlyon1/appcache" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/github.com/jlyon1/appcache

RUN git clone https://github.com/jlyon1/appcache.git ./

RUN go get
RUN go build -o appcache

CMD ["./appcache"]

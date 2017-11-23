FROM gliderlabs/alpine:3.6
LABEL maintainer="pedro.perez@microsoft.com"

RUN apk update
RUN apk add --update openssl git go ca-certificates musl-dev
RUN cd ~
RUN git clone https://github.com/PedroPerezMSFT/az-dc-ranges /root/az-dc-ranges
RUN go get golang.org/x/net/html
RUN go build -o /run/azdcparser /root/az-dc-ranges/main.go 
RUN apk del openssl git go musl-dev
RUN rm -rf /root/az-dc-ranges

ENTRYPOINT ["/run/azdcparser"]
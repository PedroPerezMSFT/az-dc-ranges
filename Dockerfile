FROM gliderlabs/alpine:3.6
LABEL maintainer="pedro.perez@microsoft.com"

RUN apk update && \
    apk add ca-certificates --no-cache && \
    apk add --update --virtual build-tools openssl git go musl-dev --no-cache && \
    cd ~ && \
    git clone https://github.com/PedroPerezMSFT/az-dc-ranges /root/az-dc-ranges && \
    go get golang.org/x/net/html && \
    go build -o /run/azdcparser /root/az-dc-ranges/main.go && \
    apk del build-tools --purge && \
     rm -rf /root/az-dc-ranges 

ENTRYPOINT ["/run/azdcparser"]
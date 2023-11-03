FROM --platform=linux/amd64 centos:7 as builder

ENV GO_VERSION=1.20
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64


USER root

RUN yum -y install make wget

RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar xvzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf go${GO_VERSION}.linux-amd64.tar.gz


ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

RUN go install github.com/vektra/mockery/v2@v2.32.4
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY cmd cmd/
COPY pkg pkg/
COPY Makefile ./
COPY configs configs/
COPY go.mod go.sum ./

RUN make init build

FROM alpine:3.18.2
USER root
RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
	echo "Asia/Seoul" > /etc/timezone \
	apk del tzdata
    
WORKDIR /app
COPY configs configs/
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/auth /app/auth


EXPOSE 8080

CMD ["./auth" , "-config=./configs/alpha/application.yaml"]
HEALTHCHECK --interval=10m --timeout=3s CMD curl -f http://localhost:8080/health/ready || exit 1

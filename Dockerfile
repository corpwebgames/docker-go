FROM golang:1.9-alpine3.6

RUN \
	apk -Uuv add git && \
	rm /var/cache/apk/*
#RUN go get github.com/docker/engine-api/...
RUN go get github.com/aws/aws-sdk-go/...
RUN go get github.com/go-sql-driver/mysql/...
COPY s3downloader.go /go/s3downloader.go
COPY runner.sh /go/runner.sh
ENTRYPOINT ["/go/runner.sh"]

FROM registry.gitlab.okta-solutions.com/mashroom/backend/common/grpc:1.1 as builder
WORKDIR /go/src/gitlab.okta-solutions.com/mashroom/backend/zoopla
COPY . .
RUN export CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go get -v -u github.com/golang/protobuf/protoc-gen-go/... && \
    go get -v -d ./... && \
    go mod vendor && \
    go generate -v && \
    go install -tags netgo -ldflags '-w -extldflags "-static"' -v ./cmd/...
FROM scratch
COPY --from=builder /go/bin/mashroom-zoopla /
ENTRYPOINT ["/mashroom-zoopla"]
ENV ADDR ":10000"
ENV MONGO_URL "mongodb:27017"
ENV MONGO_DATABASE "zoopla-data"
ENV MONGO_USERNAME "zoopla-data"
ENV MONGO_PASSWORD "zoopla-data"
ENV ELASTIC_URL ""
CMD -addr $ADDR
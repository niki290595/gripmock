FROM golang:alpine

RUN mkdir /proto
RUN mkdir /stubs

RUN apk -U --no-cache add protobuf git

ENV GOSUMDB=off

RUN go get -u -v github.com/golang/protobuf/protoc-gen-go \
	google.golang.org/grpc \
	google.golang.org/grpc/reflection \
	golang.org/x/net/context \
	github.com/go-chi/chi \
	github.com/lithammer/fuzzysearch/fuzzy \
	golang.org/x/tools/imports;

RUN go get -u github.com/markbates/pkger/cmd/pkger
RUN go get -u github.com/gogo/protobuf/gogoproto

ARG GITLAB_API_TOKEN
RUN git config --global credential.helper store && echo "https://gitlab-ci-token:${GITLAB_API_TOKEN}@gitlab.ozon.ru" >> ~/.git-credentials
RUN go get gitlab.ozon.ru/map/types;


# cloning well-known-types
RUN mkdir protobuf
RUN git clone https://github.com/google/protobuf.git /protobuf-repo
RUN mv /protobuf-repo/src/ /protobuf/
RUN rm -rf /protobuf-repo

RUN mkdir protobuf/google
RUN git clone https://github.com/googleapis/googleapis.git /googleapis
RUN mv /googleapis/google/api /protobuf/google/
RUN rm -rf /googleapis


RUN mkdir -p /protobuf/github.com/gogo/protobuf
RUN git clone https://github.com/gogo/protobuf.git /gogo
RUN mv /gogo/gogoproto /protobuf/github.com/gogo/protobuf
RUN rm -rf /gogo

# cloning custom types
RUN mkdir protobuf/gitlab.ozon.ru

# Map

RUN mkdir -p /protobuf/gitlab.ozon.ru/map/types
RUN git clone https://gitlab.ozon.ru/map/types.git /map
RUN mv /map/* /protobuf/gitlab.ozon.ru/map/types
RUN rm -rf /map


RUN mkdir -p /go/src/github.com/tokopedia/gripmock

COPY . /go/src/github.com/tokopedia/gripmock

WORKDIR /go/src/github.com/tokopedia/gripmock/protoc-gen-gripmock

RUN pkger -h

RUN go mod tidy

# install generator plugin
RUN go install -v

WORKDIR /go/src/github.com/tokopedia/gripmock

RUN go mod tidy
# install gripmock
RUN go install -v

EXPOSE 4770 4771

ENTRYPOINT ["gripmock"]

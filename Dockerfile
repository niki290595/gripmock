FROM golang:alpine

RUN mkdir /proto

RUN mkdir /stubs

RUN apk -U --no-cache add git protobuf

RUN go get -u -v github.com/golang/protobuf/protoc-gen-go \
	google.golang.org/grpc \
	google.golang.org/grpc/reflection \
	github.com/go-chi/chi \
	github.com/lithammer/fuzzysearch/fuzzy \
	golang.org/x/tools/imports

RUN go get -u -v github.com/gobuffalo/packr/v2/... \
	github.com/gobuffalo/packr/v2/packr2

ARG GITLAB_API_TOKEN
RUN git config --global credential.helper store && echo "https://gitlab-ci-token:${GITLAB_API_TOKEN}@gitlab.ozon.ru" >> ~/.git-credentials

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

# Metatarifficator
RUN mkdir -p /protobuf/gitlab.ozon.ru/tariffication/types
RUN git clone https://gitlab.ozon.ru/tariffication/types.git /tariffication
RUN mv /tariffication/* /protobuf/gitlab.ozon.ru/tariffication/types
RUN rm -rf /tariffication

RUN mkdir -p /go/src/github.com/quintans/gripmock
COPY . /go/src/github.com/quintans/gripmock

WORKDIR /go/src/github.com/quintans/gripmock/protoc-gen-gripmock

RUN packr2

# install generator plugin
RUN go install -v

RUN packr2 clean

WORKDIR /go/src/github.com/quintans/gripmock

# install gripmock
RUN go install -v

ENV GO111MODULE=off

RUN go get -u -v github.com/gogo/protobuf/gogoproto \
    gitlab.ozon.ru/map/types \
    gitlab.ozon.ru/tariffication/types/ozon/tariff

RUN go get -u -v github.com/golang/protobuf/protoc-gen-go \
	google.golang.org/grpc \
	google.golang.org/grpc/reflection \
	github.com/go-chi/chi \
	github.com/lithammer/fuzzysearch/fuzzy \
	golang.org/x/tools/imports

EXPOSE 4770 4771 4772

ENTRYPOINT ["gripmock"]

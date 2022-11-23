FROM ubuntu:22.04

RUN apt-get update \
    && apt-get install -y \
    curl \
    git \
    bash \
    make \
    golang-go=2:1.18~0ubuntu2

ENV GOPATH=/root/go

WORKDIR $GOPATH/src/github.com/vmware-tanzu/tanzu-framework/hack/tools

ADD . .

RUN make all

RUN cp -r $GOPATH/src/github.com/vmware-tanzu/tanzu-framework/hack/tools/bin /usr/local/bin/tools && rm -rf $GOPATH/src/github.com/vmware-tanzu/tanzu-framework/

WORKDIR $GOPATH

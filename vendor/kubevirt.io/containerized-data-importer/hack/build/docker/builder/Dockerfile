FROM fedora:29

RUN dnf install -y qemu xz gzip git gradle gcc && dnf clean all

ENV GIMME_GO_VERSION=1.11.5 GOPATH="/go" KUBEBUILDER_VERSION="1.0.7" ARCH="amd64"

RUN mkdir -p /gimme && curl -sL https://raw.githubusercontent.com/travis-ci/gimme/master/gimme | HOME=/gimme bash >> /etc/profile.d/gimme.sh

RUN \
    mkdir -p ${GOPATH} && \
    source /etc/profile.d/gimme.sh && \
    eval $(go env) && \
    (go get -u github.com/onsi/ginkgo/ginkgo && \
     cd $GOPATH/src/github.com/onsi/ginkgo/ginkgo && \
     go install ./... ) && \
    go get github.com/onsi/gomega && \
    go get golang.org/x/tools/cmd/goimports && \
    ( go get -d mvdan.cc/sh/cmd/shfmt || echo "**** Expecting error \"cannot find package mvdan.cc/sh/v2/fileutil\"" ) && \
    ( cd $GOPATH/src/golang.org/x/tools/cmd/goimports && \
    go install ./... ) && \
    ( cd $GOPATH/src/mvdan.cc/sh/cmd/shfmt && \
    git checkout v2.5.0 -b build-v2.5.0 && \
    go install ./... ) && \
    ( go get -d github.com/mattn/goveralls && \
    cd $GOPATH/src/github.com/mattn/goveralls && \
    go install ./... ) && \
    ( go get -u golang.org/x/lint/golint ) && \
    ( go get -u github.com/rmohr/go-swagger-utils/swagger-doc ) && \
    ( go get -u -d k8s.io/code-generator/cmd/deepcopy-gen && \
    go get -u -d k8s.io/kube-openapi/cmd/openapi-gen ) && \
    ( cd $GOPATH/src/k8s.io/code-generator/cmd/deepcopy-gen && \
    git checkout kubernetes-1.12.4 && \
    go install ./... ) && \
    ( cd $GOPATH/src/k8s.io/kube-openapi/cmd/openapi-gen && \
    git checkout 0317810137be915b9cf888946c6e115c1bfac693 && \
    go install ./... ) && \
    (go get -u -d github.com/Masterminds/glide && \
     cd $GOPATH/src/github.com/Masterminds/glide && \
     git checkout v0.13.2 -b build-v0.13.2 && \
     go install ./...) && \
    (curl -L -O "https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH}.tar.gz" && \
     tar -zxvf kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH}.tar.gz && \
     mv kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH} /usr/local/kubebuilder && \
     rm kubebuilder_${KUBEBUILDER_VERSION}_linux_${ARCH}.tar.gz )

ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

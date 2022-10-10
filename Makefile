all: build
.PHONY: all

build:
	go build github.com/efucloud/api/...
.PHONY: build


verify:
	chmod a+x ./vendor/k8s.io/code-generator/generate-groups.sh
	hack/verify-deepcopy.sh
	hack/verify-protobuf.sh
	hack/verify-swagger-docs.sh
	hack/verify-codegen.sh
.PHONY: verify


.PHONY: generate
generate:
	chmod a+x ./vendor/k8s.io/code-generator/generate-groups.sh
	hack/update-deepcopy.sh
	hack/update-protobuf.sh
	hack/update-swagger-docs.sh
	hack/update-codegen.sh
	controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=manifests/crd


CONTROLLER_GEN = $(GOPATH)/bin/controller-gen

.PHONY: manifests
manifests: generate
	controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=manifests/crd


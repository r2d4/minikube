# Copyright 2016 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Bump these on release
VERSION_MAJOR ?= 0
VERSION_MINOR ?= 12
VERSION_BUILD ?= 1
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)
DEB_VERSION ?= $(VERSION_MAJOR).$(VERSION_MINOR)-$(VERSION_BUILD)
INSTALL_SIZE ?= $(shell du out/minikube-windows-amd64.exe | cut -f1)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BUILD_DIR ?= ./out
ORG := k8s.io
REPOPATH ?= $(ORG)/minikube
BUILD_IMAGE ?= gcr.io/google_containers/kube-cross:v1.7.1-0
IS_EXE ?=

ifeq ($(IN_DOCKER),1)
	GOPATH := /go
else
	GOPATH := $(shell pwd)/_gopath
endif

export GOPATH


# Use system python if it exists, otherwise use Docker.
PYTHON := $(shell command -v python || echo "docker run --rm -it -v $(shell pwd):/minikube -w /minikube python python")
BUILD_OS := $(shell uname -s)

# Set the version information for the Kubernetes servers, and build localkube statically
K8S_VERSION_LDFLAGS := $(shell $(PYTHON) hack/get_k8s_version.py 2>&1)
MINIKUBE_LDFLAGS := -X k8s.io/minikube/pkg/version.version=$(VERSION)
LOCALKUBE_LDFLAGS := "$(K8S_VERSION_LDFLAGS) $(MINIKUBE_LDFLAGS) -s -w -extldflags '-static'"

LOCALKUBEFILES := GOPATH=$(GOPATH) go list  -f '{{join .Deps "\n"}}' ./cmd/localkube/ | grep k8s.io | GOPATH=$(GOPATH) xargs go list -f '{{ range $$file := .GoFiles }} {{$$.Dir}}/{{$$file}}{{"\n"}}{{end}}'
MINIKUBEFILES := GOPATH=$(GOPATH) go list  -f '{{join .Deps "\n"}}' ./cmd/minikube/ | grep k8s.io | GOPATH=$(GOPATH) xargs go list -f '{{ range $$file := .GoFiles }} {{$$.Dir}}/{{$$file}}{{"\n"}}{{end}}'

ifeq ($(GOOS),windows)
	IS_EXE = ".exe"
endif
out/minikube$(IS_EXE): out/minikube-$(GOOS)-$(GOARCH)$(IS_EXE)
	cp $(BUILD_DIR)/minikube-$(GOOS)-$(GOARCH)$(IS_EXE) $(BUILD_DIR)/minikube$(IS_EXE)

out/localkube: $(GOPATH)/src/$(ORG) $(shell $(LOCALKUBEFILES))
ifeq ($(BUILD_OS),Linux)
	CGO_ENABLED=1 go build -ldflags=$(LOCALKUBE_LDFLAGS) -o $(BUILD_DIR)/localkube ./cmd/localkube
else
	docker run -w /go/src/$(REPOPATH) -e IN_DOCKER=1 -v $(shell pwd):/go/src/$(REPOPATH) $(BUILD_IMAGE) make out/localkube
endif

out/minikube-darwin-amd64: $(GOPATH)/src/$(ORG) pkg/minikube/assets/assets.go $(shell $(MINIKUBEFILES))
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build --installsuffix cgo -ldflags="$(MINIKUBE_LDFLAGS) $(K8S_VERSION_LDFLAGS)" -a -o $(BUILD_DIR)/minikube-darwin-amd64 ./cmd/minikube

out/minikube-linux-amd64: $(GOPATH)/src/$(ORG) pkg/minikube/assets/assets.go $(shell $(MINIKUBEFILES))
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build --installsuffix cgo -ldflags="$(MINIKUBE_LDFLAGS) $(K8S_VERSION_LDFLAGS)" -a -o $(BUILD_DIR)/minikube-linux-amd64 ./cmd/minikube

out/minikube-windows-amd64.exe: $(GOPATH)/src/$(ORG) pkg/minikube/assets/assets.go $(shell $(MINIKUBEFILES))
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build --installsuffix cgo -ldflags="$(MINIKUBE_LDFLAGS) $(K8S_VERSION_LDFLAGS)" -a -o $(BUILD_DIR)/minikube-windows-amd64.exe ./cmd/minikube

localkube-image: out/localkube
	make -C deploy/docker VERSION=$(VERSION)

iso:
	cd deploy/iso/boot2docker && ./build.sh

minikube-iso:
	[ ! -e $(BUILD_DIR)/buildroot ] && \
		mkdir -p $(BUILD_DIR) && \
		git clone https://github.com/buildroot/buildroot $(BUILD_DIR)/buildroot && \
		cd $(BUILD_DIR)/buildroot && \
		git checkout 2016.08 && \
		make BR2_EXTERNAL=../../deploy/iso/minikube-iso minikube_defconfig && \
		cd ../..
	cd $(BUILD_DIR)/buildroot && make

.PHONY: integration
integration: out/minikube
	go test -v $(REPOPATH)/test/integration --tags=integration

.PHONY: test
test: $(GOPATH)/src/$(ORG) pkg/minikube/assets/assets.go
	./test.sh

pkg/minikube/assets/assets.go: out/localkube $(GOPATH)/bin/go-bindata deploy/addons
	$(GOPATH)/bin/go-bindata -nomemcopy -o pkg/minikube/assets/assets.go -pkg assets ./out/localkube deploy/addons

$(GOPATH)/bin/go-bindata: $(GOPATH)/src/$(ORG)
	GOBIN=$(GOPATH)/bin go get github.com/jteeuwen/go-bindata/...

.PHONY: cross
cross: out/localkube out/minikube-linux-amd64 out/minikube-darwin-amd64 out/minikube-windows-amd64.exe

.PHONE: checksum
checksum:
	for f in out/localkube out/minikube-linux-amd64 out/minikube-darwin-amd64 out/minikube-windows-amd64.exe ; do \
		if [ -f "$${f}" ]; then \
			openssl sha256 "$${f}" | awk '{print $$2}' > "$${f}.sha256" ; \
		fi ; \
	done

.PHONY: clean
clean:
	rm -rf $(GOPATH)
	rm -rf $(BUILD_DIR)
	rm -f pkg/minikube/assets/assets.go

.PHONY: gendocs
gendocs: docs/minikube.md

docs/minikube.md: $(GOPATH)/src/$(ORG) $(shell find cmd) $(shell find pkg/minikube/constants) pkg/minikube/assets/assets.go
	cd $(GOPATH)/src/$(REPOPATH) && go run -ldflags="$(K8S_VERSION_LDFLAGS) $(MINIKUBE_LDFLAGS)" -tags gendocs gen_help_text.go

out/minikube_$(DEB_VERSION).deb: out/minikube-linux-amd64
	cp -r installers/linux/deb/minikube_deb_template out/minikube_$(DEB_VERSION)
	chmod 0755 out/minikube_$(DEB_VERSION)/DEBIAN
	sed -E -i 's/--VERSION--/'$(DEB_VERSION)'/g' out/minikube_$(DEB_VERSION)/DEBIAN/control
	mkdir -p out/minikube_$(DEB_VERSION)/usr/bin
	cp out/minikube-linux-amd64 out/minikube_$(DEB_VERSION)/usr/bin
	dpkg-deb --build out/minikube_$(DEB_VERSION)
	rm -rf out/minikube_$(DEB_VERSION)

out/minikube-installer.exe: out/minikube-windows-amd64.exe
	rm -rf out/windows_tmp
	cp -r installers/windows/ out/windows_tmp
	cp -r LICENSE out/windows_tmp/LICENSE
	awk 'sub("$$", "\r")' out/windows_tmp/LICENSE > out/windows_tmp/LICENSE.txt
	sed -E -i 's/--VERSION_MAJOR--/'$(VERSION_MAJOR)'/g' out/windows_tmp/minikube.nsi
	sed -E -i 's/--VERSION_MINOR--/'$(VERSION_MINOR)'/g' out/windows_tmp/minikube.nsi
	sed -E -i 's/--VERSION_BUILD--/'$(VERSION_BUILD)'/g' out/windows_tmp/minikube.nsi
	sed -E -i 's/--INSTALL_SIZE--/'$(INSTALL_SIZE)'/g' out/windows_tmp/minikube.nsi
	cp out/minikube-windows-amd64.exe out/windows_tmp/minikube.exe
	makensis out/windows_tmp/minikube.nsi
	mv out/windows_tmp/minikube-installer.exe out/minikube-installer.exe
	rm -rf out/windows_tmp

.PHONY: gopath
gopath: $(GOPATH)/src/$(ORG)

$(GOPATH)/src/$(ORG):
	mkdir -p $(GOPATH)/src/$(ORG)
	ln -s -f $(shell pwd) $(GOPATH)/src/$(ORG)

.PHONY: check-release
check-release:
	go test -v ./deploy/minikube/release_sanity_test.go -tags=release

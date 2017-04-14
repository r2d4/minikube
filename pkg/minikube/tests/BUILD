load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "api_mock.go",
        "dir_utils.go",
        "driver_mock.go",
        "host_mock.go",
        "provision_mock.go",
        "ssh_mock.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//vendor/github.com/docker/machine/libmachine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/auth:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/engine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnerror:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnflag:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision/pkgaction:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision/serviceaction:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/swarm:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)

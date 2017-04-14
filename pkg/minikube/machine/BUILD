load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "client_linux.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//vendor/github.com/docker/machine/drivers/virtualbox:go_default_library",
        "//vendor/github.com/docker/machine/libmachine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/auth:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/cert:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/check:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers/plugin:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers/plugin/localbinary:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers/rpc:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/engine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/persist:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/ssh:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/swarm:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/version:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["client_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//vendor/github.com/docker/machine/drivers/virtualbox:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers/plugin/localbinary:go_default_library",
    ],
)

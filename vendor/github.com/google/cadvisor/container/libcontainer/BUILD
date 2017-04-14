load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["helpers.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/container:go_default_library",
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups:go_default_library",
    ],
)

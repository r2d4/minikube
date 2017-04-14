load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["selinux.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opencontainers/runc/libcontainer/system:go_default_library"],
)

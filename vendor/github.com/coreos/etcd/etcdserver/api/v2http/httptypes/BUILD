load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "errors.go",
        "member.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/etcd/pkg/types:go_default_library",
        "//vendor/github.com/coreos/pkg/capnslog:go_default_library",
    ],
)

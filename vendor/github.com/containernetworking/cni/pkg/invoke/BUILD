load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "args.go",
        "delegate.go",
        "exec.go",
        "find.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/containernetworking/cni/pkg/types:go_default_library"],
)

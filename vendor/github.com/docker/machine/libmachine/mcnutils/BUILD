load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "b2d.go",
        "utils.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/libmachine/log:go_default_library",
        "//vendor/github.com/docker/machine/version:go_default_library",
    ],
)

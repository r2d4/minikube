load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["error.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnflag:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
    ],
)

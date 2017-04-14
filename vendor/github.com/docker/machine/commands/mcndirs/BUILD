load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["utils.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/machine/libmachine/mcnutils:go_default_library"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "filestore.go",
        "store.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnerror:go_default_library",
    ],
)

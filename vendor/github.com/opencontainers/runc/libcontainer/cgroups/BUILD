load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cgroups.go",
        "stats.go",
        "utils.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
    ],
)

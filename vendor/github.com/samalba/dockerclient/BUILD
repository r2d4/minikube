load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "auth.go",
        "dockerclient.go",
        "example_responses.go",
        "interface.go",
        "tls.go",
        "types.go",
        "utils.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/go-units:go_default_library"],
)

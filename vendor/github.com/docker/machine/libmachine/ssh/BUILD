load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "keys.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/log:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnutils:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/terminal:go_default_library",
    ],
)

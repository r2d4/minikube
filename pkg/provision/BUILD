load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["buildroot.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/util:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/auth:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/engine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/log:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision/pkgaction:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/swarm:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["libmachine.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/drivers/errdriver:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/auth:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/cert:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/check:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers/plugin/localbinary:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers/rpc:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/engine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/log:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnerror:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnutils:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/persist:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/ssh:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/swarm:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/version:go_default_library",
    ],
)

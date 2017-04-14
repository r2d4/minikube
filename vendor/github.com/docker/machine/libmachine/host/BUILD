load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "host.go",
        "host_v0.go",
        "host_v2.go",
        "migrate.go",
        "migrate_v0_v1.go",
        "migrate_v1_v2.go",
        "migrate_v2_v3.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/drivers/none:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/auth:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/engine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/log:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcndockerclient:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnerror:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnutils:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision/pkgaction:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision/serviceaction:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/ssh:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/swarm:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/version:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/versioncmp:go_default_library",
    ],
)

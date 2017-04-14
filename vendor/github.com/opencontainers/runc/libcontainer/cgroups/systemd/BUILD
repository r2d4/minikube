load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["apply_systemd.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-systemd/dbus:go_default_library",
        "//vendor/github.com/coreos/go-systemd/util:go_default_library",
        "//vendor/github.com/godbus/dbus:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups/fs:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
    ],
)

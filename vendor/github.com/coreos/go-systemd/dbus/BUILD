load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dbus.go",
        "methods.go",
        "properties.go",
        "set.go",
        "subscription.go",
        "subscription_set.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/godbus/dbus:go_default_library"],
)

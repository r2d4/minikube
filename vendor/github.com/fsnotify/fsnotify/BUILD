load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fsnotify.go",
        "inotify.go",
        "inotify_poller.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/sys/unix:go_default_library"],
)

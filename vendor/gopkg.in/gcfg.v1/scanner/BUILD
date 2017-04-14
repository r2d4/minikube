load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "errors.go",
        "scanner.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/gcfg.v1/token:go_default_library"],
)

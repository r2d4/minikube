load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "go1_2.go",
        "read.go",
        "set.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/gopkg.in/gcfg.v1/scanner:go_default_library",
        "//vendor/gopkg.in/gcfg.v1/token:go_default_library",
        "//vendor/gopkg.in/gcfg.v1/types:go_default_library",
    ],
)

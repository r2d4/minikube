load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["oomparser.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/utils:go_default_library",
        "//vendor/github.com/google/cadvisor/utils/tail:go_default_library",
    ],
)

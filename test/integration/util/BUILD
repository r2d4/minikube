load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["util.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/util:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["util.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/apimachinery/pkg/api/validation:go_default_library"],
)

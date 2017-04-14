load("@io_bazel_rules_go//go:def.bzl", "cgo_library", "go_library")

cgo_library(
    name = "cgo_default_library",
    srcs = ["util.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/coreos/pkg/dlopen:go_default_library"],
)

go_library(
    name = "go_default_library",
    library = ":cgo_default_library",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/coreos/pkg/dlopen:go_default_library"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "errors.go",
        "requests.go",
        "results.go",
        "urls.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mitchellh/mapstructure:go_default_library",
        "//vendor/github.com/rackspace/gophercloud:go_default_library",
    ],
)

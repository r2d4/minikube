load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["csr.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/cloudflare/cfssl/errors:go_default_library",
        "//vendor/github.com/cloudflare/cfssl/helpers:go_default_library",
        "//vendor/github.com/cloudflare/cfssl/log:go_default_library",
    ],
)

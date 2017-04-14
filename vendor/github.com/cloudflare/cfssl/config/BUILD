load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["config.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/cloudflare/cfssl/auth:go_default_library",
        "//vendor/github.com/cloudflare/cfssl/errors:go_default_library",
        "//vendor/github.com/cloudflare/cfssl/helpers:go_default_library",
        "//vendor/github.com/cloudflare/cfssl/log:go_default_library",
        "//vendor/github.com/cloudflare/cfssl/ocsp/config:go_default_library",
    ],
)

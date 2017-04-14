load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["pkix.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/google/certificate-transparency/go/asn1:go_default_library"],
)

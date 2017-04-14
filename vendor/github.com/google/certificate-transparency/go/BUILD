load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "serialization.go",
        "signatures.go",
        "types.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/google/certificate-transparency/go/x509:go_default_library"],
)

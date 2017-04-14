load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "checksum_validator.go",
        "download.go",
        "retries.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/go-multierror:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/gopkg.in/cheggaaa/pb.v1:go_default_library",
    ],
)

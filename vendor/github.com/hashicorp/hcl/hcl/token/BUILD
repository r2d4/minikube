load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "position.go",
        "token.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/hcl/hcl/strconv:go_default_library"],
)

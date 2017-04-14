load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "descriptor.pb.go",
        "gostring.go",
        "helper.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto:go_default_library"],
)

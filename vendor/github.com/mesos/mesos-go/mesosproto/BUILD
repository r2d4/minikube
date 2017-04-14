load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "authentication.pb.go",
        "internal.pb.go",
        "mesos.pb.go",
        "messages.pb.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto:go_default_library"],
)

filegroup(
    name = "go_default_library_protos",
    srcs = [
        "authentication.proto",
        "internal.proto",
        "mesos.proto",
        "messages.proto",
    ],
    visibility = ["//visibility:public"],
)

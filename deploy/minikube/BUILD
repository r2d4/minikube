load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["schema_check.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/xeipuuv/gojsonschema:go_default_library"],
)

go_binary(
    name = "minikube",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

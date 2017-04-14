load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//cmd/minikube/cmd:go_default_library",
        "//pkg/minikube/machine:go_default_library",
        "//pkg/provision:go_default_library",
        "//vendor/github.com/pkg/profile:go_default_library",
    ],
)

go_binary(
    name = "minikube",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_prefix")

go_prefix("k8s.io/minikube")

go_library(
    name = "go_default_library",
    srcs = ["gen_help_text.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//cmd/minikube/cmd:go_default_library",
        "//vendor/github.com/spf13/cobra/doc:go_default_library",
    ],
)

go_binary(
    name = "minikube",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

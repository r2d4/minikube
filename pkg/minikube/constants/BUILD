load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "constants.go",
        "constants_linux.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/version:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
        "//vendor/k8s.io/client-go/util/homedir:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/version:go_default_library",
    ],
)

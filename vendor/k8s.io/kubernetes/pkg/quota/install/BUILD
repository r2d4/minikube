package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["registry.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/quota:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/quota/evaluator/core:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)

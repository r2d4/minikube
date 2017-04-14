load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["kubernetes_versions.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["kubernetes_versions_test.go"],
    library = ":go_default_library",
)

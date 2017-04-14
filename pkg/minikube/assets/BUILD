load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "addons.go",
        "assets.go",
        "vm_assets.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/config:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/util:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

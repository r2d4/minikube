load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["notify.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/config:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/version:go_default_library",
        "//vendor/github.com/blang/semver:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/spf13/viper:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["notify_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/config:go_default_library",
        "//pkg/minikube/tests:go_default_library",
        "//pkg/version:go_default_library",
        "//vendor/github.com/blang/semver:go_default_library",
        "//vendor/github.com/spf13/viper:go_default_library",
    ],
)

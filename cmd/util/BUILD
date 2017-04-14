load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["util.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//cmd/minikube/cmd/config:go_default_library",
        "//pkg/minikube/config:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/version:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/spf13/viper:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/terminal:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["util_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg/version:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["config.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd/api:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd/api/latest:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["config_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd/api:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "constants.go",
        "crypto.go",
        "downloader.go",
        "extra_options.go",
        "utils.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//pkg/minikube/kubernetes_versions:go_default_library",
        "//pkg/version:go_default_library",
        "//vendor/github.com/blang/semver:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/jimmidyson/go-download:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "config_test.go",
        "crypto_test.go",
        "downloader_test.go",
        "extra_options_test.go",
        "utils_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//pkg/minikube/tests:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
    ],
)

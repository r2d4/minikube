load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["webhook.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/authentication/authenticator:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/cache:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/webhook:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/typed/authentication/v1beta1:go_default_library",
        "//vendor/k8s.io/client-go/pkg/api:go_default_library",
        "//vendor/k8s.io/client-go/pkg/apis/authentication/install:go_default_library",
        "//vendor/k8s.io/client-go/pkg/apis/authentication/v1beta1:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["service.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/cluster:go_default_library",
        "//pkg/util:go_default_library",
        "//vendor/github.com/docker/machine/libmachine:go_default_library",
        "//vendor/github.com/pkg/browser:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/typed/core/v1:go_default_library",
        "//vendor/k8s.io/client-go/pkg/api/v1:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["service_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/constants:go_default_library",
        "//pkg/minikube/tests:go_default_library",
        "//vendor/github.com/docker/machine/libmachine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/typed/core/v1:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/typed/core/v1/fake:go_default_library",
        "//vendor/k8s.io/client-go/pkg/api/v1:go_default_library",
    ],
)

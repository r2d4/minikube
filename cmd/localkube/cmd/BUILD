load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "options.go",
        "root.go",
        "start.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/localkube:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/util:go_default_library",
        "//pkg/version:go_default_library",
        "//vendor/github.com/coreos/pkg/capnslog:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/feature:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/capabilities:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/types:go_default_library",
    ],
)

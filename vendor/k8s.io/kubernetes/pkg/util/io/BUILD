package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "io.go",
        "writer.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api/v1:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["io_test.go"],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/util/io:go_default_library",
        "//pkg/volume:go_default_library",
        "//vendor:github.com/pborman/uuid",
        "//vendor:k8s.io/apimachinery/pkg/api/equality",
        "//vendor:k8s.io/apimachinery/pkg/runtime",
        "//vendor:k8s.io/client-go/util/testing",
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

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["sshutil.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/assets:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/ssh:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["sshutil_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/tests:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
    ],
)

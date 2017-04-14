load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cluster.go",
        "cluster_linux.go",
        "cluster_non_darwin_panic.go",
        "cluster_non_windows_panic.go",
        "commands.go",
        "credentials.go",
        "localkube_caching.go",
        "types.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/assets:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/minikube/sshutil:go_default_library",
        "//pkg/util:go_default_library",
        "//vendor/github.com/docker/machine/drivers/virtualbox:go_default_library",
        "//vendor/github.com/docker/machine/libmachine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/engine:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/jimmidyson/go-download:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cluster_test.go",
        "commands_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/assets:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/minikube/tests:go_default_library",
        "//pkg/util:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/host:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/provision:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
    ],
)

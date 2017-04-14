load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "addons.go",
        "addons_list.go",
        "config.go",
        "config_view.go",
        "configure.go",
        "disable.go",
        "enable.go",
        "get.go",
        "open.go",
        "prompt.go",
        "set.go",
        "unset.go",
        "util.go",
        "validations.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/minikube/assets:go_default_library",
        "//pkg/minikube/cluster:go_default_library",
        "//pkg/minikube/config:go_default_library",
        "//pkg/minikube/constants:go_default_library",
        "//pkg/minikube/machine:go_default_library",
        "//pkg/minikube/service:go_default_library",
        "//pkg/minikube/sshutil:go_default_library",
        "//pkg/minikube/storageclass:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/viper:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "config_test.go",
        "disable_test.go",
        "enable_test.go",
        "set_test.go",
        "util_test.go",
        "validations_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//pkg/minikube/assets:go_default_library",
        "//pkg/minikube/config:go_default_library",
        "//pkg/minikube/sshutil:go_default_library",
        "//pkg/minikube/tests:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "disk.go",
        "ip.go",
        "misc.go",
        "network.go",
        "vbm.go",
        "virtualbox.go",
        "virtualbox_linux.go",
        "vm.go",
        "vtx.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/libmachine/drivers:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/log:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnflag:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/mcnutils:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/ssh:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/state:go_default_library",
    ],
)

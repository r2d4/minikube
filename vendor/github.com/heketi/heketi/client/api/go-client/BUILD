load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "backup.go",
        "client.go",
        "cluster.go",
        "device.go",
        "node.go",
        "topology.go",
        "volume.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dgrijalva/jwt-go:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/glusterfs/api:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
    ],
)

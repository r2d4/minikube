load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "inmem_socket.go",
        "proxy.go",
        "sockets.go",
        "sockets_unix.go",
        "tcp_socket.go",
        "unix_socket.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/user:go_default_library",
        "//vendor/golang.org/x/net/proxy:go_default_library",
    ],
)

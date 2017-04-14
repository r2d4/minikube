load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "attrs.go",
        "attrs_unix.go",
        "client.go",
        "conn.go",
        "packet.go",
        "release.go",
        "server.go",
        "server_statvfs_impl.go",
        "server_statvfs_linux.go",
        "server_unix.go",
        "sftp.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/kr/fs:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)

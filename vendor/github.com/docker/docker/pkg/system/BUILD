load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "chtimes.go",
        "chtimes_unix.go",
        "errors.go",
        "filesys.go",
        "lstat.go",
        "meminfo.go",
        "meminfo_linux.go",
        "mknod.go",
        "path_unix.go",
        "stat.go",
        "stat_linux.go",
        "syscall_unix.go",
        "umask.go",
        "utimes_linux.go",
        "xattrs_linux.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/go-units:go_default_library"],
)

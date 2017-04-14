load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dmsetup_client.go",
        "doc.go",
        "thin_ls_client.go",
        "thin_pool_watcher.go",
        "util.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/glog:go_default_library"],
)

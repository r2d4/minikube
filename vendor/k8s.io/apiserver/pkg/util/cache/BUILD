load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cache.go",
        "lruexpirecache.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/groupcache/lru:go_default_library"],
)

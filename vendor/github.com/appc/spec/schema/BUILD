load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "image.go",
        "kind.go",
        "pod.go",
        "version.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/go4.org/errorutil:go_default_library",
    ],
)

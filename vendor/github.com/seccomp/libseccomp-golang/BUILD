load("@io_bazel_rules_go//go:def.bzl", "cgo_library", "go_library")

cgo_library(
    name = "cgo_default_library",
    srcs = [
        "seccomp.go",
        "seccomp_internal.go",
    ],
    clinkopts = [
        "-lseccomp",
        "-lseccomp",
    ],
    visibility = ["//visibility:private"],
)

go_library(
    name = "go_default_library",
    library = ":cgo_default_library",
    visibility = ["//visibility:public"],
)

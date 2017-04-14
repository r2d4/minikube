load("@io_bazel_rules_go//go:def.bzl", "cgo_library", "go_library")

cgo_library(
    name = "cgo_default_library",
    srcs = ["tc_linux_cgo.go"],
    visibility = ["//visibility:private"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "ascii.go",
        "term.go",
    ],
    library = ":cgo_default_library",
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "delegate.go",
        "doc.go",
        "results.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mitchellh/mapstructure:go_default_library",
        "//vendor/github.com/rackspace/gophercloud:go_default_library",
        "//vendor/github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumes:go_default_library",
        "//vendor/github.com/rackspace/gophercloud/pagination:go_default_library",
    ],
)

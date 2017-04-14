load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "auth_env.go",
        "client.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/rackspace/gophercloud:go_default_library",
        "//vendor/github.com/rackspace/gophercloud/openstack:go_default_library",
        "//vendor/github.com/rackspace/gophercloud/openstack/utils:go_default_library",
        "//vendor/github.com/rackspace/gophercloud/rackspace/identity/v2/tokens:go_default_library",
    ],
)

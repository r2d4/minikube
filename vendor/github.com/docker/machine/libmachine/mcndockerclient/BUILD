load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "docker_client.go",
        "docker_host.go",
        "docker_versioner.go",
        "fake_docker_versioner.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/machine/libmachine/auth:go_default_library",
        "//vendor/github.com/docker/machine/libmachine/cert:go_default_library",
        "//vendor/github.com/samalba/dockerclient:go_default_library",
    ],
)

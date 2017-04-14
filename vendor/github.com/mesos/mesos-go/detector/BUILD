load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "factory.go",
        "interface.go",
        "standalone.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/mesos/mesos-go/mesosproto:go_default_library",
        "//vendor/github.com/mesos/mesos-go/mesosutil:go_default_library",
        "//vendor/github.com/mesos/mesos-go/upid:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

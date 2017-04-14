load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client2.go",
        "detect.go",
        "doc.go",
        "plugin.go",
        "types.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/mesos/mesos-go/detector:go_default_library",
        "//vendor/github.com/mesos/mesos-go/mesosproto:go_default_library",
        "//vendor/github.com/samuel/go-zookeeper/zk:go_default_library",
    ],
)

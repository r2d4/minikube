load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "apply_raw.go",
        "blkio.go",
        "cpu.go",
        "cpuacct.go",
        "cpuset.go",
        "devices.go",
        "freezer.go",
        "hugetlb.go",
        "memory.go",
        "name.go",
        "net_cls.go",
        "net_prio.go",
        "perf_event.go",
        "pids.go",
        "utils.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/system:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/utils:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "collector.go",
        "counter.go",
        "desc.go",
        "doc.go",
        "expvar.go",
        "gauge.go",
        "go_collector.go",
        "histogram.go",
        "http.go",
        "metric.go",
        "process_collector.go",
        "push.go",
        "registry.go",
        "summary.go",
        "untyped.go",
        "value.go",
        "vec.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/beorn7/perks/quantile:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/prometheus/client_model/go:go_default_library",
        "//vendor/github.com/prometheus/common/expfmt:go_default_library",
        "//vendor/github.com/prometheus/procfs:go_default_library",
    ],
)

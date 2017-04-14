load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cert_pool.go",
        "pem_decrypt.go",
        "pkcs1.go",
        "pkcs8.go",
        "root.go",
        "root_unix.go",
        "sec1.go",
        "verify.go",
        "x509.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/certificate-transparency/go/asn1:go_default_library",
        "//vendor/github.com/google/certificate-transparency/go/x509/pkix:go_default_library",
    ],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bar",
    srcs = ["bar.go"],
    importpath = "github.com/dt/externdemo/pkg/bar",
    visibility = ["//visibility:public"],
    deps = ["//pkg/foo"],
)

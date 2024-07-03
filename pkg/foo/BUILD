load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "foo",
    srcs = ["foo.go"],
    importpath = "github.com/dt/externdemo/pkg/foo",
    visibility = ["//visibility:public"],
)

go_test(
    name = "foo_test",
    srcs = [
        "foo_external_test.go",
        "foo_test.go",
    ],
    embed = [":foo"],
    deps = ["//pkg/bar"],
)

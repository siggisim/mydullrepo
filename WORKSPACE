load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "gcloud",
    build_file_content = """package(default_visibility = ["//visibility:public"])\nexports_files(["gcloud", "gsutil", "bq"])""",
    patch_cmds = [
        "ln -s google-cloud-sdk/bin/gcloud gcloud",
        "ln -s google-cloud-sdk/bin/gsutil gsutil",
        "ln -s google-cloud-sdk/bin/bq bq",
    ],
    urls = ["https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-440.0.0-linux-x86_64.tar.gz"],
)

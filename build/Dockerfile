# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1.20-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

COPY Makefile .
RUN make setup

RUN ls -a
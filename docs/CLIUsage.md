# CLI Usage

This document demonstrates how to use the CLI but also shows what happens in `Kudo` under the hood, which can be helpful when troubleshooting.

## Kubectl kudo Plugin

### Requirements

- GitHub set up
  - GitHub OTP token in `$HOME/.git-credentials`
  - GitHub Basic Auth via `GIT_USER` and `GIT_PASSWORD` environment variables exposed

### Install

Install the plugin from your `$GOPATH/src/github.com/kudobuilder/kudo` root folder via:

- `go install ./cmd/kubectl-kudo`

### Commands

|  Syntax | Description  | 
|---|---|
| `kubectl kudo install --frameworkname=<FrameworkName>`  |  Installs a framework from the `kudo App Store` with dependencies | 

.PHONY: binary

# For now, it is v1
APIVersion=v1

# The building time, something like:
# 2016-01-15.02:12:02
BuildDate=$(shell date -u +%Y-%m-%d.%H:%M:%S)

# Check whether the git repository is dirty or not.
GitDirty=$(shell git status --porcelain --untracked-files=no)
GitCommit=$(shell git rev-parse --short HEAD)
ifneq ($(GitDirty),)
	GitCommit:= $(GitCommit)-dirty
endif

# Which will be used in go build
GO_LDFLAGS=-ldflags "-X `go list ./version`.APIVersion=$(APIVersion) \
	   -X `go list ./version`.GitCommit=$(GitCommit) \
	   -X `go list ./version`.BuildDate=$(BuildDate)"

binary:
	go build $(GO_LDFLAGS)


#
# Copyright (c) 2020 HyTrust, Inc. All Rights Reserved.
#
THISDIR = $(shell pwd)

GOCMD = /usr/local/go/bin/go

KMIPCLI_SRCDIR = $(THISDIR)
PARENTDIR = $(THISDIR)/..

BUILDDIR = kmipcli-build
WORKSPACE = $(PARENTDIR)/$(BUILDDIR)
WORKSPACE_SRCDIR = $(WORKSPACE)/src
WORKSPACE_BINDIR = $(WORKSPACE)/bin

KMIPCLI_MAINGO = main.go

KMIPCLI_LINUX = kmipcli

all:
	@/usr/bin/echo "Clearing old Workspace if any at $(WORKSPACE).."
	@/usr/bin/rm -rf $(WORKSPACE)
	@/usr/bin/echo "Creating new kmipcli Workspace at $(WORKSPACE).."
	@/usr/bin/mkdir $(WORKSPACE)
	@/usr/bin/mkdir $(WORKSPACE_SRCDIR)
	@/usr/bin/cp -r $(KMIPCLI_SRCDIR)/. $(WORKSPACE_SRCDIR)
	@/usr/bin/cp -r $(PARENTDIR)/getpasswd/ $(WORKSPACE_SRCDIR)/
	@/usr/bin/cp -r $(PARENTDIR)/cmd/. $(WORKSPACE_SRCDIR)/cmd/
	@cd $(WORKSPACE_SRCDIR) && $(GOCMD) mod init cli
	@cd $(WORKSPACE_SRCDIR) && $(GOCMD) mod tidy
	@/usr/bin/echo "Compiling kmipcli for Linux..."
	@cd $(WORKSPACE_SRCDIR) && (env GOOS=linux GOARCH=amd64 $(GOCMD) build -o $(WORKSPACE_BINDIR)/$(KMIPCLI_LINUX) $(KMIPCLI_MAINGO))
	@/usr/bin/echo "Please find Linux kmipcli binary, $(KMIPCLI_LINUX) at $(WORKSPACE_BINDIR)"

install:

clean:
	@/usr/bin/echo "Clearing old Workspace if any at $(WORKSPACE).."
	@/usr/bin/rm -rf $(WORKSPACE)
	@/usr/bin/echo "Clean up complete..."

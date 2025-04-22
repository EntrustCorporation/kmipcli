#
# Copyright (c) 2020 HyTrust, Inc. All Rights Reserved.
#
THISDIR = $(shell pwd)

GOCMD = /usr/local/go21/bin/go

KMIPCLI_SRCDIR = $(THISDIR)
PARENTDIR = $(THISDIR)/..

BUILDDIR = kmipcli-build
WORKSPACE = $(PARENTDIR)/$(BUILDDIR)
WORKSPACE_SRCDIR = $(WORKSPACE)/src
WORKSPACE_BINDIR = $(WORKSPACE)/bin

KMIPCLI_MAINGO = main.go

KMIPCLI_LINUX = kmipcli
KMIPCLI_WINDOWS = kmipcli.exe
KMIPCLI_MAC = kmipcli.app

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
	@/usr/bin/echo "Compiling kmipcli for Linux, Windows & Mac..."
	@cd $(WORKSPACE_SRCDIR) && (env GOOS=linux GOARCH=amd64 $(GOCMD) build -o $(WORKSPACE_BINDIR)/$(KMIPCLI_LINUX) $(KMIPCLI_MAINGO))
	@cd $(WORKSPACE_SRCDIR) && (env GOOS=windows GOARCH=amd64 $(GOCMD) build -o $(WORKSPACE_BINDIR)/$(KMIPCLI_WINDOWS) $(KMIPCLI_MAINGO))
	@cd $(WORKSPACE_SRCDIR) && (env GOOS=darwin GOARCH=amd64 $(GOCMD) build -o $(WORKSPACE_BINDIR)/$(KMIPCLI_MAC) $(KMIPCLI_MAINGO))
	@/usr/bin/echo "Please find respective Linux, Windows & Mac kmipcli binaries, $(KMIPCLI_LINUX), $(KMIPCLI_WINDOWS) & $(KMIPCLI_MAC) at $(WORKSPACE_BINDIR)"

install:

clean:
	@/usr/bin/echo "Clearing old Workspace if any at $(WORKSPACE).."
	@/usr/bin/rm -rf $(WORKSPACE)
	@/usr/bin/echo "Clean up complete..."

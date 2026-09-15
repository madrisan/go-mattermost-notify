## 1.3.3 -- (Sep 16, 2026)

SECURITY FIXES:

 * github.com/go-viper/mapstructure/v2: may leak sensitive information in
   logs when processing malformed data
   Affected versions: github.com/go-viper/mapstructure/v2 < 2.3.0
   Fix: 2.5.0
   See: https://pkg.go.dev/vuln/GO-2025-3787

 * Update golang.org/x/text and golang.org/x/sys to their latest patched
   releases, addressing further advisories reported against modules
   required by the project.

IMPROVEMENTS:

 * Add a Trivy vulnerability scan to the GitHub Workflows, with results
   published to the GitHub Security tab.
 * Update the Go compiler and golangci-lint version used in the build
   scripts.
 * Update GitHub Actions (setup-go, checkout, golangci-lint-action,
   codeql-action/upload-sarif, trivy-action) to versions that no longer
   run on deprecated Node.js runtimes, and address the upcoming CodeQL v3
   deprecation.
 * Drop github.com/golangci/golangci-lint and github.com/mitchellh/gox
   from go.mod: they had been pulled in as real module requirements by
   `make bootstrap`'s use of `go get -u` instead of `go install`, even
   though neither is imported by this module's code.

## 1.3.2 -- (Mar 21, 2025)

BUG FIXES:

 * Fix SA1019: "io/ioutil" has been deprecated since Go 1.19.
 * Fix issues in the GitHub workflows.

## 1.3.1 -- (Mar 21, 2025)

BUG FIXES:

 * Set the Content-Type header for outgoing HTTP requests.
 * Fix test cmd/post_test.go.

IMPROVEMENTS:

 * Update the Go version and dependencies.

## 1.3.0 -- (Apr 26, 2024)

SECURITY FIXES:

 * Bump golang.org/x/text from 0.3.7 to 0.3.8.

IMPROVEMENTS:

 * Add the OS/arch and compiler information to the output of the version
   command.
 * Require Go 1.17.10+ in the Containerfiles.

BUG FIXES:

 * Fix the --insecure flag for the get method.
 * Fix a formatting issue reported by gofmt.

## 1.2.1 -- (Mar 15, 2023) -- security update

SECURITY FIXES:

 * Bump golang.org/x/sys from 0.0.0-20220209214540-3681064d5158 to 0.1.0.

BUG FIXES:

 * Fix build error: 'undefined: unsafe.Slice'.

IMPROVEMENTS:

 * Update the Go compiler version.

## 1.2.0 -- (May 10, 2022)

FEATURES:

 * New command-line option --insecure (false by default).
 * New command-line option --timeout (10s by default).

BUG FIXES:

 * Fix golangci-lint execution in the GitHub Workflows.
 * Fix lint issues spotted by golint.

IMPROVEMENTS:

 * Update the Go dependencies.

## 1.1.1 -- (Jul 4, 2021)

IMPROVEMENTS:

 * Add a Containerfile for building a container running
   go-mattermost-notify, with better documentation for the containerized
   version.

BUG FIXES:

 * Do not repeat the default value for --level twice in the help message.
 * README: fix "cannot find package hcl/hcl/printer".

## 1.1.0 -- (Feb 13, 2021)

IMPROVEMENTS:

 * Split the mattermost package into multiple files and add more unit
   tests.
 * Handle errors in a way more consistent with cobra.

BUG FIXES:

 * Use a better post image for small devices.
 * Do not print the error message twice.

## 1.0.0 -- (Feb 10, 2021)

FEATURES:

 * New command get for executing HTTP GET queries to Mattermost.

IMPROVEMENTS:

 * Improve unit tests and code refactoring in cmd/post.go.

BUG FIXES:

 * Fix an issue reported by Go lint (gosimple).
 * Fix the URL of the Go Reference badge.

## 0.9.0 -- (Feb 7, 2021)

BUG FIXES:

 * Fix the setup of the URL and access code from the configuration file.

IMPROVEMENTS:

 * Add a project logo image.

## 0.7.0 -- (Feb 6, 2021)

IMPROVEMENTS:

 * Mark the project as stable.
 * Add a coverage script and more unit tests.

BUG FIXES:

 * Fix an issue reported by golangci-lint.
 * Fix source code formatting reported by gofmt.

## 0.6.1 -- (Feb 3, 2021)

IMPROVEMENTS:

 * Add a simple Dockerfile for go-mattermost-notify.
 * Manage all fatal/usage errors with a single function HandleError().

## 0.6.0 -- (Feb 2, 2021)

 * First public release.

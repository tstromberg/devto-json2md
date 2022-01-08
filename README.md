# devto-markdown

Your escape hatch from https://dev.to/ - this tool converts the dev.to JSON export into Markdown files suitable for usage by other static content management systems.

## Features

* Exports Markdown compatible files using Hugo-compatible front-matter fields
* Properly marks unpublished articles as "draft"

## Requirements

* JSON export of dev.to articles from https://dev.to/settings/account
* go v1.17 or newer

## Installation

`go install github.com/tstromberg/devto-json2md@latest`

## Usage

1. Visit https://dev.to/settings/account
2. Click check-box next to "Request an export of your content"
3. Click "Submit Data Request" button
4. Download the `devto-export-<date>.zip` file you should have received in your e-mail Inbox
5. Unzip `devto-export-<date.zip`
6. Run `devto-json2md articles.json .`

## Example

Run against an export from http://dev.to/tstromberg:

```
> devto-json2md ~/Downloads/articles.json /tmp
I0103 15:08:39.129692    6005 devto-json2md.go:71] #1: Writing out /tmp/another-puzzle-to-solve-3b9j.md ...
I0103 15:08:39.130528    6005 devto-json2md.go:71] #2: Writing out /tmp/dreaming-of-a-kernel-cafe-191f.md ...
I0103 15:08:39.130785    6005 devto-json2md.go:71] #3: Writing out /tmp/supply-vs-demand-20e7.md ...
I0103 15:08:39.131031    6005 devto-json2md.go:71] #4: Writing out /tmp/kernel-cafe-toward-alpha-1-28m6.md ...
I0103 15:08:39.131243    6005 devto-json2md.go:71] #5: Writing out /tmp/persistent-multi-user-docker-on-macos-32em.md ...
I0103 15:08:39.131468    6005 devto-json2md.go:71] #6: Writing out /tmp/reasoning-out-the-kernel-cafe-infrastructure-4a0m-temp-slug-9692600.md ...
I0103 15:08:39.131855    6005 devto-json2md.go:71] #7: Writing out /tmp/go-secondary-groups-a-kaniko-adventure-6mn.md ...
I0103 15:08:39.132118    6005 devto-json2md.go:71] #8: Writing out /tmp/tesla-model-y-family-camping-in-below-zero-temperatures-od3.md ...
I0103 15:08:39.132382    6005 devto-json2md.go:71] #9: Writing out /tmp/motivating-software-engineering-teams-n1n.md ...
I0103 15:08:39.132607    6005 devto-json2md.go:71] #10: Writing out /tmp/choose-open-source-3d6f-temp-slug-1805710.md ...
I0103 15:08:39.132850    6005 devto-json2md.go:71] #11: Writing out /tmp/the-makings-of-a-great-playbook-entry-2f43-temp-slug-6540710.md ...
I0103 15:08:39.133075    6005 devto-json2md.go:71] #12: Writing out /tmp/the-anatomy-of-a-great-playbook-entry-35od.md ...
```

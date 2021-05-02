# xprurl

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

`xrpurl` is a command line tool to make calling `rippled` servers easy.

## Installation

```
$ go get github.com/go-xrp/xrpurl
```

## CLI Parameters

| Shortcut | Longform | Required | Description |
|----------|----------|----------|-------------|
| `m` | `method` | yes | `rippled` API method |
| `d` | `data` | no | Post Body, JSON RPC `params` property only |
| `t` | `testdata` | no | Use test data request instead of custom `data` |
| `x` | `exec` | no | Execute API call |
| `v` | `verbose` | no | Verbose |
| `p` | `pretty` | no Pretty print JSON response |

## Examples

Example request include:

### Retrieve Account Lines with test data

```
$ xpurl -m account_lines -t -x -p
```

 [build-status-svg]: https://github.com/go-xrp/xrpurl/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/go-xrp/xrpurl/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/go-xrp/xrpurl
 [goreport-url]: https://goreportcard.com/report/github.com/go-xrp/xrpurl
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/go-xrp/xrpurl
 [docs-godoc-url]: https://pkg.go.dev/github.com/go-xrp/xrpurl
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/go-xrp/xrpurl/blob/master/LICENSE
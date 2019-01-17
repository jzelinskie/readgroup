# readgroup
[![Go Report Card](https://goreportcard.com/badge/github.com/jzelinskie/readgroup)](https://goreportcard.com/report/github.com/jzelinskie/readgroup)
[![GoDoc](https://godoc.org/github.com/jzelinskie/readgroup?status.svg)](https://godoc.org/github.com/jzelinskie/readgroup)
![Lines of Code](https://tokei.rs/b1/github/jzelinskie/readgroup)
[![License](https://img.shields.io/badge/license-BSD-blue.svg)](https://en.wikipedia.org/wiki/BSD_licenses#2-clause_license_.28.22Simplified_BSD_License.22_or_.22FreeBSD_License.22.29)

Have you ever wanted to process the exact same io.Reader data concurrently?
This is library attempts to paint over the synchronization footguns in doing so.

A canonical example is transcoding a video that is being uploaded to your server into multiple formats.
You don't wait for the whole video to be uploaded and have to keep the whole thing in memory before you can read over the entire set of bytes multiple times.
Instead, you want to fire off multiple goroutines that all get to read the same data as it comes in and block until they're all finished or one fails.

This library was heavily by [errgroup](https://golang.org/x/sync/errgroup) and mimics its API.

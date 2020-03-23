# mplayer

<div align="center">
  <a href="https://godoc.org/github.com/bisoncorps/mplayer">
    <img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" alt="Documentation">
  </a>
  <a href="https://goreportcard.com/report/github.com/bisoncorps/mplayer">
    <img src="https://goreportcard.com/badge/github.com/bisoncorps/mplayer" alt="Go Report Card">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT">
  </a>
</div>
A library to allow playing media files (remote and local) from your gocode


## Install
```
go get github.com/bisoncorps/mplayer
```

## What is mplayer
mplayer is just a means to allow for playing media files from golang. It allows for streaming online media
files and local media files as well. It hopes to incorporate functionalities that support using any known
media player. Currently only browser is supported but other players are currently being considered. It supports

- playing local files
- streaming online resource

## Usage
```go
package main

import  (
  "log"
  "github.com/bisoncorps/mplayer"
)

func main() {
		p, err := mplayer.GetPlayer("browser")
		if err != nil {
			log.Fatal(err)
		}
    // urls can be remote too
		p.URL("/home/manasseh/Videos/Jumanji.mp4")
		p.SetTitle("Jumanji MP4")
		p.Play()
}
```

## TODO
- [ ] Add other players (mpv, vlc)
- [ ] Add subtitle options

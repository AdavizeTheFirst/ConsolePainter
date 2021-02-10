package main

import (
	"fmt"
	"time"
)

type Screen struct {
	width       int
	height      int
	pixels      int
	returnChars int
	refreshRate time.Time
	frameStream []byte
	backgroundChar rune
}

type screenConfig struct {
	width       int
	height      int
	backgroundChar rune
	refreshRate time.Time
}

type renderInfo struct {
	startX int
	startY int
	endX int
	endY int
	frame []byte
}

type screenItem interface {
	getRenderInfo () *renderInfo renderInfo{}
}

func (s *Screen) GetPositionInStream(iPx int, jPx int) int {
	returnCharsInStream := iPx - 1
	return ((s.width * (iPx - 1)) - AbsInt(s.width-jPx) + returnCharsInStream)
}

func (s *Screen) Init(sc *screenConfig) {
	s.width = sc.width
	s.height = sc.height
	s.returnChars = sc.height;
	s.pixels = (s.width * s.height)  + s.returnChars
	s.frameStream = make([]byte, s.pixels, s.pixels)
	s.backgroundChar = sc.backgroundChar
}

func (s *Screen) render() {
		s.updateFrameStream()
		fmt.Printf(string(s.frameStream))
}

func (s *Screen) updateFrameStream(c rune) {
	s.resetStream()
	//loop through hooked items
}

func (s *Screen) resetStream(c rune) {
		// put default background 
		for i, _ := range s.frameStream {
			if i % (s.width + 1) == 0 {
				s.frameStream[i] = '\n'
				continue
			}
			s.frameStream[i] = byte(s.backgroundChar)
		}
}
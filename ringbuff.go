package main

import (
	"errors"
)

type Row interface {
	Stringfy() string
}

type PageBuff struct {
	// how much pages
	pages int
	// one page size
	pageBuffSize int
	// current pointer to  the table(tracking pointer)
	currIndex int
	// currnt page(slice) | this start with 0
	currPage int
	// this buffer hold buffSize * maxLoadPage amount of data
	buff      [][]string
	buffSize  int
	startPage []int
	endPage   []int
	extra     int
}

func NewPageBuff(initBuff [][]string, pagebuff int) PageBuff {
	bufflenth := len(initBuff)
	p := bufflenth / pagebuff
	return PageBuff{
		buff:         initBuff,
		buffSize:     bufflenth,
		pages:        p,
		pageBuffSize: pagebuff,
		currIndex:    0,
		currPage:     0,
		startPage:    []int{0, pagebuff - 1},
		endPage:      []int{((p - 1) * pagebuff), (p * pagebuff) - 1},
		extra:        bufflenth % pagebuff,
	}
}

// buff = |0|1|2|3|4|5|6|7|8|9|
// p.buffSize = 10
// p.pageBuffSize = 3
// p.startPage = [0, 2]

// buff = |0|1|2|3|4|5|6|*|*|*|
// p.buffSize = 10
// p.pageBuffSize = 3
// p.startPage = [7, 9]
func (p *PageBuff) PushToStart(set [][]string) {
	if p.pageBuffSize != len(set) {
		set = set[:p.buffSize]
	}
	if p.endPage[1] != (p.buffSize-p.extra)-1 {
		p.endPage[0], p.endPage[1] = p.StepPageIndex(p.endPage, 1)
	}
	fpoint := (p.startPage[0] - p.pageBuffSize + p.buffSize) % p.buffSize
	epoint := (p.startPage[1] - p.pageBuffSize + p.buffSize) % p.buffSize

	trackId := fpoint
	for _, i := range set {
		p.buff[trackId] = i
		trackId = (trackId + 1) % p.buffSize

	}

	p.startPage[0] = fpoint
	p.startPage[1] = epoint
}

// buff = |*|*|*|3|4|5|6|7|8|9|
// p.buffSize = 10
// p.pageBuffSize = 3
// p.startPage = [3, 5]
func (p *PageBuff) PushToEnd(set [][]string) {
	if p.pageBuffSize != len(set) {
		set = set[:p.buffSize]
	}
	if p.endPage[1] == (p.buffSize-p.extra)-1 {
		p.startPage[0] = (p.startPage[0] + p.pageBuffSize) % p.buffSize
		p.startPage[1] = (p.startPage[1] + p.pageBuffSize) % p.buffSize
	} else {
		p.endPage[0], p.endPage[1] = p.StepPageIndex(p.endPage, 1)
	}
	fpoint := (p.startPage[0] + p.buffSize - (p.pageBuffSize + p.extra)) % p.buffSize
	// epoint := (p.startPage[1] + p.buffSize - (p.pageBuffSize + p.extra)) % p.buffSize
	trackId := fpoint
	for _, v := range set {
		p.buff[trackId] = v
		trackId = (trackId + 1) % p.buffSize
	}

}

func (p *PageBuff) GetPageRealIndex(pageIndex int) (int, int, error) {
	if pageIndex > p.pages {
		return 0, 0, errors.New("Invalid page index")
	}
	start := (p.startPage[0] + (pageIndex * p.pageBuffSize)) % p.buffSize
	end := (p.startPage[1] + (pageIndex * p.pageBuffSize)) % p.buffSize
	return start, end, nil
}

func (p *PageBuff) StepPageIndex(page []int, step int) (int, int) {
	if step < 0 {
		s := p.pageBuffSize * step
		return page[0] + s, page[1] + s
	} else {
		s := p.pageBuffSize * step
		return page[0] + s, page[1] + s
	}
}

func (p *PageBuff) UpdatePageBuff(buffSize int) {
	buffLenth := len(p.buff)
	p.pageBuffSize = buffSize
	p.pages = buffLenth / buffSize
	p.extra = buffLenth % buffSize
}

func (p *PageBuff) ChangePageBuffSize(newSize int) {
	if newSize <= 0 {
		p.pageBuffSize = 0
		return
	}

	p.pageBuffSize = newSize
	p.SetPages(p.buffSize / p.pageBuffSize)
	p.extra = p.buffSize % p.pageBuffSize
	p.startPage[1] = (p.startPage[0] + (p.pageBuffSize - 1)) % p.buffSize
	p.pages = p.GetCurrPage()
	// TODO: continue
}

func (p *PageBuff) GetCurrPage() int {
	if p.currIndex >= p.startPage[0] {
		return (p.currIndex - p.startPage[0]) / p.pageBuffSize
	} else {
		return (p.currIndex + p.buffSize - p.startPage[0]) / p.pageBuffSize
	}
}

func (p *PageBuff) SetPages(pages int) {
	p.pages = pages
}

func IncODec(value int) (bool, int) {
	if value < 0 {
		return true, -value
	} else {
		return false, value
	}

}

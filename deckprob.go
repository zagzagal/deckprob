package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type deckToCheck struct {
	num  int // cards in deck
	draw int // cards to draw
	t    int // number of target cards

	nCount bool // find average till found?
	bios   bool // is runner Bios?
	numEmu int  // number of time to emulate
}

type results struct {
	numSuc int       // number of successes
	st     time.Time // start time
	count  int       // total of cards to find on failure
}

func main() {
	num := flag.Int("size", 45, "Starting Deck Size")
	draw := flag.Int("draw", 5, "Number of cards to draw")
	targ := flag.Int("targets", 1, "Number of target cards")
	numEmu := flag.Int("run-times", 10,
		"Number of times to run simulation")
	bios := flag.Bool("bios", false, "Simulate for Bios")
	nCount := flag.Bool("nCount", false, "Report Average till find")

	flag.Parse()

	p := deckToCheck{num: *num,
		draw:   *draw,
		t:      *targ,
		bios:   *bios,
		numEmu: *numEmu,
		nCount: *nCount,
	}

	r := process(p)
	output(p, r)

}

func process(p deckToCheck) (r results) {
	r.st = time.Now()
	var wg sync.WaitGroup
	r.numSuc = 0
	r.count = 0
	for i := 0; i < p.numEmu; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d := NewDeck(p.num)
			f := func() bool {
				d.Shuffle()
				if d.Check(p.draw, p.t) {
					r.numSuc++
					return true
				}
				return false
			}
			//fmt.Printf("%v\n%v\n", d.cards, d.Check(p.draw, p.t))
			if p.bios {
				d.Shuffle()
				if d.Check(6, p.t) {
					r.numSuc++
					return
				}
				d.Resize(41)
			}
			if f() {
				return
			}
			if f() {
				return
			}
			if p.nCount {
				r.count += d.Ncheck(p.t) - p.draw
			}
		}()
	}
	wg.Wait()
	return
}

func output(p deckToCheck, r results) {
	fmt.Printf("Results:\n")
	fmt.Println("------------------------------")
	fmt.Printf("Trys: %d\n", p.numEmu)
	fmt.Printf("draw size: %d\n", p.draw)
	fmt.Printf("number of target cards: %d\n", p.t)
	fmt.Printf("%v%%: %v of %v\n",
		(float64(r.numSuc)/float64(p.numEmu))*100,
		r.numSuc,
		p.numEmu)
	if p.nCount {
		fmt.Printf("Average draws to find on Failure: %v\n",
			(float64(r.count) / float64(p.numEmu-r.numSuc)))
		fmt.Printf("  %d of %d\n", r.count, p.numEmu-r.numSuc)
	}

	fmt.Printf("total time: %v\n", time.Now().Sub(r.st).String())

}

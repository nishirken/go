package main

import (
	"fmt"
	"time"
)

type Chopstick struct {
	index int
	isUsed bool
}

func (c *Chopstick) get() {
	fmt.Printf("Chopstick is used %d\n", c.index)
	c.isUsed = true
}

func (c *Chopstick) put() {
	fmt.Printf("Chopstick not used %d\n", c.index)
	c.isUsed = false
}

type Philosopher struct {
	index int
	isEating bool
	hasLeft bool
	hasRight bool
	eatTimes int
}

func (p *Philosopher) eat() {
	if !(p.hasRight && p.hasLeft) {
		panic("Can't eat without both chopsticks")
	}

	p.isEating = true
	p.eatTimes++
}

func (p *Philosopher) think() {
	p.isEating = false
}

func (p *Philosopher) getLeft() {
	p.hasLeft = true
}

func (p *Philosopher) getRight() {
	p.hasRight = true
}

func (p *Philosopher) putLeft() {
	p.hasLeft = false
}

func (p *Philosopher) putRight() {
	p.hasRight = false
}

func getChopsticks(p *Philosopher, cs *[]Chopstick) (*Chopstick, *Chopstick) {
	var right = p.index - 1
	if right < 0 {
		right = 4
	}
	return &(*cs)[p.index], &(*cs)[right]
}

func almostAllUsed(cs *[]Chopstick) bool {
	var counter int = 0
	for _, c := range *cs {
		if c.isUsed {
			counter += 1
		}
		if counter >= 4 {
			return true
		}
	}
	return false
}

func runTest(xs *[]Philosopher, eatTimes int) {
	ticker := time.NewTicker(50 * time.Millisecond)
	quit := make(chan struct{})

	for {
		select {
		case <- ticker.C:
			if (*xs)[0].isEating && (*xs)[1].isEating {
				panic("Cant eat 0 and 1 at the same time")
			}

			if (*xs)[0].isEating && (*xs)[4].isEating {
				panic("Cant eat 0 and 4 at the same time")
			}

			if (*xs)[1].isEating && (*xs)[2].isEating {
				panic("Cant eat 1 and 2 at the same time")
			}

			if (*xs)[2].isEating && (*xs)[3].isEating {
				panic("Cant eat 2 and 3 at the same time")
			}

			if (*xs)[3].isEating && (*xs)[4].isEating {
				panic("Cant eat 3 and 4 at the same time")
			}

			for _, p := range *xs {
				if p.eatTimes > eatTimes {
					panic("Cant eat more than 3 times")
				}
			}

			var allEats = true
			for _, p := range *xs {
				if p.eatTimes != eatTimes {
					allEats = false
					break
				}
			}

			if allEats {
				panic("All philosophers ate")
			}
		case <- quit:
			ticker.Stop()
			return
		}
	}
}

func eat(p *Philosopher, left *Chopstick, right *Chopstick) {
	fmt.Printf("Philosopher %d start eating with %d %d\n", p.index, left.index, right.index)
	p.eat()
	left.get()
	right.get()
	p.eatTimes += 1
}

func think(p *Philosopher, left *Chopstick, right *Chopstick) {
	fmt.Printf("Philosopher %d start thinking\n", p.index)
	p.think()
	left.put()
	right.put()
}

func tryAny(p *Philosopher, left *Chopstick, right *Chopstick, isRightFirst bool) {
	fmt.Printf("Try any for %d\n", p.index)
	var first *Chopstick = left
	var second *Chopstick = right

	if isRightFirst {
		first = right
		second = left
	}

	if first.isUsed {
		if second.isUsed {
			time.Sleep(100 * time.Millisecond)
			go tryAny(p, left, right, !isRightFirst)
		} else {
			second.get()
			if isRightFirst {
				p.getLeft()
			} else {
				p.getRight()
			}
		}
	} else {
		first.get()
		if isRightFirst {
			p.getRight()
		} else {
			p.getLeft()
		}
	}
}

func tryOne(p *Philosopher, left *Chopstick, right *Chopstick) {
	fmt.Printf("Try one for %d\n", p.index)
	if p.hasLeft {
		if !right.isUsed {
			right.get()
			p.getRight()
		} else {
			p.putLeft()
			left.put()
			go tryAny(p, left, right, false)
		}
	} else if p.hasRight {
		if !left.isUsed {
			left.get()
			p.getLeft()
		} else {
			p.putRight()
			right.put()
			go tryAny(p, left, right, false)
		}
	}
}

func servePhilosopher(philosopher *Philosopher, cs *[]Chopstick, eatTimes int) {
	fmt.Printf("Serve %d\n", philosopher.index)
	if philosopher.isEating || philosopher.eatTimes == eatTimes {
		return
	}
	left, right := getChopsticks(philosopher, cs)

	var denyForLast = almostAllUsed(cs)

	if denyForLast {
		fmt.Printf("Deny for %d\n", philosopher.index)
		time.Sleep(150 * time.Millisecond)
		go servePhilosopher(philosopher, cs, eatTimes)
		return
	}

	if !philosopher.hasLeft && !philosopher.hasRight {
		go tryAny(philosopher, left, right, true)
	} else if philosopher.hasLeft && philosopher.hasRight {
		eat(philosopher, left, right)
		time.Sleep(100 * time.Millisecond)
		think(philosopher, left, right)
		go servePhilosopher(philosopher, cs, eatTimes)
	} else {
		go tryOne(philosopher, left, right)
	}
}

func main() {
	const howMany = 5
	const eatTimes = 3
	philosophers := make([]Philosopher, howMany)
	chopsticks := make([]Chopstick, howMany)
	freeCs := make(chan *Chopstick, howMany)

	for i := 0; i < howMany; i++ {
		chopsticks[i].index = i
		freeCs <- &chopsticks[i]
	}

	for i := 0; i < howMany; i++ {
		(philosophers[i]).index = i
		go servePhilosopher(&philosophers[i], &chopsticks, eatTimes)
	}

	runTest(&philosophers, eatTimes)
}

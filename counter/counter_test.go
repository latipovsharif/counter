package counter_test

import (
	"sync"
	"testing"

	"github.com/latipovsharif/counter/counter"
)

const (
	threadCount        = 100
	iterationPerThread = 10000
)

func TestCounterIncrement_ThreadCheck(t *testing.T) {
	cnt := counter.Counter{}
	wg := sync.WaitGroup{}

	wg.Add(threadCount)

	for i := 0; i < threadCount; i++ {
		go incrementor(&cnt, &wg)
	}

	wg.Wait()

	expected := threadCount * iterationPerThread

	if cnt.Value() != uint(expected) {
		t.Fatalf("TestCounterIncrement_ThreadCheck failed got: %v, expected: %v", cnt.Value(), expected)
	}
}

func incrementor(c *counter.Counter, wg *sync.WaitGroup) {
	for i := 0; i < iterationPerThread; i++ {
		c.Increment()
	}

	wg.Done()
}

func TestCounterIncrement_OverflowCheck(t *testing.T) {
	cnt := counter.Counter{}

	cnt.SetMaximumValue(iterationPerThread)

	for i := uint(0); i <= iterationPerThread; i++ {
		cnt.Increment()
	}

	expected := 0
	if cnt.Value() != uint(expected) {
		t.Fatalf("TestCounterIncrement_OverflowCheck failed got: %v, expected: %v", cnt.Value(), 0)
	}
}

func TestCounterIncrement_MaxGreaterThanValue(t *testing.T) {
	cnt := counter.Counter{}

	for i := uint(0); i < iterationPerThread; i++ {
		cnt.Increment()
	}

	cnt.SetMaximumValue(iterationPerThread - 1)

	expected := 0
	if cnt.Value() != uint(expected) {
		t.Fatalf("TestCounterIncrement_MaxGreaterThanValue failed got: %v, expected: %v", cnt.Value(), 0)
	}
}

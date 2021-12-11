package main

import (
	"fmt"
	"strconv"

	"github.com/fkarakas/adventofcode_2021/utils"
)

type bit int

const (
	file = "data.txt"
)

var (
	ZERO bit = 0
	ONE  bit = 1
)

func NoFilter(b Bits) bool {
	return true
}

func FilterByIndex(index int, b bit) func(bs Bits) bool {
	return func(bs Bits) bool {
		return bs.get(index) == b
	}
}

func bits(filter func(b Bits) bool) chan interface{} {
	return utils.Data(file, func(v string) interface{} {
		bs := Bits{
			data: []rune(v),
		}
		if filter(bs) {
			return bs
		}
		return nil
	})
}

func bitsAsSlice(filter func(b Bits) bool) []Bits {
	result := []Bits{}

	for v := range utils.Data(file, func(v string) interface{} {
		bs := Bits{
			data: []rune(v),
		}
		if filter(bs) {
			return bs
		}
		return nil
	}) {
		result = append(result, v.(Bits))
	}

	return result
}

func bitsAsSliceFiltered(arr []Bits, filter func(b Bits) bool) []Bits {
	result := []Bits{}

	for _, v := range arr {
		if filter(v) {
			result = append(result, v)
		}
	}

	return result
}

type Bits struct {
	data []rune
	// less significant bit first
	lsb bool
}

func (b Bits) length() int {
	return len(b.data)
}

func (b Bits) get(pos int) bit {
	if pos < 0 || pos >= len(b.data) {
		return -1
	}
	if b.lsb {
		return bit(int(b.data[len(b.data)-1-pos] - '0'))
	}
	return bit(int(b.data[pos] - '0'))
}

func (b Bits) toInt() int64 {
	return binaryToInt(string(b.data))
}

func (b Bits) String() string {
	return string(b.data)
}

type Counter struct {
	oneBitCounter  int
	zeroBitCounter int
}

type BitCounter struct {
	counters []Counter
}

func (bc *BitCounter) Inc(index int, b bit) {
	if index >= len(bc.counters) {
		// nb element to add to the slice
		nb := (index - len(bc.counters)) + 1
		for i := 1; i <= nb; i++ {
			bc.counters = append(bc.counters, Counter{})
		}
	}
	ct := bc.counters[index]
	switch b {
	case ZERO:
		ct.zeroBitCounter++
	case ONE:
		ct.oneBitCounter++
	}
	bc.counters[index] = ct
}

func flip(b bit) bit {
	if b == ONE {
		return ZERO
	}
	return ONE
}

func (bc BitCounter) GetMostCommonBitByIndex(index int) bit {
	c := bc.counters[index]
	if c.oneBitCounter > c.zeroBitCounter || c.oneBitCounter == c.zeroBitCounter {
		return ONE
	}

	return ZERO
}

func (bc BitCounter) GetLeastCommonBitByIndex(index int) bit {
	return flip(bc.GetMostCommonBitByIndex(index))
}

func (bc BitCounter) GetMostAndLeastCommonBits() (most string, least string) {
	for _, c := range bc.counters {
		if c.oneBitCounter > c.zeroBitCounter {
			most += "1"
			least += "0"
		} else {
			most += "0"
			least += "1"
		}
	}
	return
}

func binaryToInt(binary string) int64 {
	res, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return -1
	}
	return res
}

// should be 198 (sample) or 3895776 (data)
func response1() int64 {
	bc := BitCounter{}

	for b := range bits(NoFilter) {
		bs := b.(Bits)
		for i := 0; i < bs.length(); i++ {
			bc.Inc(i, bs.get(i))
		}
	}
	m, l := bc.GetMostAndLeastCommonBits()
	return binaryToInt(m) * binaryToInt(l)
}

func findOxygenGeneratorRating(arr []Bits, index int) Bits {
	if len(arr) == 1 {
		return arr[0]
	}
	bc := BitCounter{}
	for _, bs := range arr {
		for i := 0; i < bs.length(); i++ {
			bc.Inc(i, bs.get(i))
		}
	}

	b := bc.GetMostCommonBitByIndex(index)
	arrf := bitsAsSliceFiltered(arr, FilterByIndex(index, b))
	index++
	return findOxygenGeneratorRating(arrf, index)
}

func findCO2ScrubberRating(arr []Bits, index int) Bits {
	if len(arr) == 1 {
		return arr[0]
	}
	bc := BitCounter{}
	for _, bs := range arr {
		for i := 0; i < bs.length(); i++ {
			bc.Inc(i, bs.get(i))
		}
	}

	b := bc.GetLeastCommonBitByIndex(index)
	arrf := bitsAsSliceFiltered(arr, FilterByIndex(index, b))
	index++
	return findCO2ScrubberRating(arrf, index)
}

// should be 7928162
func response2() int64 {
	ogr := findOxygenGeneratorRating(bitsAsSlice(NoFilter), 0)
	csr := findCO2ScrubberRating(bitsAsSlice(NoFilter), 0)
	return ogr.toInt() * csr.toInt()
}

func main() {
	fmt.Printf("response1=%v\n", response1())
	fmt.Printf("response2=%v\n", response2())
}

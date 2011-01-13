package gobash

import (
  "fmt"
  "utf8"
)

const (
  stringBuilderInitialCapacity = 10
)

type StringBuilder struct {
  runes []int
  cnt int
}

func NewStringBuilder() *StringBuilder {
  sb := new(StringBuilder)
  sb.runes = make([]int, stringBuilderInitialCapacity)
  return sb
}

func (sb *StringBuilder) Len() int {
  return sb.cnt
}

func (sb *StringBuilder) At(index int) int {
  if index < 0 {
    panic(fmt.Sprintf("StringBuilder.At(index=%d): index must be non-negative", index))
  }
  if index >= sb.cnt {
    panic(fmt.Sprintf("StringBuilder.At(index=%d): index out of range. Current length is %d", index, sb.cnt))
  }
  return sb.runes[index]
}

func (sb *StringBuilder) AtFirst() int {
  return sb.At(0)
}

func (sb *StringBuilder) AtLast() int {
  return sb.At(sb.Len() - 1)
}

func (sb *StringBuilder) Add(rune int) {
  sb.ensureCapacity(sb.cnt + 1)
  sb.runes[sb.cnt] = rune
  sb.cnt++
}

func (sb *StringBuilder) HasRune(rune int) bool {
  for _, v := range sb.runes {
    if v == rune {
      return true
    }
  }
  return false
}

func (sb *StringBuilder) Append(another *StringBuilder) {
  if another == nil {
    return
  }
  sb.ensureCapacity(sb.cnt + another.cnt)
  for i, v := range another.runes {
    sb.runes[sb.cnt + i] = v
  }
  sb.cnt += another.cnt
}

func (sb *StringBuilder) String() string {
	return runesToString(sb.runes[:sb.cnt])
}

func (sb *StringBuilder) RangeString(from, to int) string {
    if from < 0 {
      panic(fmt.Sprintf("StringBuilder.RangeString(from=%d, to=%d): from must be non negative", from, to))
    }
    if to < from {
      panic(fmt.Sprintf("StringBuilder.RangeString(from=%d, to=%d): from must be non greater than to", from, to))
    }
    if to >= sb.cnt {
      panic(fmt.Sprintf("StringBuilder.RangeString(from=%d, to=%d): to is out of bounds. Current length is %d", from, to, sb.cnt))
    }
    return runesToString(sb.runes[from : to])
}

func runesToString(runes []int) string {
    length := 0
	for _, v := range runes {
		length += utf8.RuneLen(v)
	}
	data := make([]byte, length)
	cur := data
	for _, v := range runes {
		rlen := utf8.RuneLen(v)
		utf8.EncodeRune(cur[0:rlen], v)
		cur = cur[rlen:]
	}
	return string(data)
}

func (sb *StringBuilder) ensureCapacity(capacity int) {
  if len(sb.runes) >= capacity {
    return
  }
  newlen := len(sb.runes) * 2
  if newlen < capacity {
    newlen = capacity
  }
  runes := make([]int, newlen)
  for i, v := range sb.runes {
    runes[i] = v
  }
  sb.runes = runes
}


package util

// Comparator is an interface used to compare values
type Comparator interface {
  // Compare two values of the Same Type
  Compare(a, b interface{}) int
  
  // Return the shortest data between a and b
  FindShortestSep(a, b interface{}) interface{}
  
  // Name of this comparator
  Name() string
}

type Compare func(a, b interface{}) int

// BinaryCompare is a function compare two bytes array
func BinaryCompare(fval, sval interface{}) int {
  first, second := fval.([]byte), sval.([]byte)
  var clen int
  if len(first) > len(second) {
    clen = len(second)
  } else {
    clen = len(first)
  }

  for i := 0; i < clen; i++ {
    switch true {
    case first[i] > second[i]:
      return 1;
    case first[i] < second[i]:
      return -1;
    }
  }

  if len(first) > len(second) {
    return 1
  } else if len(first) == len(second) {
    return 0
  } else {
    return -1
  }
}


type binaryCmp struct {
}

// BinaryComparator is an binary comparator instance
var BinaryComparator binaryCmp


func (binary binaryCmp) Compare(a, b interface{}) int {
  return BinaryCompare(a.([]byte), b.([]byte))
}

func (binary binaryCmp) Name() string {
  return "binary_comparator"
}

func (binary binaryCmp) FindShortestSep(val1, val2 interface{}) interface{} {
  var minLen int
  key1, key2 := val1.([]byte), val2.([]byte)
  if len(key1) > len(key2) {
    minLen = len(key2)
  } else {
    minLen = len(key1)
  }

  var pos int
  for pos = 0; pos < minLen; pos++ {
    if key1[pos] + 1 >= key2[pos] {
      continue
    }
  }

  if pos < minLen {
    key1[pos]++
    return key1[:pos + 1]
  }
  return key1
}

// Copyright 2019 Dmitry A. Mottl. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.

package circularqueue

import "fmt"

func ExampleLen() {
	q := NewQueue(10)

	q.Push("abc")
	q.Push("def")
	q.Push("ghi")

	fmt.Println(q.Len())
	// Output: 3
}

func ExampleCap() {
	q := NewQueue(10)

	q.Push("abc")
	q.Push("def")
	q.Push("ghi")

	fmt.Println(q.Cap())
	// Output: 10
}

func ExampleVacant() {
	q := NewQueue(10)

	q.Push("abc")
	q.Push("def")
	q.Push("ghi")

	fmt.Println(q.Vacant())
	// Output: 7
}

func ExamplePush() {
	q := NewQueue(2)

	vacant1, err1 := q.Push("abc")
	vacant2, err2 := q.Push("def")
	vacant3, err3 := q.Push("ghi")

	fmt.Println(vacant1, err1)
	fmt.Println(vacant2, err2)
	fmt.Println(vacant3, err3)
	// Output:
	// 1 <nil>
	// 0 <nil>
	// 0 Can't push: queue is full (capacity=2)
}

func ExamplePopAt() {
	q := NewQueue(10)
	q.Push("abc")
	q.Push("def")
	val1, err1 := q.PopAt(0)
	val2, err2 := q.PopAt(0)
	val3, err3 := q.PopAt(0)
	fmt.Println(val1, err1)
	fmt.Println(val2, err2)
	fmt.Println(val3, err3)
	// Output:
	// abc <nil>
	// def <nil>
	// <nil> Can't pop: queue is empty
}

func ExamplePop() {
	q := NewQueue(10)
	q.Push("abc")
	q.Push("def")
	val1, err1 := q.Pop()
	val2, err2 := q.Pop()
	val3, err3 := q.Pop()
	fmt.Println(val1, err1)
	fmt.Println(val2, err2)
	fmt.Println(val3, err3)
	// Output:
	// def <nil>
	// abc <nil>
	// <nil> Can't pop: queue is empty
}

func ExampleString() {
	q := NewQueue(10)
	q.Push("abc")
	q.Push("def")
	q.Push("ghi")
	q.PopAt(0)
	fmt.Println(q)
	// Output: Queue (len=2, cap=10) [ def, ghi ]
}

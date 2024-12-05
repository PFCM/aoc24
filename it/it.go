// package it holds some iterator helpers.
package it

import "iter"

// Zip zips two iter.Seq into an iter.Seq2 that yields paired up
// values from each iterator. Stops as soon as either iterator runs
// out of elements.
func Zip[A, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		nextA, stopA := iter.Pull(a)
		defer stopA()
		nextB, stopB := iter.Pull(b)
		defer stopB()
		for {
			aVal, aOK := nextA()
			if !aOK {
				return
			}
			bVal, bOK := nextB()
			if !bOK {
				return
			}
			if !yield(aVal, bVal) {
				return
			}
		}
	}
}

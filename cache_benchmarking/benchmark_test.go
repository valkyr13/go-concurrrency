package cachebenchmarking

import "testing"

//: go test -bench=. -race -cpu=1,2,4,8

func BenchmarkGetWithReadLocking(b *testing.B) {
	c := NewCache()
	c.Set("k", "v")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get_With_Read_Lock("k")
		}
	})
}

func BenchmarkGetWithWriteLocking(b *testing.B) {
	c := NewCache()
	c.Set("k", "v")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get_With_Write_Lock("k")
		}
	})
}

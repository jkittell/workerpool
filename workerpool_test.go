package workerpool

import (
	"testing"
)

func workerpool(numberOfJobs int) {
	in, out := Start[int](10)

	go func() {
		for i := 0; i < numberOfJobs; i++ {
			i := i
			in <- func() (int, error) {
				return i, nil
			}
		}
		close(in)
	}()

	// consume the produced numbers and sum them up
	sum := 0
	for o := range out {
		if o.Err != nil {
			panic(o.Err)
		}
		sum += o.Value
	}
}

func TestStart(t *testing.T) {
	workerpool(10)
}

func benchmarkStart(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		workerpool(i)
	}
}

func BenchmarkStart1(b *testing.B)   { benchmarkStart(1, b) }
func BenchmarkStart2(b *testing.B)   { benchmarkStart(2, b) }
func BenchmarkStart3(b *testing.B)   { benchmarkStart(3, b) }
func BenchmarkStart10(b *testing.B)  { benchmarkStart(10, b) }
func BenchmarkStart20(b *testing.B)  { benchmarkStart(20, b) }
func BenchmarkStart40(b *testing.B)  { benchmarkStart(40, b) }
func BenchmarkStart80(b *testing.B)  { benchmarkStart(80, b) }
func BenchmarkStart160(b *testing.B) { benchmarkStart(160, b) }

package ec

import "testing"

func BenchmarkDebug(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Debug("error msg")
	}
}

// func BenchmarkDebug1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Debug1("error msg")
// 	}
// }

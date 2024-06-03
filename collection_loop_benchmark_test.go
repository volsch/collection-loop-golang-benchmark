// Copyright (c) 2024, Volker Schmidt
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright
//    notice,this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS”
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package collection_loop_golang_benchmark

import (
	"fmt"
	"testing"
)

var loopIterations = [...]int{0, 1, 10, 100, 1000}

func BenchmarkLoop(b *testing.B) {
	for _, testLoopIterations := range loopIterations {
		b.Run(fmt.Sprintf("loop_iterations_%d", testLoopIterations), func(b *testing.B) {
			values := make([]int32, 0, testLoopIterations)
			for i := 0; i < testLoopIterations; i++ {
				values = append(values, int32(i))
			}

			var result []int32
			b.ResetTimer()
			for n := 0; n < b.N; n++ {

				length := len(values)
				result = make([]int32, 0, length)
				for i := 0; i < length; i++ {
					value := values[i]
					if value%2 == 0 {
						result = append(result, value)
					}
				}

			}
		})
	}
}

func BenchmarkRangeLoop(b *testing.B) {
	for _, testLoopIterations := range loopIterations {
		b.Run(fmt.Sprintf("loop_iterations_%d", testLoopIterations), func(b *testing.B) {
			values := make([]int32, 0, testLoopIterations)
			for i := 0; i < testLoopIterations; i++ {
				values = append(values, int32(i))
			}

			var result []int32
			b.ResetTimer()
			for n := 0; n < b.N; n++ {

				result = make([]int32, 0, len(values))
				for _, value := range values {
					if value%2 == 0 {
						result = append(result, value)
					}
				}

			}
		})
	}
}

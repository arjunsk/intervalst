// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/arjunsk/intervalst"
)

func main() {
	ist := intervalst.NewIntervalTree()
	ist.Insert(intervalst.NewInt64Interval(1, 3), "user1")
	ist.Insert(intervalst.NewInt64Interval(9, 13), "user2")
	ist.Insert(intervalst.NewInt64Interval(7, 20), "user3")

	// Output :
	// given interval {Begin:4 End:8} overlaps with these existing intervals
	// interval_st interval: &{Ivl:{Begin:7 End:20} Val:user3}
	findIntervalIntersections(ist, intervalst.NewInt64Interval(4, 8))

	// Output :
	// given point {Begin:10 End:11} overlaps with these existing intervals
	// interval_st interval: &{Ivl:{Begin:7 End:20} Val:user3}
	// interval_st interval: &{Ivl:{Begin:9 End:13} Val:user2}
	findPointIntersections(ist, intervalst.NewInt64Point(10))

	// Output : 3
	fmt.Println(ist.Len())

	// Output:
	// given interval {Begin:4 End:8} does not overlap with any existing intervals
	ist.Delete(intervalst.NewInt64Interval(7, 20))
	findIntervalIntersections(ist, intervalst.NewInt64Interval(4, 8))

}

func findPointIntersections(ist intervalst.IntervalTree, newPoint intervalst.Interval) {

	if ist.Intersects(newPoint) {
		fmt.Printf("given point %+v overlaps with these existing intervals \n", newPoint)

		rs := ist.Stab(newPoint)
		for _, v := range rs {
			fmt.Printf("interval_st interval: %+v\n", v)
		}
		fmt.Println()
	}
}

func findIntervalIntersections(ist intervalst.IntervalTree, newInterval intervalst.Interval) {
	if ist.Intersects(newInterval) {
		fmt.Printf("given interval %+v overlaps with these existing intervals \n", newInterval)
		rs := ist.Stab(newInterval)
		for _, v := range rs {
			fmt.Printf("interval_st interval: %+v\n", v)
		}
		fmt.Println()
	} else {
		fmt.Printf("given interval %+v does not overlap with any existing intervals \n", newInterval)
	}
}

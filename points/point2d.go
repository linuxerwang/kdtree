/*
 * Copyright 2018 Dennis Kuhnert
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package points

import "fmt"

// Point2D ...
type Point2D struct {
	X float64
	Y float64
}

// Dimensions ...
func (p *Point2D) Dimensions() int {
	return 2
}

// Dimension ...
func (p *Point2D) Dimension(i int) float64 {
	if i == 0 {
		return p.X
	}
	return p.Y
}

// String ...
func (p *Point2D) String() string {
	return fmt.Sprintf("{%.2f %.2f}", p.X, p.Y)
}

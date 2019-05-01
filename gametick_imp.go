// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gametick

import "time"

func (t GameTick) ToTimeDuration() time.Duration {
	return time.Duration(t)
}

func (t GameTick) AddDuration(dur time.Duration) GameTick {
	return t + GameTick(dur)
}

func FromTimeDurationToTickType(s time.Duration) GameTick {
	return GameTick(s)
}

func (t GameTick) AddIntSec(s int) GameTick {
	return t + GameTick(s)*Nano
}

func (t GameTick) AddInt64Sec(s int64) GameTick {
	return t + GameTick(s)*Nano
}

func (t GameTick) AddFloat64Sec(s float64) GameTick {
	return t + GameTick(s*float64(Nano))
}

func (t GameTick) ToInt64Sec() int64 {
	return int64(t / Nano)
}

func (t GameTick) ToInt64NanoSec() int64 {
	return int64(t)
}

func (t GameTick) ToFloat64Sec() float64 {
	return float64(t) / float64(Nano)
}

func (t GameTick) ToFloat64Day() float64 {
	return float64(t) / float64(Nano) / 3600 / 24
}

func (t GameTick) ToIntSec() int {
	return int(t / Nano)
}

func (t GameTick) MulFloat64(f float64) GameTick {
	return GameTick(float64(t) * f)
}

func FromInt64NanoSecToTickType(s int64) GameTick {
	return GameTick(s)
}

func FromInt64SecToTickType(s int64) GameTick {
	return GameTick(s) * Nano
}

func FromIntSecToTickType(s int) GameTick {
	return GameTick(s) * Nano
}

func FromFloat64SecToTickType(s float64) GameTick {
	return GameTick(s * float64(Nano))
}

func MakeIn(x GameTick, r1, r2 GameTick) GameTick {
	if x < r1 {
		x = r1
	}
	if x >= r2 {
		x = r2 - 1
	}
	return x
}

// no more used
// func TimeToInt64(s time.Time) int64 {
// 	return s.UnixNano()
// }

// func Int64ToTime(i int64) time.Time {
// 	return time.Unix(0, i).UTC()
// }

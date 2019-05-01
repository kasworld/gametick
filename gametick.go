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

// 게임 컨텐츠용 시간, 게임 서버가 실행되어 있는 총 시간의 합
package gametick

import (
	"fmt"
	"strconv"
)

type GameTick int64

const Nano = GameTick(1000000000)
const GameTick_None = GameTick(0)

// no more need , no convert to humantime
// const GameTick_None = GameTick(math.MinInt64 / 2)

func (t GameTick) IsNone() bool {
	return t == GameTick_None
}

func (t GameTick) IsNotNone() bool {
	return t != GameTick_None
}

func (t GameTick) String() string {
	ns := int64(t % Nano)
	sec := int64(t / Nano)
	min := sec / 60
	sec -= min * 60
	hour := min / 60
	min -= hour * 60
	days := hour / 24
	hour -= days * 24

	return fmt.Sprintf(
		"GameTick[%d / %v %v:%v:%v.%09d]",
		t, days, hour, min, sec, ns)
}

func (t GameTick) ToIntString() string {
	return fmt.Sprintf("%v", int(t))
}

func FromIntString(s string) (GameTick, error) {
	int64v, err := strconv.ParseInt(s, 0, 64)
	return GameTick(int64v), err
}

func (t GameTick) ToDBInt64() int64 {
	return int64(t)
}

func FromDBInt64(v int64) GameTick {
	return GameTick(v)
}

// Copyright 2022 Akiomi Kamakura
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

package twitter

import (
	"bytes"
	"time"
)

type RubyDate time.Time

func (t *RubyDate) UnmarshalJSON(buf []byte) error {
	s := bytes.Trim(buf, `"`)
	parsed, err := time.Parse(time.RubyDate, string(s))
	if err != nil {
		return err
	}

	*t = RubyDate(parsed)
	return nil
}

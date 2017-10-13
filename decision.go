// Copyright (c) 2017 ircbot authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/sorcix/irc"
)

type Decision struct {
}

func NewDecision() (*Decision, error) {
	return &Decision{}, nil
}

func (dec *Decision) HandleMessage(conn *Conn, m *irc.Message) {
	msg := AcceptPRIVMSG(m)
	if msg == nil || msg.channel == "" {
		return
	}

	nickHandler := fmt.Sprintf("%s:", *nickname)
	if strings.HasPrefix(m.Trailing, nickHandler) {
		separator := ""

		// for historical reason suports both or and ou (pt) separators
		for _, s := range []string{" or ", " ou "} {
			if strings.Contains(m.Trailing, s) {
				separator = s
			}
		}

		if separator == "" {
			return
		}

		text := strings.Replace(m.Trailing, fmt.Sprintf("%s:", *nickname), "", 1)
		text = strings.TrimSuffix(text, "?")

		opts := strings.Split(text, separator)
		r := rand.Int()
		rep := opts[1]
		if r%2 == 0 {
			rep = opts[0]
		}

		say(conn, msg.channel, fmt.Sprintf("%s, of course!\n", rep))
	}
}

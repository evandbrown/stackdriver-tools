/*
 * Copyright 2019 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package session

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

// Emitter adds spinner logs to Stackdriver.
type Emitter interface {

	// Emit adds logs to Stackdriver which can be identified by the needle.
	Emit(needle string) (int, error)
}

// Probe searches for spinner logs in Stackdriver.
type Probe interface {

	// Find searches Stackdriver for logs with the given needle.
	Find(start time.Time, needle string, count int) (int, error)
}

// Session represents an instance of the Spinner.
type Session struct {
	emitter Emitter
	probe   Probe
}

// Result is a summary of the log loss from a Run of a Session.
type Result struct {
	GUID  string
	Found int
	Loss  float64
}

// NewSession constructs a Session.
func NewSession(emitter Emitter, probe Probe) Session {
	return Session{emitter, probe}
}

// Run tests the Stackdriver Nozzle by Emitting logs and Probing for them after a given interval.
func (s Session) Run(burstInterval time.Duration) (Result, error) {
	needle := generateNeedle()
	emitted, err := s.emitter.Emit(needle)
	if err != nil {
		return Result{}, err
	}

	queryTime := time.Now().Add(-burstInterval - 10)
	time.Sleep(burstInterval)

	found, err := s.probe.Find(queryTime, needle, emitted)
	if err != nil {
		return Result{}, err
	}
	return Result{needle, found, float64(emitted-found) / float64(emitted)}, nil
}

func generateNeedle() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic("failed to generate a needle")
	}
	needle := fmt.Sprintf("%x", uuid)
	return needle
}

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// generates a uid suitable for use as a container name
func uid() (string, error) {
	utcSeconds := time.Now().UTC().Unix()
	secondBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(secondBytes, uint64(utcSeconds))
	// encode to hex and drop leading zeros
	encodedSeconds := strings.TrimLeft(hex.EncodeToString(secondBytes), "0")
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", encodedSeconds, hex.EncodeToString(randomBytes)), nil
}

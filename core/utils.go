/*
 * Copyright 2021-2022 JetBrains s.r.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"os"
	"strings"
)

// lower a shortcut to strings.ToLower.
func lower(s string) string {
	return strings.ToLower(s)
}

// contains checks if a string is in a given slice.
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Append appends a string to a slice if it's not already there.
//
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func Append(slice []string, elems ...string) []string {
	if !contains(slice, elems[0]) {
		slice = append(slice, elems[0])
	}
	return slice
}

// CheckDirFiles checks if a directory contains files.
func CheckDirFiles(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return len(files) > 0
}

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: Evan Spendlove (@evanSpendlove)

package hermes

// FileOperation is an int used as part of the FileOperation enum within the intent log of Hermes' StateJournal.
type FileOperation int

// FileOperation enum is used for marking the operation intent within the StateJournal of Hermes.
// FileOperation has two possible values: CREATE, DELETE.
const (
	_                    = iota
	CREATE FileOperation = iota
	DELETE
)

// Intent stores the next intended file operation of Hermes.
// This is used as part of the StateJournal of Hermes.
type Intent struct {
	operation FileOperation // Stores the file operation intent, either CREATE or DELETE
	filename  string        // Stores the filename that the operation is being performed on.
}

// StateJournal stores the state of Hermes in two parts: the next operation intent and a map of filenames.
// The intent stores the operation to be performed and the name of the file that the operation is being performed on.
// The filenames map is a map of file IDs to filenames.
// If an entry does not exist for a given ID, then the file does not exist (i.e. has been deleted)
type StateJournal struct {
	intent    Intent         // intent stores the next intended file operation and the name of the file that the operation is being performed on.
	filenames map[int]string // filenames is a map of file IDs to filenames.
}
// Copyright 2013 Robert A. Uhl.  All rights reserved.
//
// This file is part of cypherback.
//
// Cypherback is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Cypherback is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Cypherback.  If not, see <http://www.gnu.org/licenses/>.

package cypherback

import (
	//"fmt"
	memoryBackend "cypherback/backends/memory"
	"os"
	"testing"
)

func TestDirectoryInfo(t *testing.T) {
	backupset := BackupSet{}
	info, err := os.Lstat(".")
	if err != nil {
		t.Fatal(err)
	}
	rec, err := backupset.fileRecordFromFileInfo(".", info)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rec)
}

func TestFileInfo(t *testing.T) {
	secrets, err := generateSecrets()
	defer ZeroSecrets(secrets)
	if err != nil {
		t.Fatal(err)
	}
	backend := memoryBackend.New()
	backupSet, err := EnsureBackupSet(backend, secrets, "foo")
	if err != nil {
		t.Fatal(err)
	}
	ProcessPath(backupSet, ".")
}

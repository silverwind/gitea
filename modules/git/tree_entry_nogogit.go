// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build all !gogit

package git

import (
	"strconv"
	"strings"
)

// TreeEntry the leaf in the git tree
type TreeEntry struct {
	ID SHA1

	ptree *Tree

	entryMode EntryMode
	name      string

	size     int64
	sized    bool
	fullName string
}

// Name returns the name of the entry
func (te *TreeEntry) Name() string {
	if te.fullName != "" {
		return te.fullName
	}
	return te.name
}

// Mode returns the mode of the entry
func (te *TreeEntry) Mode() EntryMode {
	return te.entryMode
}

// Size returns the size of the entry
func (te *TreeEntry) Size() int64 {
	if te.IsDir() {
		return 0
	} else if te.sized {
		return te.size
	}

	stdout, err := NewCommand("cat-file", "-s", te.ID.String()).RunInDir(te.ptree.repo.Path)
	if err != nil {
		return 0
	}

	te.sized = true
	te.size, _ = strconv.ParseInt(strings.TrimSpace(stdout), 10, 64)
	return te.size
}

// IsSubModule if the entry is a sub module
func (te *TreeEntry) IsSubModule() bool {
	return te.entryMode == EntryModeCommit
}

// IsDir if the entry is a sub dir
func (te *TreeEntry) IsDir() bool {
	return te.entryMode == EntryModeTree
}

// IsLink if the entry is a symlink
func (te *TreeEntry) IsLink() bool {
	return te.entryMode == EntryModeSymlink
}

// IsRegular if the entry is a regular file
func (te *TreeEntry) IsRegular() bool {
	return te.entryMode == EntryModeBlob
}

// IsExecutable if the entry is an executable file (not necessarily binary)
func (te *TreeEntry) IsExecutable() bool {
	return te.entryMode == EntryModeExec
}

// Blob returns the blob object the entry
func (te *TreeEntry) Blob() *Blob {
	return &Blob{
		ID:       te.ID,
		repoPath: te.ptree.repo.Path,
		name:     te.Name(),
		size:     te.size,
		gotSize:  te.sized,
	}
}

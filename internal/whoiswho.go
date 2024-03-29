// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"os"

	"github.com/mitchellh/go-ps"
)

// Executable name running this process. This is not a path to the
// executable.
func RunningExeName() string {
	proc, _ := ps.FindProcess(os.Getpid())
	return proc.Executable()
}

func ParentExeName() string {
	proc, _ := ps.FindProcess(os.Getppid())
	return proc.Executable()
}

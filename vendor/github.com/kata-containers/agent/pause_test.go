//
// Copyright (c) 2017 HyperHQ Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForkPauseBin(t *testing.T) {
	skipUnlessRoot(t)

	cmd := &exec.Cmd{
		Path: selfBinPath,
		Env:  []string{fmt.Sprintf("%s=%s", pauseBinKey, pauseBinValue)},
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}

	err := cmd.Start()
	assert.Nil(t, err, "Failed to fork pause binary: %s", err)

	_, err = os.Stat(fmt.Sprintf("/proc/%d/ns/pid", cmd.Process.Pid))
	assert.Nil(t, err, "Failed to stat pidns of pid %d: %s", cmd.Process.Pid, err)

	err = cmd.Process.Kill()
	assert.Nil(t, err, "Failed to kill pause binary: %s", err)

	_, err = cmd.Process.Wait()
	assert.Nil(t, err, "Failed to wait killed pause binary: %s", err)
}

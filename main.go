// A generated module for DaggerTools functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"dagger/dagger-tools/internal/dagger"
	"fmt"
)

type DaggerTools struct{}

func (m *DaggerTools) Update(version string) (*dagger.File) {
    container := dag.Container().
        From("alpine:latest").
        WithExec([]string{"wget", fmt.Sprintf("https://github.com/dagger/dagger/releases/download/%s/dagger_%s_linux_amd64.tar.gz", version, version)}).
		WithExec([]string{"tar", "xvf", fmt.Sprintf("dagger_%s_linux_amd64.tar.gz", version)})
	
    return container.File("dagger")
}
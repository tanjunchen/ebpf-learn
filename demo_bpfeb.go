// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadDemo returns the embedded CollectionSpec for demo.
func loadDemo() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_DemoBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load demo: %w", err)
	}

	return spec, err
}

// loadDemoObjects loads demo and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*demoObjects
//	*demoPrograms
//	*demoMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadDemoObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadDemo()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// demoSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type demoSpecs struct {
	demoProgramSpecs
	demoMapSpecs
}

// demoSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type demoProgramSpecs struct {
	KprobeExecve *ebpf.ProgramSpec `ebpf:"kprobe_execve"`
}

// demoMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type demoMapSpecs struct {
	KprobeMap *ebpf.MapSpec `ebpf:"kprobe_map"`
}

// demoObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadDemoObjects or ebpf.CollectionSpec.LoadAndAssign.
type demoObjects struct {
	demoPrograms
	demoMaps
}

func (o *demoObjects) Close() error {
	return _DemoClose(
		&o.demoPrograms,
		&o.demoMaps,
	)
}

// demoMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadDemoObjects or ebpf.CollectionSpec.LoadAndAssign.
type demoMaps struct {
	KprobeMap *ebpf.Map `ebpf:"kprobe_map"`
}

func (m *demoMaps) Close() error {
	return _DemoClose(
		m.KprobeMap,
	)
}

// demoPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadDemoObjects or ebpf.CollectionSpec.LoadAndAssign.
type demoPrograms struct {
	KprobeExecve *ebpf.Program `ebpf:"kprobe_execve"`
}

func (p *demoPrograms) Close() error {
	return _DemoClose(
		p.KprobeExecve,
	)
}

func _DemoClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed demo_bpfeb.o
var _DemoBytes []byte

package validate

import (
	"syscall"
	"testing"

	"github.com/docker/libcontainer/configs"
)

func containProc(values []*configs.Mount) bool {
	for _, m := range values {
		if m.Source == "proc" &&
			m.Destination == "/proc" &&
			m.Flags == syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV {
			return true
		}
	}
	return false
}

func TestAddProcMount(t *testing.T) {
	v := &ConfigValidator{}

	config := &configs.Config{
		Namespaces: configs.Namespaces{{Type: configs.NEWNS}},
	}
	v.addProcMount(config)
	if len(config.Mounts) != 1 && !containProc(config.Mounts) {
		t.Fatalf("addProc failed to add an proc mount")
	}

	config = &configs.Config{
		Namespaces: configs.Namespaces{{Type: configs.NEWPID}},
	}
	v.addProcMount(config)
	if len(config.Mounts) != 0 {
		t.Fatalf("mounts should have 0 items but reports %d", len(config.Mounts))
	}
}

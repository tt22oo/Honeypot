package session

import (
	"honeypot/core/config"

	"github.com/tt22oo/fakeshell/loader"
	"github.com/tt22oo/fakeshell/shell"
)

func initShell(cfgs *config.Config) (*shell.Shell, error) {
	shellCfg := &shell.Config{
		HomePath: cfgs.Shell.HomePath,
		DirPath:  cfgs.Shell.DirPath,
		Proc: &loader.ProcConfig{
			ProcessPath: cfgs.Shell.ProcessPath,
			CpuInfoPath: cfgs.Shell.CpuinfoPath,
			MemInfoPath: cfgs.Shell.MeminfoPath,
			VersionPath: cfgs.Shell.VersionPath,
		},
	}

	return shellCfg.New()
}

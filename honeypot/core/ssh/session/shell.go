package session

import (
	"github.com/tt22oo/fakeshell/loader"
	"github.com/tt22oo/fakeshell/shell"
)

func (cfgs *Config) initShell() (*shell.Shell, error) {
	shellCfg := &shell.Config{
		HomePath: cfgs.Config.Shell.HomePath,
		DirPath:  cfgs.Config.Shell.DirPath,
		Proc: &loader.ProcConfig{
			ProcessPath: cfgs.Config.Shell.ProcessPath,
			CpuInfoPath: cfgs.Config.Shell.CpuinfoPath,
			MemInfoPath: cfgs.Config.Shell.MeminfoPath,
			VersionPath: cfgs.Config.Shell.VersionPath,
		},
	}

	return shellCfg.New()
}

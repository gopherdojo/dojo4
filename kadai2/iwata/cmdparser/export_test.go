package cmdparser

func NewCmdConfig(dir, fromExt, toExt string) *CmdConfig {
	return &CmdConfig{dir, fromExt, toExt}
}

package client

type MegaCLI struct {
	c *Client
}

// exec executes an ExecFunc using 'megadl'.
func (ytdl *MegaCLI) exec(args ...string) ([]byte, error) {
	return ytdl.c.exec("megadl", args...)
}

// `megadl` = specify --path for output
func (ytdl *MegaCLI) DownloadWithDirName(dir, url string) error {
	cmds := []string{"--path", dir, url}
	_, err := ytdl.exec(cmds...)
	return err
}

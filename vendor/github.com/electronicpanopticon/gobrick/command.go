package gobrick

/// ExecCommand is a struct used to hold the command values passed to exec.Command()
type ExecCommand struct {
	Name string
	Args []string
}

package run

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nathanburkett/diff-chart/algorithm"
	"github.com/nathanburkett/diff-chart/input"
	"github.com/nathanburkett/diff-chart/output"
	"github.com/nathanburkett/diff-chart/transform"
	"os/exec"
)

// ErrRuntime indicates an error only detectable at runtime
var ErrRuntime = errors.New("runtime error")

// CmdFlags possible command invocation flag values
type CmdFlags struct {
	DownstreamRef string
	UpstreamRef   string
	InputType     string
	ReducerType   string
	OutputType    string
	SortType      string
}

// App managing struct which runs entire invocation process end-to-end
type App struct {
	Flags   *CmdFlags
	Reader  input.DiffReader
	Reducer transform.Reducer
	Writer  output.Writer
	Sorter  algorithm.Sorter
}

// NewApp factory func for run.App
func NewApp() *App {
	return &App{
		Flags: &CmdFlags{},
	}
}

// Attach attach instances based upon flags
func (a *App) Attach() error {
	read, err := input.Make(a.Flags.InputType)
	if err != nil {
		return err
	}

	a.Reader = read

	reduce, err := transform.Make(a.Flags.ReducerType)
	if err != nil {
		return err
	}

	a.Reducer = reduce

	w, err := output.Make(a.Flags.OutputType)
	if err != nil {
		return err
	}

	a.Writer = w

	s, err := algorithm.Make(a.Flags.SortType)
	if err != nil {
		return err
	}

	a. Sorter = s

	return nil
}

// Run run end-to-end process of invocation
func (a *App) Run() error {
	out, err := getCliDiffBytes(a.Flags.UpstreamRef, a.Flags.DownstreamRef)
	if err != nil {
		return fmt.Errorf("%s: %s", ErrRuntime, err)
	}

	diff, err := input.Read(a.Reader, bytes.NewBuffer(out))
	if err != nil {
		return fmt.Errorf("%s: %s", ErrRuntime, err)
	}

	diff, err = transform.Reduce(a.Reducer, diff)
	if err != nil {
		return fmt.Errorf("%s: %s", ErrRuntime, err)
	}

	diff, err = algorithm.Sort(a.Sorter, diff)
	if err != nil {
		return fmt.Errorf("%s: %s", ErrRuntime, err)
	}

	if err := a.Writer.Write(diff); err != nil {
		return fmt.Errorf("%s: %s", ErrRuntime, err)
	}

	return nil
}

func getCliDiffBytes(down, up string) ([]byte, error) {
	var b []byte

	dwnHsh, err := getOutputGitPointerHash(down)
	if err != nil {
		fmt.Println(string(dwnHsh))
		fmt.Println(err)
		return b, err
	}

	upHsh, err := getOutputGitPointerHash(up)
	if err != nil {
		fmt.Println(string(upHsh))
		fmt.Println(err)
		return b, err
	}

	cmdArgs := []string{
		"diff",
		fmt.Sprintf("%s..%s", upHsh, dwnHsh),
		"--numstat",
	}
	return runGitCmd(cmdArgs)
}

func getOutputGitPointerHash(ptr string) ([]byte, error) {
	args := []string{
		"rev-parse",
		"--short",
		ptr,
	}

	return runGitCmd(args)
}

func runGitCmd(args []string) ([]byte, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return bytes.Trim(out, "\n"), nil
}

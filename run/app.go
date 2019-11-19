package run

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nathanburkett/diff_table/algorithm"
	"github.com/nathanburkett/diff_table/input"
	"github.com/nathanburkett/diff_table/output"
	"github.com/nathanburkett/diff_table/transform"
	"os/exec"
)

var ErrRuntime = errors.New("runtime error")

type CmdArgs struct {
	DownstreamRef string
	UpstreamRef   string
	InputType     string
	ReducerType   string
	OutputType    string
	SortType      string
}

type App struct {
	Args    *CmdArgs
	Reader  input.DiffReader
	Reducer transform.Reducer
	Writer  output.Writer
	Sorter  algorithm.Sorter
}

func NewApp() *App {
	return &App{
		Args: &CmdArgs{},
	}
}

func (a *App) Attach() error {
	read, err := input.Make(a.Args.InputType)
	if err != nil {
		return err
	}

	a.Reader = read

	reduce, err := transform.Make(a.Args.ReducerType)
	if err != nil {
		return err
	}

	a.Reducer = reduce

	w, err := output.Make(a.Args.OutputType)
	if err != nil {
		return err
	}

	a.Writer = w

	s, err := algorithm.Make(a.Args.SortType)
	if err != nil {
		return err
	}

	a. Sorter = s

	return nil
}

func (a *App) Run() error {
	out, err := getCliDiffBytes(a.Args.UpstreamRef, a.Args.DownstreamRef)
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrRuntime, err)
	}

	diff, err := input.Read(a.Reader, bytes.NewBuffer(out))
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrRuntime, err)
	}

	diff, err = transform.Reduce(a.Reducer, diff)
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrRuntime, err)
	}

	diff, err = algorithm.Sort(a.Sorter, diff)
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrRuntime, err)
	}

	if err := a.Writer.Write(diff); err != nil {
		return fmt.Errorf("%s: %s\n", ErrRuntime, err)
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

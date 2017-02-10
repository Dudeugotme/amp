package cli

import (
	"errors"
	"regexp"
	"testing"
	"time"

	gexpect "github.com/Thomasrooney/gexpect"
)

// SuiteSpec defines the suite fields and the suite data structures.
type SuiteSpec struct {
	Name        string   `yaml:"name"`
	Timeout     string   `yaml:"timeout"`
	LookupDir   string   `yaml:"lookupdir"`
	SetupDir    string   `yaml:"setupdir"`
	TestDirs    []string `yaml:"testdirs"`
	TearDownDir string   `yaml:"teardowndir"`
	Setup       []TestSpec
	Tests       []TestSpec
	TearDown    []TestSpec
}

// TestSpec defines the test fields and test data structures.
type TestSpec struct {
	Name        string        `yaml:"name"`
	TestTimeout string        `yaml:"test-timeout"`
	CmdTimeout  string        `yaml:"cmd-timeout"`
	Commands    []CommandSpec `yaml:"commands"`
	Cache       map[string]string
}

// CommandSpec defines the command fields and data types.
type CommandSpec struct {
	Cmd         string   `yaml:"cmd"`
	Args        []string `yaml:"args"`
	Options     []string `yaml:"options"`
	Input       []string `yaml:"input"`
	Expectation string   `yaml:"expectation"`
	Skip        bool     `yaml:"skip"`
	Retry       int      `yaml:"retry"`
	Timeout     string   `yaml:"timeout"`
	Delay       string   `yaml:"delay"`
}

var (
	// Array of suite directories to run.
	suiteDir = []string{"./suite"}

	// Array of suites read
	suites = []*SuiteSpec{}

	// Map of all lookup regexes.
	regexMap map[string]string
)

// TestCli is the primary test func, tests all cli commands.
func TestCli(t *testing.T) {
	var err error
	// Parse all suites specified in the array of suite directories.
	suites, err = createSuite(suiteDir)
	if err != nil {
		t.Errorf("Unable to parse suites, reason: %v", err)
	}

	// Create all suites timeouts.
	cancels := []func(){}
	for _, suite := range suites {
		// Create suite duration.
		duration, err := time.ParseDuration(suite.Timeout)
		if err != nil {
			t.Fatal("Unable to create duration for timeout:", suite.Name, "Error:", err)
		}
		// Create suite timeout.
		cancel := createTimeout(t, duration, suite.Name)
		cancels = append(cancels, cancel)
	}

	// Run setup tests.
	t.Run("Setup", testSetup)

	// Run main tests in parallel.
	t.Run("Cmds", testCmds)

	// Run teardown tests.
	t.Run("TearDown", testTearDown)

	// End of suite timeouts.
	for _, cancel := range cancels {
		cancel()
	}
}

// testSetup is the test func for the setup commands.
func testSetup(t *testing.T) {
	// Run setup files from each suite.
	for _, suite := range suites {
		for _, setup := range suite.Setup {
			runTestSpec(t, setup)
		}
	}
}

// testCmds is the test func for the main commands.
func testCmds(t *testing.T) {
	// Run test files from each suite in parallel.
	for _, suite := range suites {
		for _, test := range suite.Tests {
			test := test
			t.Run(test.Name, func(t *testing.T) {
				t.Parallel()
				runTestSpec(t, test)
			})
		}
	}
}

// testTearDown is the test func for the teardown commands.
func testTearDown(t *testing.T) {
	// Run teardown files from each suite.
	for _, suite := range suites {
		for _, teardown := range suite.TearDown {
			runTestSpec(t, teardown)
		}
	}
}

// runTestspec creates the testspec timeout and executes each of its commands.
func runTestSpec(t *testing.T, testSpec TestSpec) {
	// Create testspec duration.
	duration, err := time.ParseDuration(testSpec.TestTimeout)
	if err != nil {
		t.Fatal("Unable to create duration for timeout:", testSpec.Name, "Error:", err)
	}
	// Create testspec timeout.
	cancel := createTimeout(t, duration, testSpec.Name)
	defer cancel()

	// Run all commands in testspec.
	for _, cmdSpec := range testSpec.Commands {
		runCmdSpec(t, cmdSpec, testSpec.Cache)
	}
}

// runCmdSpec creates the CommandSpec timeout and executes the command.
func runCmdSpec(t *testing.T, cmdSpec CommandSpec, cache map[string]string) {
	// Retry counter.
	i := 0

	// Generates the command string.
	cmd := genCmdStr(cmdSpec)

	// Create commandspec duration.
	duration, err := time.ParseDuration(cmdSpec.Timeout)
	if err != nil {
		t.Fatal("Unable to create duration for timeout:", cmd, "Error:", err)
	}
	// Create commandspec timeout.
	cancel := createTimeout(t, duration, cmd)
	defer cancel()

	// Template command string.
	cmd, err = templating(cmd, cache)
	if err != nil {
		t.Fatal("Executing templating failed:", cmd, "Error:", err)
	}

	// Check if expectation has corresponding regex
	if cmdSpec.Expectation != "" && regexMap[cmdSpec.Expectation] == "" {
		t.Fatal("Unable to fetch regex for command:", cmd, "reason: no regex for given expectation:", cmdSpec.Expectation)
	}

	// Template regex.
	regex, err := templating(regexMap[cmdSpec.Expectation], cache)
	if err != nil {
		t.Fatal("Executing templating failed:", cmdSpec.Expectation, "Error:", err)
	}

	// Retry loop.
	for i = 0; i <= cmdSpec.Retry; i++ {
		// err is set to nil to ensure no carryover.
		err = nil

		// Execute command.
		output := runCmd(cmd, cmdSpec)

		// Check ouptput against corresponding regex.
		expectedOutput := regexp.MustCompile(regex)
		if !expectedOutput.MatchString(string(output)) {
			err = errors.New("Error: mismatched return: command, " + cmd + ". regex, " + regex + ". ouput, " + output)
		}

		// Delay command execution if delay is set.
		if cmdSpec.Delay != "" {
			del, err := time.ParseDuration(cmdSpec.Delay)
			if err != nil {
				t.Fatal("Invalid delay specified: ", cmdSpec.Delay, "Error:", err)
			}
			time.Sleep(del)
		}

		// Given no error, stop executing.
		if err == nil {
			break
		}
	}
	// If the command retried, output retry counter.
	if i > 1 {
		t.Log("The command :", cmd, "has re-run", i, "times.")
	}
	// If the command still failed after retrying, fail test.
	if err != nil {
		t.Fatal(err)
	}
}

// runCmd executes the command, passing any inputs to it.
func runCmd(cmd string, spec CommandSpec) string {
	// Spawns new execution of command.
	child, err := gexpect.Spawn(cmd)
	if err != nil {
		panic(err)
	}

	// Sends each input to command.
	for _, input := range spec.Input {
		child.Send(input + "\r")
		child.Expect(input)
	}

	// Reads ouput until return on cli.
	output, _ := child.ReadUntil('$')

	return string(output)
}

// genCmdStr creates the command by concatenating the cmd, args and options fields.
func genCmdStr(cmdSpec CommandSpec) string {
	// Get command string.
	cmdStr := cmdSpec.Cmd

	// Add arguments.
	for _, arg := range cmdSpec.Args {
		cmdStr = cmdStr + " " + arg
	}

	// Add options.
	for _, opts := range cmdSpec.Options {
		cmdStr = cmdStr + " " + opts
	}

	return cmdStr
}

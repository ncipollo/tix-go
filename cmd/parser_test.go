package cmd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"tix/logger"
)

const TixVersion = "1.0"

func TestMain(m *testing.M) {
	logger.SetLogLevel(logger.LogLevelNormal)
	os.Exit(m.Run())
}

func TestParser_Parse_QuietLogLevel(t *testing.T) {
	parser := setupParser("-quiet", "foo.md")
	parser.Parse()
	assert.Equal(t, logger.LogLevelQuiet, logger.CurrentLogLevel())
}

func TestParser_Parse_QuietLogLevel_Shorthand(t *testing.T) {
	parser := setupParser("-q", "foo.md")
	parser.Parse()
	assert.Equal(t, logger.LogLevelQuiet, logger.CurrentLogLevel())
}

func TestParser_Parse_TixCommand(t *testing.T) {
	parser := setupParser("foo.md")
	cmd := parser.Parse()
	assert.IsType(t, &TixCommand{}, cmd)
}

func TestParser_Parse_VerboseLogLevel(t *testing.T) {
	parser := setupParser("-verbose", "foo.md")
	parser.Parse()
	assert.Equal(t, logger.LogLevelVerbose, logger.CurrentLogLevel())
}

func TestParser_Parse_VerboseLogLevel_Shorthand(t *testing.T) {
	parser := setupParser("-v", "foo.md")
	parser.Parse()
	assert.Equal(t, logger.LogLevelVerbose, logger.CurrentLogLevel())
}

func TestParser_Parse_VersionCommand(t *testing.T) {
	parser := setupParser("--version")
	cmd := parser.Parse()
	assert.IsType(t, &VersionCommand{}, cmd)
}

func setupParser(args ...string) *Parser {
	os.Args = append([]string{"fnew"}, args...)
	return NewParser([]string{}, TixVersion)
}

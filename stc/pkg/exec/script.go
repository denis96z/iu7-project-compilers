package exec

import (
	"strings"
)

type Script struct {
	lineIdx  int
	objects  map[string]Object
	commands []*Command
	stack    *Stack
}

func NewScript() *Script {
	return &Script{
		objects:  makeCoreObjects(),
		commands: make([]*Command, 0),
		stack:    NewStack(),
	}
}

func (s *Script) Execute(src string) {
	lines := strings.Split(src, "\n")
	for _, line := range lines {
		s.lineIdx++

		pLine := s.parseLine(line)
		if pLine == nil {
			continue
		}

		s.executeLine(pLine)
	}
}

type ParsedLine struct {
	Command string
	Params  []string
}

func (s *Script) parseLine(line string) *ParsedLine {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil
	}

	pLine := &ParsedLine{
		Params: make([]string, 0),
	}

	spcIdx := strings.Index(line, " ")
	if spcIdx <= 0 {
		s.logWrongFormat("no command")
	}

	pLine.Command = line[:spcIdx]
	line = strings.TrimSpace(line[spcIdx:])

	pArr := strings.Split(line, ",")
	for _, p := range pArr {
		pLine.Params = append(pLine.Params, strings.TrimSpace(p))
	}

	return pLine
}

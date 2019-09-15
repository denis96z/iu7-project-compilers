package exec

import (
	"fmt"
	"log"
)

func (s *Script) logUndefined(name string) {
	s.logError("%q is undefined", name)
}

func (s *Script) logRedefined(name string) {
	s.logError("%q is redefined", name)
}

func (s *Script) logWrongFormat(format string, args ...interface{}) {
	s.logError(fmt.Sprintf("wrong format: %s", format), args...)
}

func (s *Script) logError(format string, args ...interface{}) {
	log.Fatalf(fmt.Sprintf("[line %d]: %s", s.lineIdx, format), args...)
}

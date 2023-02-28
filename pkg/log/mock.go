package log
// NilLogger is a dummy struct that allows to avoid logging anything
type NilLogger struct{}

func (n *NilLogger)Debug(args ...any){ return }
func (n *NilLogger)Debugf(format string, args ...any){ return }
func (n *NilLogger)Fatal(args ...any){ return }
func (n *NilLogger)Fatalf(format string, args ...any){ return }

package neutrino

import (
	"github.com/gcash/bchd/addrmgr"
	"github.com/gcash/bchd/blockchain"
	"github.com/gcash/bchd/peer"
	"github.com/gcash/bchd/txscript"
	"github.com/gcash/bchlog"

	"github.com/dcrlabs/neutrino-bch/blockntfns"
	"github.com/dcrlabs/neutrino-bch/pushtx"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log bchlog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	log = bchlog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using bchlog.
func UseLogger(logger bchlog.Logger) {
	log = logger
	blockchain.UseLogger(logger)
	txscript.UseLogger(logger)
	peer.UseLogger(logger)
	addrmgr.UseLogger(logger)
	blockntfns.UseLogger(logger)
	pushtx.UseLogger(logger)
}

// logClosure is used to provide a closure over expensive logging operations so
// don't have to be performed when the logging level doesn't warrant it.
type logClosure func() string

// String invokes the underlying function and returns the result.
func (c logClosure) String() string {
	return c()
}

// newLogClosure returns a new closure over a function that returns a string
// which itself provides a Stringer interface so that it can be used with the
// logging system.
func newLogClosure(c func() string) logClosure {
	return logClosure(c)
}

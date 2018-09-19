/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package core

// Logger is the interface for loggers used in the Insolar components.
type Logger interface {
	// SetLevel sets log level.
	SetLevel(string) error
	// GetLevel gets log level.
	GetLevel() string

	// Debug logs a event at level Debug.
	Debug(...interface{})
	// Debugln logs a event at level Debug.
	Debugln(...interface{})
	// Debugf formatted logs a event at level Debug.
	Debugf(string, ...interface{})

	// Info logs a event at level Info.
	Info(...interface{})
	// Infoln logs a event at level Info.
	Infoln(...interface{})
	// Infof formatted logs a event at level Info.
	Infof(string, ...interface{})

	// Warn logs a event at level Warn.
	Warn(...interface{})
	// Warnln logs a event at level Warn.
	Warnln(...interface{})
	// Warnf formatted logs a event at level Warn.
	Warnf(string, ...interface{})

	// Error logs a event at level Error.
	Error(...interface{})
	// Errorln logs a event at level Error.
	Errorln(...interface{})
	// Errorf formatted logs a event at level Error.
	Errorf(string, ...interface{})

	// Fatal logs a event at level Fatal and than call os.exit().
	Fatal(...interface{})
	// Fatalln logs a event at level Fatal and than call os.exit().
	Fatalln(...interface{})
	// Fatalf formatted logs a event at level Fatal and than call os.exit().
	Fatalf(string, ...interface{})

	// Panic logs a event at level Panic and than call panic().
	Panic(...interface{})
	// Panicln logs a event at level Panic and than call panic().
	Panicln(...interface{})
	// Panicf formatted logs a event at level Panic and than call panic().
	Panicf(string, ...interface{})
}

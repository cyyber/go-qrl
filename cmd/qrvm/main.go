// Copyright 2014 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// qrvm executes QRVM code snippets.
package main

import (
	"fmt"
	"math/big"
	"os"
	"slices"

	"github.com/theQRL/go-qrl/cmd/qrvm/internal/t8ntool"
	"github.com/theQRL/go-qrl/internal/flags"
	"github.com/urfave/cli/v2"
)

var (
	DebugFlag = &cli.BoolFlag{
		Name:     "debug",
		Usage:    "output full trace logs",
		Category: flags.VMCategory,
	}
	MemProfileFlag = &cli.StringFlag{
		Name:     "memprofile",
		Usage:    "creates a memory profile at the given path",
		Category: flags.LoggingCategory,
	}
	CPUProfileFlag = &cli.StringFlag{
		Name:     "cpuprofile",
		Usage:    "creates a CPU profile at the given path",
		Category: flags.LoggingCategory,
	}
	StatDumpFlag = &cli.BoolFlag{
		Name:     "statdump",
		Usage:    "displays stack and heap memory information",
		Category: flags.VMCategory,
	}
	CodeFlag = &cli.StringFlag{
		Name:     "code",
		Usage:    "QRVM code",
		Category: flags.VMCategory,
	}
	CodeFileFlag = &cli.StringFlag{
		Name:     "codefile",
		Usage:    "File containing QRVM code. If '-' is specified, code is read from stdin ",
		Category: flags.VMCategory,
	}
	GasFlag = &cli.Uint64Flag{
		Name:     "gas",
		Usage:    "gas limit for the qrvm",
		Value:    10000000000,
		Category: flags.VMCategory,
	}
	PriceFlag = &flags.BigFlag{
		Name:     "price",
		Usage:    "price set for the qrvm",
		Value:    new(big.Int),
		Category: flags.VMCategory,
	}
	ValueFlag = &flags.BigFlag{
		Name:     "value",
		Usage:    "value set for the qrvm",
		Value:    new(big.Int),
		Category: flags.VMCategory,
	}
	DumpFlag = &cli.BoolFlag{
		Name:     "dump",
		Usage:    "dumps the state after the run",
		Category: flags.VMCategory,
	}
	InputFlag = &cli.StringFlag{
		Name:     "input",
		Usage:    "input for the QRVM",
		Category: flags.VMCategory,
	}
	InputFileFlag = &cli.StringFlag{
		Name:     "inputfile",
		Usage:    "file containing input for the QRVM",
		Category: flags.VMCategory,
	}
	VerbosityFlag = &cli.IntFlag{
		Name:     "verbosity",
		Usage:    "sets the verbosity level",
		Category: flags.LoggingCategory,
	}
	BenchFlag = &cli.BoolFlag{
		Name:     "bench",
		Usage:    "benchmark the execution",
		Category: flags.VMCategory,
	}
	CreateFlag = &cli.BoolFlag{
		Name:     "create",
		Usage:    "indicates the action should be create rather than call",
		Category: flags.VMCategory,
	}
	GenesisFlag = &cli.StringFlag{
		Name:     "prestate",
		Usage:    "JSON file with prestate (genesis) config",
		Category: flags.VMCategory,
	}
	MachineFlag = &cli.BoolFlag{
		Name:     "json",
		Usage:    "output trace logs in machine readable format (json)",
		Category: flags.VMCategory,
	}
	SenderFlag = &cli.StringFlag{
		Name:     "sender",
		Usage:    "The transaction origin",
		Category: flags.VMCategory,
	}
	ReceiverFlag = &cli.StringFlag{
		Name:     "receiver",
		Usage:    "The transaction receiver (execution context)",
		Category: flags.VMCategory,
	}
	DisableMemoryFlag = &cli.BoolFlag{
		Name:     "nomemory",
		Value:    true,
		Usage:    "disable memory output",
		Category: flags.VMCategory,
	}
	DisableStackFlag = &cli.BoolFlag{
		Name:     "nostack",
		Usage:    "disable stack output",
		Category: flags.VMCategory,
	}
	DisableStorageFlag = &cli.BoolFlag{
		Name:     "nostorage",
		Usage:    "disable storage output",
		Category: flags.VMCategory,
	}
	DisableReturnDataFlag = &cli.BoolFlag{
		Name:     "noreturndata",
		Value:    true,
		Usage:    "enable return data output",
		Category: flags.VMCategory,
	}
)

var stateTransitionCommand = &cli.Command{
	Name:    "transition",
	Aliases: []string{"t8n"},
	Usage:   "executes a full state transition",
	Action:  t8ntool.Transition,
	Flags: []cli.Flag{
		t8ntool.TraceFlag,
		t8ntool.TraceEnableMemoryFlag,
		t8ntool.TraceDisableStackFlag,
		t8ntool.TraceEnableReturnDataFlag,
		t8ntool.OutputBasedir,
		t8ntool.OutputAllocFlag,
		t8ntool.OutputResultFlag,
		t8ntool.OutputBodyFlag,
		t8ntool.InputAllocFlag,
		t8ntool.InputEnvFlag,
		t8ntool.InputTxsFlag,
		t8ntool.ForknameFlag,
		t8ntool.ChainIDFlag,
		t8ntool.VerbosityFlag,
	},
}

var transactionCommand = &cli.Command{
	Name:    "transaction",
	Aliases: []string{"t9n"},
	Usage:   "performs transaction validation",
	Action:  t8ntool.Transaction,
	Flags: []cli.Flag{
		t8ntool.InputTxsFlag,
		t8ntool.ChainIDFlag,
		t8ntool.ForknameFlag,
		t8ntool.VerbosityFlag,
	},
}

var blockBuilderCommand = &cli.Command{
	Name:    "block-builder",
	Aliases: []string{"b11r"},
	Usage:   "builds a block",
	Action:  t8ntool.BuildBlock,
	Flags: []cli.Flag{
		t8ntool.OutputBasedir,
		t8ntool.OutputBlockFlag,
		t8ntool.InputHeaderFlag,
		t8ntool.InputWithdrawalsFlag,
		t8ntool.InputTxsRlpFlag,
		t8ntool.VerbosityFlag,
	},
}

// vmFlags contains flags related to running the QRVM.
var vmFlags = []cli.Flag{
	CodeFlag,
	CodeFileFlag,
	CreateFlag,
	GasFlag,
	PriceFlag,
	ValueFlag,
	InputFlag,
	InputFileFlag,
	GenesisFlag,
	SenderFlag,
	ReceiverFlag,
}

// traceFlags contains flags that configure tracing output.
var traceFlags = []cli.Flag{
	BenchFlag,
	DebugFlag,
	DumpFlag,
	MachineFlag,
	StatDumpFlag,
	DisableMemoryFlag,
	DisableStackFlag,
	DisableStorageFlag,
	DisableReturnDataFlag,
}

var app = flags.NewApp("the qrvm command line interface")

func init() {
	app.Flags = slices.Concat(vmFlags, traceFlags, []cli.Flag{
		VerbosityFlag,
		MemProfileFlag,
		CPUProfileFlag,
	})
	app.Commands = []*cli.Command{
		runCommand,
		blockTestCommand,
		stateTestCommand,
		stateTransitionCommand,
		transactionCommand,
		blockBuilderCommand,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		code := 1
		if ec, ok := err.(*t8ntool.NumberedError); ok {
			code = ec.ExitCode()
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}
}

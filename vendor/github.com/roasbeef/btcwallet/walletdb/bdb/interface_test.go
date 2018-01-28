// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// This file intended to be copied into each backend driver directory.  Each
// driver should have their own driver_test.go file which creates a database and
// invokes the testInterface function in this file to ensure the driver properly
// implements the interface.  See the bdb backend driver for a working example.
//
// NOTE: When copying this file into the backend driver folder, the package name
// will need to be changed accordingly.

package bdb_test

import (
	"os"
	"testing"

	"github.com/roasbeef/btcwallet/walletdb/walletdbtest"
)

// TestInterface performs all interfaces tests for this database driver.
func TestInterface(t *testing.T) {
	dbPath := "interfacetest.db"
	defer os.RemoveAll(dbPath)
	walletdbtest.TestInterface(t, dbType, dbPath)
}

// Copyright 2018 The issues2markdown Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/repejota/issues2markdown"
	"github.com/spf13/cobra"
)

var (
	verboseFlag bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "issues2markdown",
	Short: "Convert a list of issues to markdown",
	Long:  `issues2markdown converts a list of github issues to markdown list format`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)

		// by default logging is off
		log.SetOutput(ioutil.Discard)

		// --verbose
		// enable logging if verbose mode
		if verboseFlag {
			log.SetOutput(os.Stdout)
		}

		i2md := issues2markdown.NewIssuesToMarkdown()

		log.Println("Querying data ...")
		qoptions := issues2markdown.NewQueryOptions()
		issues, err := i2md.Query(qoptions)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Rendering data ...")
		roptions := &issues2markdown.RenderOptions{}
		result, err := i2md.Render(issues, roptions)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	},
}

// Execute adds all child commands to the root command and sets flags
// appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Setup Cobra
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "enable verbose mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Unimplemented
}

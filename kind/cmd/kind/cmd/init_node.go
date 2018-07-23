/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var initNodeCmd = &cobra.Command{
	// TODO(bentheelder): more detailed usage?
	Use: "init-node",
	// this is not a user-facing command, it runs inside the "node" containers
	Hidden: true,
	Short:  "Creates a cluster",
	Long:   "Creates a Kubernetes cluster",
	Run:    runInitNode,
}

func runInitNode(cmd *cobra.Command, args []string) {
	// TODO(bentheelder): do actual node init here
	println("init node")
}

func init() {
	rootCmd.AddCommand(initNodeCmd)
}

/**
 * Created by zhouwenzhe on 2023/5/21
 */

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of zscaffold",
	Long:  `All software has versions. This is zscaffold's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zscaffold v1 -- HEAD")
	},
}

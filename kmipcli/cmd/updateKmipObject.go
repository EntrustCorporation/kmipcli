/*
 Copyright 2020-2025 Entrust Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
    // standard
    "os"
    "fmt"
    "bytes"
    "strconv"
    "encoding/json"
    
    
    // external
    "github.com/spf13/cobra"
)


// updateKmipObjectCmd represents the update-kmip-object command
var updateKmipObjectCmd = &cobra.Command{
    Use:   "update-kmip-object",
    Short: "Update Kmip Object",
    Run: func(cmd *cobra.Command, args []string) {
        flags := cmd.Flags()
        params := map[string]interface{}{}

        

        // Kmip Object uuid
        uuid, _ := flags.GetString("uuid")
        params["uuid"] = uuid

        // Action to perform on the object
        action, _ := flags.GetString("action")
        params["action"] = action

        // revcode
        if flags.Changed("revcode") {
            revcode, _ := flags.GetString("revcode")

            revcode_int, err := strconv.Atoi(revcode)
            if err != nil {
                fmt.Println(err)
                os.Exit(2)
            }

            params["revcode"] = revcode_int
        }

        // revmsg
        if flags.Changed("revmsg") {
            revmsg, _ := flags.GetString("revmsg")
            params["revmsg"] = revmsg
        }

        // JSONify
        jsonParams, err := json.Marshal(params)
        if (err != nil) {
            fmt.Println("Error building JSON request: ", err)
            os.Exit(1)
        }

        // now POST
        endpoint := GetEndPoint("", "1.0", "UpdateKmipObject")
        ret, err := DoPost(endpoint,
                               GetCACertFile(),
                               AuthTokenKV(),
                               jsonParams,
                               "application/json")
        if err != nil {
            fmt.Printf("\nHTTP request failed: %s\n", err)
            os.Exit(4)
        } else {
            // type assertion
            retBytes := ret["data"].(*bytes.Buffer)
            retStatus := ret["status"].(int)
            retStr := retBytes.String()

            if (retStr == "" && retStatus == 404) {
                fmt.Println("\nAction denied\n")
                os.Exit(5)
            }

			if retStatus == 204 {
				fmt.Println("\nUpdate successful\n")
				os.Exit(0)
			} else {
				fmt.Println("\n" + retStr + "\n")
				// make a decision on what to exit with
				retMap := JsonStrToMap(retStr)
				if _, present := retMap["error"]; present {
					os.Exit(3)
				} else {
					fmt.Println("\nUnknown error\n")
					os.Exit(0)
				}
			}
        }
    },
}

func init() {
    rootCmd.AddCommand(updateKmipObjectCmd)
    updateKmipObjectCmd.Flags().StringP("uuid", "u", "",
                        "uuid of the KMIP Object to update")
    updateKmipObjectCmd.Flags().StringP("action", "a", "",
                        "action to perform on the KMIP Object")
    updateKmipObjectCmd.Flags().StringP("revcode", "c", "",
                        "Revocation code if action is revoke")
    updateKmipObjectCmd.Flags().StringP("revmsg", "m", "",
                        "Revocation msg if action is revoke")

    // mark mandatory fields as required
    updateKmipObjectCmd.MarkFlagRequired("uuid")
    updateKmipObjectCmd.MarkFlagRequired("action")
}

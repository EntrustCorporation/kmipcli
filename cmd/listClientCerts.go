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
    "encoding/json"
    
    
    // external
    "github.com/spf13/cobra"
)


// listClientCertsCmd represents the list-client-certs command
var listClientCertsCmd = &cobra.Command{
    Use:   "list-client-certs",
    Short: "List all Kmip Client Certificate",
    Run: func(cmd *cobra.Command, args []string) {
        flags := cmd.Flags()
        params := map[string]interface{}{}

        

        // prefix
        if flags.Changed("prefix") {
            prefix, _ := flags.GetString("prefix")
            params["prefix"] = prefix
        }
        // filters
        if flags.Changed("filters") {
            filters, _ := flags.GetString("filters")
            params["filters"] = filters
        }
        // max items
        if flags.Changed("max-items") {
            maxItems, _ := flags.GetInt("max-items")
            params["max_items"] = maxItems
        }
        // fields
        if flags.Changed("field") {
            fieldsArray, _ := flags.GetStringArray("field")
            params["fields"] = fieldsArray
        }
        // next token
        if flags.Changed("next-token") {
            nextToken, _ := flags.GetString("next-token")
            params["next_token"] = nextToken
        }

        // JSONify
        jsonParams, err := json.Marshal(params)
        if (err != nil) {
            fmt.Println("Error building JSON request: ", err)
            os.Exit(1)
        }

        // now POST
        endpoint := GetEndPoint("", "1.0", "ListClientCertificates")
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

            fmt.Println("\n" + retStr + "\n")

            // make a decision on what to exit with
            retMap := JsonStrToMap(retStr)
            if _, present := retMap["error"]; present {
                os.Exit(3)
            } else {
                os.Exit(0)
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(listClientCertsCmd)
    listClientCertsCmd.Flags().StringP("prefix", "p", "",
                                 "Prefix to list Client Certs with")
    listClientCertsCmd.Flags().StringP("filters", "l", "",
                                 "Conditional expression to list filtered " +
                                 "Client Certs")
    listClientCertsCmd.Flags().IntP("max-items", "m", 0,
                              "Maximum number of items to include " +
                              "in the response")
    listClientCertsCmd.Flags().StringArrayP("field", "f", []string{},
                            "Client Cert field to include in the response")
    listClientCertsCmd.Flags().StringP("next-token", "n", "",
                            "Token from which subsequent Client Certs would " +
                            "be listed")
}

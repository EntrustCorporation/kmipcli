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
    "strings"
    "encoding/json"
    
    
    // external
    "github.com/spf13/cobra"
)


// getKmipObjectCountCmd represents the get-kmip-object-count command
var getKmipObjectCountCmd = &cobra.Command{
    Use:   "get-kmip-object-count",
    Short: "Get count of Kmip Objects",
    Run: func(cmd *cobra.Command, args []string) {
        flags := cmd.Flags()
        params := map[string]interface{}{}
        filterParams := []interface{}{}

        allowedFilters := map[string]bool {
            "uuid": true,
            "created_at": true,
            "lastmod_at": true,
            "State": true,
            "type": true,
        }

        

        if flags.Changed("filter") {
            filters, _ := flags.GetStringArray("filter")
            for i := 0; i < len(filters); i +=1 {
                filter := strings.Split(filters[i], "||")

                // remove trailing/leading whitespace
                for index := range filter {
                    filter[index] = strings.TrimSpace(filter[index])
                }

                if len(filter) != 3 {
                    fmt.Printf("\nInvalid filter argument: %s\n", filters[i])
                    os.Exit(1)
                }

                filterMap := map[string]interface{}{}
                filterMap["field"] = filter[0]
                if _, ok := allowedFilters[filter[0]]; ok {
                    filterMap["op"] = filter[1]
                    filterMap["value"] = filter[2]
                } else {
                    fmt.Println("Invalid filter type")
                    os.Exit(1)
                }

                filterParams = append(filterParams, filterMap)
            }

            params["filters"] = filterParams
        }

        // JSONify
        jsonParams, err := json.Marshal(params)
        if (err != nil) {
            fmt.Println("Error building JSON request: ", err)
            os.Exit(1)
        }

        // now POST
        endpoint := GetEndPoint("", "1.0", "GetKmipObjectCount")
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
    rootCmd.AddCommand(getKmipObjectCountCmd)
    getKmipObjectCountCmd.Flags().StringArrayP("filter", "f", []string{},
        "|| separated string containing filter field, " +
        "operation and value of the field to be filtered")
}

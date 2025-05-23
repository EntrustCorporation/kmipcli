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


// updateKmipSettingsCmd represents the update-kmip-settings command
var updateKmipSettingsCmd = &cobra.Command{
    Use:   "update-tenant-settings",
    Short: "Update Kmip Tenant settings",
    Run: func(cmd *cobra.Command, args []string) {
        flags := cmd.Flags()
        params := map[string]interface{}{}

        // create request payload
        degradedModeAvailProvided := flags.Changed("degraded-mode-availability")
        oidcProvided := flags.Changed("oidc-enabled")
        if (!degradedModeAvailProvided && !oidcProvided) {
                fmt.Println("Specify any one of Vault settings to update")
                os.Exit(1)
        }

        if flags.Changed("degraded-mode-availability") {
            degradedModeAvailability, _ := flags.GetString("degraded-mode-availability")
            if (degradedModeAvailability == "enable" ||
                  degradedModeAvailability == "disable") {
                if (degradedModeAvailability == "enable") {
                    params["degraded_mode_availability"] = true
                }
                if (degradedModeAvailability == "disable") {
                    params["degraded_mode_availability"] = false
                }
            } else {
                fmt.Printf("\nInvalid -d, --degraded-mode-availability option %s. " +
                           "Supported: enable (or) disable\n", degradedModeAvailability)
                os.Exit(1)
            }
	    }

        if flags.Changed("oidc-enabled") {
            oidcEnabled, _ := flags.GetString("oidc-enabled")
            if (oidcEnabled == "enable" ||
                  oidcEnabled == "disable") {
                if (oidcEnabled == "enable") {
                    params["oidc_enabled"] = true
                }
                if (oidcEnabled == "disable") {
                    params["oidc_enabled"] = false
                }
            } else {
                fmt.Printf("\nInvalid -o, --oidc-enabled option %s. " +
                           "Supported: enable (or) disable\n", oidcEnabled)
                os.Exit(1)
            }
	    }

        revision, _ := flags.GetInt("revision")
        params["revision"] = revision

        // JSONify
        jsonParams, err := json.Marshal(params)
        if (err != nil) {
            fmt.Println("Error building JSON request: ", err)
            os.Exit(1)
        }

        // now POST
        endpoint := GetEndPoint("", "1.0", "UpdateTenantSettings")
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
                fmt.Println("\nKmip Settings not found\n")
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
    rootCmd.AddCommand(updateKmipSettingsCmd)
    updateKmipSettingsCmd.Flags().StringP("degraded-mode-availability", "d", "",
                                 "Degraded mode availability. ")
    updateKmipSettingsCmd.Flags().StringP("oidc-enabled", "o", "",
                                 "OIDC enabled flag. ")
    updateKmipSettingsCmd.Flags().IntP("revision", "R", 0,
                              "Revision number of the box")

    // mark mandatory fields as required
    updateKmipSettingsCmd.MarkFlagRequired("revision")
}

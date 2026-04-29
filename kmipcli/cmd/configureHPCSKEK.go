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
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	// external
	"github.com/spf13/cobra"
)

// ConfigureHPCSKEK represents the configure-hpcs-kek command
var ConfigureHPCSKEK = &cobra.Command{
	Use:    "configure-hpcs-kek",
	Short:  "Configure Encrypting KMIP objects using KEK stored in IBM Hyper Protect Crypto Services",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		params := map[string]interface{}{}

		// revision
		revision, _ := flags.GetInt("revision")
		params["revision"] = revision

		params["hsm_type"] = "HPCS"

		hpcs_url, _ := flags.GetString("hpcs_url")
		params["hpcs_url"] = hpcs_url

		hpcs_api_key, _ := flags.GetString("hpcs_api_key")
		params["hpcs_api_key"] = hpcs_api_key

		hpcs_instance_id, _ := flags.GetString("hpcs_instance_id")
		params["hpcs_instance_id"] = hpcs_instance_id

		if flags.Changed("hpcs_root_key_id") {
			hpcs_root_key_id, _ := flags.GetString("hpcs_root_key_id")
			params["hpcs_root_key_id"] = hpcs_root_key_id
		}

		if flags.Changed("kek_cache_timeout") {
			//kek_cache_timeout
			kek_cache_timeout, _ := flags.GetInt("kek_cache_timeout")
			params["kek_cache_timeout"] = kek_cache_timeout
		}

		// JSONify
		jsonParams, err := json.Marshal(params)
		if err != nil {
			fmt.Println("Error building JSON request: ", err)
			os.Exit(1)
		}

		// now POST
		endpoint := GetEndPoint("", "1.0", "UpdateKEKSetting")
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

			if retStr == "" && retStatus == 404 {
				fmt.Println("\nTenant not found\n")
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
	rootCmd.AddCommand(ConfigureHPCSKEK)
	ConfigureHPCSKEK.Flags().IntP("kek_cache_timeout", "t", 1800,
		"KEK Cache Timeout")
	ConfigureHPCSKEK.Flags().IntP("revision", "R", 0,
		"KEK Setting revision number")
	ConfigureHPCSKEK.Flags().StringP("hpcs_url", "u", "",
		"HPCS URL")
	ConfigureHPCSKEK.Flags().StringP("hpcs_api_key", "k", "",
		"HPCS API Key")
	ConfigureHPCSKEK.Flags().StringP("hpcs_instance_id", "i", "",
		"HPCS Instance ID")
	ConfigureHPCSKEK.Flags().StringP("hpcs_root_key_id", "r", "",
		"KEK Rootkey ID")
	// mark mandatory fields as required
	ConfigureHPCSKEK.MarkFlagRequired("hpcs_url")
	ConfigureHPCSKEK.MarkFlagRequired("hpcs_api_key")
	ConfigureHPCSKEK.MarkFlagRequired("hpcs_instance_id")
	ConfigureHPCSKEK.MarkFlagRequired("revision")
}

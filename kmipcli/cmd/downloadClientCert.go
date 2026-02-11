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
    "os"
    "fmt"
    "encoding/json"
    
    "github.com/spf13/cobra"
)


// downloadClientCertCmd represents the download-client-cert command
var downloadClientCertCmd = &cobra.Command{
    Use:   "download-client-cert",
    Short: "Download Kmip Client Certificate",
    Run: func(cmd *cobra.Command, args []string) {
        flags := cmd.Flags()
        params := map[string]interface{}{}

        

        // client certificate id
        certid, _ := flags.GetString("certid")
        params["client_cert_id"] = certid

        // JSONify
        jsonParams, err := json.Marshal(params)
        if (err != nil) {
            fmt.Println("Error building JSON request: ", err)
            os.Exit(1)
        }

        // now POST
        endpoint := GetEndPoint("", "1.0", "DownloadClientCertificate")
        fname, err := DoDownload(
                        endpoint,
                        "POST",
                        GetCACertFile(),
                        AuthTokenKV(),
                        jsonParams)
        if err != nil {
            fmt.Printf("\nHTTP request failed: %s\n", err)
            os.Exit(4)
        } else {
            fmt.Println("\nSuccessfully downloaded client cert bundle " +
              "as - " + fname + "\n")
        }
    },
}

func init() {
    rootCmd.AddCommand(downloadClientCertCmd)
    downloadClientCertCmd.Flags().StringP("certid", "c", "",
            "Id or name of the KMIP Client Certificate to download")

    // mark mandatory fields as required
    downloadClientCertCmd.MarkFlagRequired("certid")
}

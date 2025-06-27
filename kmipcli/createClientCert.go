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


// createClientCertCmd represents the create-client-cert command
var createClientCertCmd = &cobra.Command{
    Use:   "create-client-cert",
    Short: "Create a Kmip Client Certificate",
    Run: func(cmd *cobra.Command, args []string) {
        flags := cmd.Flags()
        params := map[string]interface{}{}

        

        // Client Certificate name
        certName, _ := flags.GetString("name")
        params["name"] = certName

        // Expiry days
        expiryDays, _ := flags.GetInt("expiry-days")
        params["expiry_days"] = expiryDays

        // csr
        if flags.Changed("csr") {
            csr, _ := flags.GetString("csr")
            params["csr"] = csr
        }

        // cert
        if flags.Changed("cert") {
            cert, _ := flags.GetString("cert")
            params["cert"] = cert
        }

        // passphrase
        if flags.Changed("passphrase") {
            passphrase, _ := flags.GetString("passphrase")
            params["passphrase"] = passphrase
        }

        // JSONify
        jsonParams, err := json.Marshal(params)
        if (err != nil) {
            fmt.Println("Error building JSON request: ", err)
            os.Exit(1)
        }

        // now POST
        endpoint := GetEndPoint("", "1.0", "CreateClientCertificate")
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
    rootCmd.AddCommand(createClientCertCmd)
    createClientCertCmd.Flags().StringP("name", "n", "",
                                    "Name of the Client Certificate")
    createClientCertCmd.Flags().IntP("expiry-days", "e", 0,
                              "Number of days after which the KMIP Client " +
                              "Certificate will expire")
    createClientCertCmd.Flags().StringP("csr", "c", "",
                                    "Certificate Signing Request to be used " +
                                    "for generating KMIP Client Certificate ")
    createClientCertCmd.Flags().StringP("cert", "t", "",
                                    "Certificate to be used for KMIP Client")
    createClientCertCmd.Flags().StringP("passphrase", "p", "",
                                    "Passphrase to be used if the KMIP " +
                                    "Client Certificate has to be encrypted")

    // mark mandatory fields as required
    createClientCertCmd.MarkFlagRequired("name")
    createClientCertCmd.MarkFlagRequired("expiry-days")
}

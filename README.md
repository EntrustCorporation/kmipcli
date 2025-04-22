#kmipcli

kmipmcli is a command line tool to manage, control Entrust KMIP type Vault. All vault management operation, e.g user management, KMIP key operations(e.g.: revoke, destroy, rekey), HSM Management, KMIP Client certificate management, can be done by kmipcli. 

Entrust KMIP vault provides secure storage in addition to all key management features described by OASIS KMIP specification.

1. The kmipcli requires Entrust KeyControl version 5.2 or later.
2. The KMIP Vault must be created by the KeyControl Vault Administrator.
3. To manage the Vault using the kmipcli, you must be the admin for that KMIP Vault and have the login API URL of that vault.
4. All users authorized to access the Vault can use the kmipcli with the login URL.

## Releases

kmipcli's for Linux & Windows for each release can be found in Releases section (https://github.com/EntrustCorporation/kmipcli/releases) 

##Build instruction

The code in this repo corresponds to the latest released version of kmipcli. In general, to use kmipcli, head over to Releases section to get pre-compiled binaries. If you do plan to build, follow instructions below.

1. Install go. pasmcli/Makefile expects to find go at /usr/local/go/bin/go
  . Check for latest stable Golang Linux version at https://go.dev/dl/
  . As root, do, wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
  . As root, tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
  . Go should be installed at, /usr/local/go/bin/go. To find go binary, adjust PATH as, export PATH=$PATH:/usr/local/go/bin

2. cd to kmipcli/
3. To build Linux & Windows cli binaries,

   ```$ gmake all```

4. To clean workspace,

   ```$ gmake clean```

For more information, see the Secrets Vault chapter in the Key Management Systems documentation at https://trustedcare.entrust.com/.

##Usage examples

Login to KMIP vault: 
```$./kmipcli login  -u "username"  -p "password"  --login-URL <VAULT_API_URL>  --cacert ca.cert```

Create KMIP client certificate: 
```$./kmipcli create-client-cert  --name "client-cert-name"  --expiry-days "expiry_days_in_int"```

Get client certificate details: 
```$./kmipcli get-client-cert  -c "client-cert-name"```

Download kmip client certificate: 
```$./kmipcli download-client-cert  -c "client-cert-name"```

Delete kmip client certificate: 
```$./kmipcli delete-client-cert -c "client-cert-name"```

List kmip vault local user : 
```$./kmipcli list-local-users```

Get KMIP policy details: 
```$./kmipcli list-policies```

Other Available Commands:
  change-ad-domain                Change Active Directory Domain
  configure-hpcs-kek              Configure Encrypting KMIP objects using KEK stored in IBM Hyper Protect Crypto Services
  configure-hsm-kek               Configure Encrypting KMIP objects using KEK stored in system HSM
  create-client-cert              Create a Kmip Client Certificate
  create-local-user               Create a Local User
  create-personal-access-token    Create a Personal Access Token
  create-policy                   Create a Kmip Policy
  delete-client-cert              Delete Kmip Client Certificate
  delete-local-user               Delete a Local User
  delete-personal-access-token    Delete a Personal Access Token
  delete-policy                   Delete Policy
  disable-kmip-kek                Disable encrypting KMIP objects using KEK
  download-audit                  Download audit log bundle
  download-client-cert            Download Kmip Client Certificate
  get-ad-group                    Search Active Directory group
  get-ad-setting                  Get AD Setting details
  get-ad-user                     Search Active Directory users
  get-audit-message-template      Given a message id, get corresponding audit message template
  get-audit-setting               Get audit settings
  get-client-cert                 Get Kmip Client Certificate details
  get-hsm-info                    Get System HSM configuration details
  get-kek-setting                 Get KMIP KEK settings
  get-kmip-object                 Get Kmip Object details
  get-kmip-object-count           Get count of Kmip Objects
  get-local-user                  Get Local User details
  get-personal-access-token       Get Personal Access Token details
  get-platform-info               Get Platform Info
  get-policy                      Get Policy details
  get-tenant-info                 Get Kmip Tenant Info
  get-tenant-settings             Get Kmip Tenant settings
  list-ad-settings                List all AD settings
  list-audit-message-templates    List available audit messages templates
  list-audit-messages             List all audit messages
  list-client-certs               List all Kmip Client Certificate
  list-kmip-objects               List all Kmip Objects
  list-local-users                List all Local Users
  list-personal-access-tokens     List Personal Access Tokens
  list-policies                   List all Policies
  list-policy-versions            List versions of a given Policy
  locate-root-key                 Locate KMIP KEK Root Key
  login                           Login to Kmip Tenant Portal
  rekey-kmip-kek                  Rekey KMIP objects with new KEK
  renew                           Renew Access Token
  set-policy-version              Set a specific version of Policy to current
  update-ad-setting               Update Active Directory Setting
  update-audit-setting            Update audit settings
  update-kmip-object              Update Kmip Object
  update-local-user               Update a given Local User
  update-personal-access-token    Update a Personal Access Token
  update-policy                   Update a given Kmip Policy
  update-tenant-auth-method-to-ad Update Tenant Auth Method To AD
  update-tenant-settings          Update Kmip Tenant settings
  version                         Version of Entrust KMIP tenant portal cli

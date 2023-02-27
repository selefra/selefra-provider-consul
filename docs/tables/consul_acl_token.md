# Table: consul_acl_token

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| partition | string | X | √ | The partition the ACL token is associated with. | 
| policies | json | X | √ | List of policies. | 
| roles | json | X | √ | List of roles | 
| node_identities | json | X | √ | The list of node identities that should be applied to the token. | 
| description | string | X | √ | The token description. | 
| expiration_time | string | X | √ | If set this represents the point after which a token should be considered revoked and is eligible for destruction. | 
| local | bool | X | √ | Flag to set the token local to the current datacenter. | 
| namespace | string | X | √ |  | 
| service_identities | json | X | √ | The list of service identities that should be applied to the token. | 
| accessor_id | string | X | √ | The token id. | 
| id | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



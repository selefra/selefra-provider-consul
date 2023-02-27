# Table: consul_acl_auth_method

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | A free form human readable description of the auth method. | 
| display_name | string | X | √ | An optional name to use instead of the name attribute when displaying information about this auth method. | 
| token_locality | string | X | √ | The kind of token that this auth method produces. This can be either 'local' or 'global'. | 
| type | string | X | √ | The type of the ACL auth method. | 
| namespace_rule | json | X | √ |  | 
| config | json | X | √ | The raw configuration for this ACL auth method. | 
| id | string | X | √ |  | 
| max_token_ttl | string | X | √ | The maximum life of any token created by this auth method. | 
| name | string | X | √ | The name of the ACL auth method. | 
| namespace | string | X | √ |  | 
| partition | string | X | √ | The partition the ACL auth method is associated with. | 
| config_json | string | X | √ | The raw configuration for this ACL auth method. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



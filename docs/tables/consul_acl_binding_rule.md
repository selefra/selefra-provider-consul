# Table: consul_acl_binding_rule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| partition | string | X | √ | The partition the ACL binding rule is associated with. | 
| selector | string | X | √ | The expression used to math this rule against valid identities returned from an auth method validation. | 
| auth_method | string | X | √ | The name of the ACL auth method this rule apply. | 
| bind_name | string | X | √ | The name to bind to a token at login-time. | 
| bind_type | string | X | √ | Specifies the way the binding rule affects a token created at login. | 
| description | string | X | √ | A free form human readable description of the binding rule. | 
| id | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



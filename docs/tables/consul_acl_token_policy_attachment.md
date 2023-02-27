# Table: consul_acl_token_policy_attachment

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| policy | string | X | √ | The policy name. | 
| token_id | string | X | √ | The token accessor id. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



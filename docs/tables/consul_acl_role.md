# Table: consul_acl_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| partition | string | X | √ | The partition the ACL role is associated with. | 
| policies | json | X | √ | The list of policies that should be applied to the role. | 
| node_identities | json | X | √ | The list of node identities that should be applied to the role. | 
| service_identities | json | X | √ | The list of service identities that should be applied to the role. | 
| description | string | X | √ | A free form human readable description of the role. | 
| id | string | X | √ |  | 
| name | string | X | √ | The name of the ACL role. | 
| namespace | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



# Table: consul_admin_partition

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| name | string | X | √ | The partition name. This must be a valid DNS hostname label. | 
| description | string | X | √ | Free form partition description. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



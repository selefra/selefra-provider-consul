# Table: consul_node

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| meta | json | X | √ |  | 
| name | string | X | √ |  | 
| partition | string | X | √ | The partition the node is associated with. | 
| token | string | X | √ |  | 
| address | string | X | √ |  | 
| datacenter | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



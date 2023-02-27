# Table: consul_service

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| datacenter | string | X | √ |  | 
| external | bool | X | √ |  | 
| node | string | X | √ |  | 
| port | float | X | √ |  | 
| address | string | X | √ |  | 
| enable_tag_override | bool | X | √ |  | 
| partition | string | X | √ | The partition the service is associated with. | 
| check | json | X | √ |  | 
| meta | json | X | √ |  | 
| namespace | string | X | √ |  | 
| service_id | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



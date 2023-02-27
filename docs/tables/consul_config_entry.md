# Table: consul_config_entry

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| partition | string | X | √ | The partition the config entry is associated with. | 
| config_json | string | X | √ |  | 
| id | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



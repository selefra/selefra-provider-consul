# Table: consul_peering_token

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| meta | json | X | √ | Specifies KV metadata to associate with the peering. This parameter is not required and does not directly impact the cluster peering process. | 
| partition | string | X | √ |  | 
| peer_name | string | X | √ | The name assigned to the peer cluster. The `peer_name` is used to reference the peer cluster in service discovery queries and configuration entries such as `service-intentions`. This field must be a valid DNS hostname label. | 
| peering_token | string | X | √ | The generated peering token | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



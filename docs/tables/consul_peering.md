# Table: consul_peering

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| peer_server_addresses | json | X | √ |  | 
| peer_server_name | string | X | √ |  | 
| peering_token | string | X | √ | The peering token fetched from the peer cluster. | 
| deleted_at | string | X | √ |  | 
| peer_ca_pems | json | X | √ |  | 
| peer_name | string | X | √ | The name assigned to the peer cluster. The `peer_name` is used to reference the peer cluster in service discovery queries and configuration entries such as `service-intentions`. This field must be a valid DNS hostname label. | 
| peer_id | string | X | √ |  | 
| state | string | X | √ |  | 
| id | string | X | √ |  | 
| meta | json | X | √ | Specifies KV metadata to associate with the peering. This parameter is not required and does not directly impact the cluster peering process. | 
| partition | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 



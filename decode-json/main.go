package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println(system_npu)

	var x map[string]interface{}
	json.Unmarshal([]byte(system_npu), &x)

	hash_drill(x, "")
}

func hash_drill(val map[string]interface{}, ident string) {
	for key, value := range val {
		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			fmt.Println(ident, key, "String", value)
		case reflect.Map:
			fmt.Println(ident, key, "Map")
			hash_drill(value.(map[string]interface{}), ident+"  ")
		case reflect.Float64:
			fmt.Println(ident, key, "Float64", value)
		case reflect.Slice:
			fmt.Println(ident, key, "Slice")
			array_drill(value.([]interface{}), ident+"  ")
		default:
			fmt.Println(ident, key, "Default!!", reflect.TypeOf(value), reflect.TypeOf(value).Kind())
		}
	}
}

func array_drill(val []interface{}, ident string) {
	for _, value := range val {
		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			fmt.Println(ident, "String", value)
		case reflect.Map:
			fmt.Println(ident, "Map")
			hash_drill(value.(map[string]interface{}), ident+"  ")
		case reflect.Float64:
			fmt.Println(ident, "Float64", value)
		case reflect.Slice:
			fmt.Println(ident, "Slice")
			array_drill(value.([]interface{}), ident+"  ")
		default:
			fmt.Println(ident, "Default!!", reflect.TypeOf(value), reflect.TypeOf(value).Kind())
		}
	}
}

const system_npu = `
{
  "http_method":"GET",
  "results":{
    "name":"npu",
    "category":"complex",
    "help":"Configure NPU attributes.",
    "children":{
      "enc-offload-antireplay":{
        "name":"enc-offload-antireplay",
        "category":"unitary",
        "type":"option",
        "help":"Enable\/disable hardware offloading for IPsec anti-replay on encryption.",
        "options":[
          {
            "name":"enable",
            "help":"Enable hardware offloading for IPsec anti-replay on encryption."
          },
          {
            "name":"disable",
            "help":"Disable hardware offloading for IPsec anti-replay on encryption."
          }
        ]
      },
      "dec-offload-antireplay":{
        "name":"dec-offload-antireplay",
        "category":"unitary",
        "type":"option",
        "help":"Enable\/disable hardware offloading for IPsec anti-replay on decryption.",
        "options":[
          {
            "name":"enable",
            "help":"Enable hardware offloading for IPsec anti-replay on decryption."
          },
          {
            "name":"disable",
            "help":"Disable hardware offloading for IPsec anti-replay on decryption."
          }
        ]
      },
      "offload-ipsec-host":{
        "name":"offload-ipsec-host",
        "category":"unitary",
        "type":"option",
        "help":"Enable\/disable hardware offloading for IPsec host encryption pass.",
        "options":[
          {
            "name":"enable",
            "help":"Enable hardware offloading for IPsec host encryption pass."
          },
          {
            "name":"disable",
            "help":"Disable hardware offloading for IPsec host encryption pass."
          }
        ]
      }
    },
    "path":"system",
    "q_type":330
  },
  "vdom":"root",
  "path":"system",
  "name":"npu",
  "action":"schema",
  "status":"success",
  "http_status":200,
  "serial":"FWF60C3G11009887",
  "version":"v5.2.5",
  "build":701
}
`

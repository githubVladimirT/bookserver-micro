# #!/bin/sh -e

# INC=$(go list -f '{{ .Dir }}' -m go.unistack.org/micro-proto/v3)
# ARGS="-I${INC}"

# protoc $ARGS -Iproto --openapiv3_out=disable_default_errors=true,allow_merge=true:./proto/ --go_out=paths=source_relative:./proto/ --micro_out=components="micro|http",debug=true,paths=source_relative:./proto/ proto/*.proto

#!/bin/sh -e

INC=$(go list -f '{{ .Dir }}' -m go.unistack.org/micro-proto/v3)
INC_CODEC=$(go list -f '{{ .Dir }}' -m go.unistack.org/micro/v3)
ARGS="-I${INC}"
CODEC_ARGS="-I${INC_CODEC}"

protoc $ARGS $CODEC_ARGS -Iproto --go_out=paths=source_relative:./proto/ proto/*.proto

protoc $ARGS $CODEC_ARGS -Iproto --go-micro_out=components="micro|http",debug=true,tag_path=./proto/,paths=source_relative:./proto/ proto/*.proto

protoc $ARGS $CODEC_ARGS -Iproto --go-micro_out=openapi_file=apidocs.swagger.json,components="openapiv3",debug=true,paths=source_relative:./proto/ proto/*.proto
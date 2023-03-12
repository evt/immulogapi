export SERVICES=\
	log \
	auth

proto:
	@for SERVICE in $(SERVICES); \
		do $(call generate-service,$$SERVICE); \
	done

define generate-service
	echo "** generate protofiles for $(1)" && \
	\
	mkdir -p $(1) && \
	\
	protoc 	-I=. \
		--go_out=. \
			--go_opt=paths=source_relative \
		--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
      --grpc-gateway_out=. \
			--grpc-gateway_opt paths=source_relative \
		v1/$(1).proto
endef
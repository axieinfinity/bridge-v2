mock: ## generate mock
	mockery --output internal/mocks/types --dir internal/types --name IListener
	mockery --output internal/mocks/types --dir internal/types --name ITransaction
	mockery --output internal/mocks/types --dir internal/types --name ILog
	mockery --output internal/mocks/types --dir internal/types --name IReceipt
	mockery --output internal/mocks/types --dir internal/types --name IBlock
	mockery --output internal/mocks/types --dir internal/types --name IJob
	mockery --output internal/mocks/utils --dir internal/utils --name IUtils
	mockery --output internal/mocks/utils --dir internal/utils --name EthClient

abigen:
	abigen --abi=./contracts/common/BridgeAdmin.abi --pkg=common --out=generated_contracts/ethereum/common/bridge_admin.go
	abigen --abi=./contracts/common/BridgeAdmin.abi --pkg=common --out=generated_contracts/ronin/common/bridge_admin.go
	abigen --abi=./contracts/ethereum/MainchainGatewayV2.abi --pkg=gateway --out=./generated_contracts/ethereum/gateway/mainchain_gateway_v2.go
	abigen --abi=./contracts/ronin/RoninGatewayV2.abi --pkg=gateway --out=./generated_contracts/ronin/gateway/ronin_gateway_v2.go
	abigen --abi=./contracts/common/RoninValidator.abi --pkg=validator --out=./generated_contracts/ronin/validator/ronin_validator.go

bridge:
	go install ./cmd/bridge
	@echo "Done building."
	@echo "Run \"bridge\" to launch bridge."

run:
	@cd cmd/bridge && go run main.go
	
# todo
MOCKS_DESTINATION=mock
.PHONY: mocks
mocks: internal/db/*.go
	@echo "Generating mocks for db package..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/db_mocks.go -package=mock; done
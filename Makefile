
test:
	@echo Run unit test
	@test -v ./...

staticcheck:
	@echo Run staticcheck
	@staticcheck ./...

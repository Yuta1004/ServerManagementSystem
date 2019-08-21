GO = go


run:
	@$(GO) run main.go

test:
	@$(GO) test -v ../server-manage/...

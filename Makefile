GO = go
ENV = ~/golang-project/src/server-manage/env.sh


run:
	@source $(ENV); $(GO) run main.go

test:
	@source $(ENV); $(GO) test -v ../server-manage/...

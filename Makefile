install:
	bash install.sh

buf-gen:
	buf generate proto

buf-update:
	cd proto; buf mod update;

run-role:
	go run apps/role/main.go

run-user:
	go run apps/user/main.go

test-repo:
	cd repository/; go test -v
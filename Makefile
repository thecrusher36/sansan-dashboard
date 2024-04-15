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

mock-gen:
	mockgen -source=repository/role.go -destination=repository/mock/role_mock.go -package=mock;
	mockgen -source=repository/user.go -destination=repository/mock/user_mock.go -package=mock;
	mockgen -source=core/caller/caller.go -destination=core/caller/caller_mock.go -package=caller;
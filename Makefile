buf-gen:
	buf generate proto

buf-update:
	cd proto; buf mod update;

run-role:
	go run apps/role/main.go
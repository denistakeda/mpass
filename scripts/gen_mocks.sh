#!/usr/bin/env sh

mockgen -destination=mocks/server/server_mock.go \
    -package=server_mock \
    -source=internal/server/server.go

mockgen -destination=mocks/auth_service/auth_service_mock.go \
    -package=auth_service_mock \
    -source=internal/auth_service/auth_service.go

#!/bin/sh

export PRISMA_CLIENT_ENGINE_TYPE=dataproxy
go run github.com/polarlightsllc/prisma-client-go-1 generate --schema schema.out.prisma

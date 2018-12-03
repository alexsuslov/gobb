#!/bin/bash
(
    cd api/graphql/types;
    go run user_gen.go;
    go run settings_gen.go;
    go run post_gen.go;
    go run board_gen.go;
    go run view_gen.go;
)
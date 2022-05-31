#!/bin/bash

cat docs_readme_head.md > README.md
gomarkdoc -o README.md -e ./ctrl
gomarkdoc -o README.md -e ./db
gomarkdoc -o README.md -e ./rdg
gomarkdoc -o README.md -e ./server

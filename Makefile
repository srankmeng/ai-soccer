.PHONY: run install i

# Run the application
run:
	bun run src/index.js

install:
	bun install

i: install


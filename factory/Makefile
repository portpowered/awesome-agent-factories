.PHONY: dev lint format typecheck test build ci validate-data linkcheck validate-pdf build-search-index

dev:
	bun run dev

lint:
	bun run lint

format:
	bun run format

typecheck:
	bun run typecheck

test:
	bun run test

build:
	bun run build
	bun ./scripts/verify-phase-1-static-routes.ts

ci: lint typecheck test build validate-data

validate-data:
	bun ./scripts/validate-registry.ts

linkcheck:
	@echo "linkcheck: skipped (not yet implemented in scaffold)"
	@exit 0

validate-pdf:
	@echo "validate-pdf: skipped (not yet implemented in scaffold)"
	@exit 0

build-search-index:
	bun run build:search-index

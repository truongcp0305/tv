# Project

Go + HTMX + Templ + Tailwind CSS

Rules

- Prefer editing existing files.
- Never create duplicate components.
- Run gofmt after editing Go files.
- Keep HTMX responses partial.
- Use Templ idioms.

## graphify

This project has a graphify knowledge graph at graphify-out/.

Rules:
- Before answering architecture or codebase questions, read graphify-out/GRAPH_REPORT.md for god nodes and community structure
- If graphify-out/wiki/index.md exists, navigate it instead of reading raw files
- After modifying code files in this session, run `graphify update .` to keep the graph current (AST-only, no API cost)

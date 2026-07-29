# Templates Package

Components and layout for the Tử Vi Đẩu Số frontend.

## Structure

```
templates/
├── base.templ        — HTML boilerplate, CSS variables, Tailwind classes
├── page.templ        — Main page layout (header + form panel + chart panel)
├── forms.templ       — Form components with inputs, select, toggle
├── chart.templ       — Chart placeholder & result rendering
├── shared.templ      — Shared data (lunar months, gender options)
├── markdown.go       — Markdown → HTML converter
├── markdown.templ    — RenderMarkdown templ component
└── client.js         — Static client-side logic (HTMX + Alpine)
```

## Tech Stack

- **HTMX** — AJAX submissions without writing JavaScript
- **Go Templ** — Type-safe server-side templating
- **Tailwind CSS** — Utility-first styling (CDN in dev, can be built for prod)
- **Alpine.js** — Light interactivity where HTMX isn't enough (gender toggle)

## Usage

1. Go to project root and run:
   ```bash
   cd D:\dev\letcoode\src
   go run .  # Start the API server
   ```

2. Open `http://localhost:3088` in browser.

3. The static demo is at `D:\dev\letcoode\src\static\index.html` — open directly in browser.

## Responsive Breakpoints

- Desktop: 40% form / 60% chart (lg breakpoint: ≥1024px)
- Mobile: stacked vertically (< 1024px)

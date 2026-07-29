# Frontend Demo — Tử Vi Đẩu Số

## Cấu trúc

```
src/
├── templates/          # Go Templ components
│   ├── base.templ     — HTML boilerplate, CSS variables, Tailwind classes
│   ├── page.templ     — Main page layout (header + form panel + chart panel)
│   ├── forms.templ    — Form components (inputs, selects, toggle)
│   ├── chart.templ    — Chart placeholder & result rendering
│   ├── shared.templ   — Shared data (lunar months, gender options)
│   ├── markdown.go    — Markdown → HTML converter
│   └── markdown.templ — RenderMarkdown templ component
│
├── static/
│   ├── index.html     — Standalone demo HTML (no server needed)
│   └── client.js      — Client-side logic for HTMX + Alpine
│
└── demo/
    ─── main.go          — Simple HTTP server to serve static files
    ─── go.mod           — Go module for demo
    ─── static/          — Embedded static files
```

## Cách chạy

### Option 1: Mở trực tiếp file HTML
Mở `D:\dev\letcoode\src\static\index.html` trong trình duyệt — không cần server.

### Option 2: Chạy demo server
```bash
cd D:\dev\letcoode\src\demo
.\tv_demo.exe
```
Mở http://localhost:3088

## Công nghệ sử dụng

- **HTMX** — AJAX submissions không cần viết JS
- **Go Templ** — Type-safe server-side templating
- **Tailwind CSS** — Utility-first styling (CDN)
- **Alpine.js** — Light interactivity (gender toggle)

## Responsive Design

| Desktop (≥1024px) | Mobile (<1024px) |
|-------------------|------------------|
| Form 40% / Chart 60% | Form trên / Chart dưới |
| Hai cột song song | Stack dọc |

## Components có sẵn

### Base (`base.templ`)
HTML boilerplate với CSS variables, Tailwind classes import.

### Page (`page.templ`)
Main layout: header sticky + form panel (trái) + chart panel (phải).

### Form (`forms.templ`)
- Input text (họ tên)
- Gender buttons (Nam/Nữ)
- Birth date (day/month/year)
- Birth time (hour/time period)
- Province/city select
- Submit button with loading animation

### Chart (`chart.templ`)
- Empty state với skeleton
- Result area với markdown rendering

## Lưu ý

- Không có backend logic — chỉ UI components
- Các form actions cần được connect với API endpoint `/api/horoscope`
- HTMX xử lý form submission và swap nội dung vào `#chart-output`

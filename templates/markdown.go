package templates

import (
	"fmt"
	"strings"
)

// simpleHTML renders markdown-like text into safe HTML blocks.
func simpleHTML(md string) string {
	raw := strings.TrimSpace(md)
	if raw == "" {
		return `<p class="text-sm text-slate-400 italic">Đang xử lý...</p>`
	}

	lines := strings.Split(raw, "\n")
	var parts []string
	inUL := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### "):
			if inUL {
				parts = append(parts, "</ul>")
				inUL = false
			}
			prefix := strings.TrimLeft(line, "#")
			text := strings.TrimSpace(prefix)
			if text == "" {
				continue
			}
			tag := "h1"
			cls := "text-xl font-semibold mt-6 mb-3"
			if prefix[0] == '#' {
				prefix = prefix[1:]
			} else {
				tag = "h2"
				cls = "text-lg font-semibold mt-5 mb-2"
				if prefix[0] == '#' {
					prefix = prefix[1:]
				} else {
					tag = "h3"
					cls = "text-base font-semibold mt-4 mb-1.5"
				}
			}
			parts = append(parts, fmt.Sprintf("<%s class=\"%s\">%s</%s>", tag, cls, text, tag))
		case strings.HasPrefix(line, "- "):
			if !inUL {
				parts = append(parts, "<ul class=\"space-y-1 ml-5 list-disc my-2 text-sm leading-relaxed\">")
				inUL = true
			}
			parts = append(parts, fmt.Sprintf("<li>%s</li>", strings.TrimSpace(strings.TrimPrefix(line, "- "))))
		default:
			if inUL {
				parts = append(parts, "</ul>")
				inUL = false
			}
			parts = append(parts, fmt.Sprintf("<p class=\"text-sm leading-relaxed text-slate-700 my-2\">%s</p>", line))
		}
	}
	if inUL {
		parts = append(parts, "</ul>")
	}

	return strings.Join(parts, "\n")
}

# tealayout

Layout managers for [Bubble Tea](https://github.com/charmbracelet/bubbletea)
TUIs.

## Components

- **`BorderLayout`** — five regions (`North`, `West`, `Center`, `East`,
  `South`), each holding a `tea.Model`. Regions can be fixed-size or flex, and
  shown/hidden at runtime. Implements `tea.Model`, `Resizable`, and
  `LayoutManager`.
- **`LayoutComponent`** — wraps any `tea.Model` in a
  [`viewport`](https://github.com/charmbracelet/bubbles) with an optional
  [`lipgloss`](https://github.com/charmbracelet/lipgloss) style; handles
  resizing.

## Showcase

```bash
go run ./showcase        # type n / s / e / w / c to toggle regions; writes debug.log
```

## Status

Experimental. The module path in `go.mod` is `github.com/dan.frohlch/tealayout`
(`dan.frohlch` is a typo for `dan-frohlich`).

# [NathanBurkett](https://github.com/NathanBurkett) / [diff-chart](https://github.com/NathanBurkett/diff-chart)

![](https://github.com/nathanburkett/diff-chart/workflows/Continuous%20Integration/badge.svg)

```bash
A fast and flexible diff statistic generator built with
        love by Nathan Burkett in Go.

Usage:
  diff-chart [command]

Available Commands:
  diff        Parse a diff
  help        Help about any command

Flags:
  -h, --help   help for diff-chart

Use "diff-chart [command] --help" for more information about a command.
```

## Example

```bash
diff-chart diff -d ae13803 -u 51b256f
```
yields
```
| Directory | +/- | Δ % |
| --- | --- | --- |
| go.sum | +167/-0 | 21.44% |
| run | +142/-0 | 18.23% |
| algorithm | +107/-8 | 14.76% |
| input | +95/-11 | 13.61% |
| transform | +88/-1 | 11.42% |
| output | +86/-0 | 11.04% |
| cmd | +62/-0 | 7.96% |
| go.mod | +12/-0 | 1.54% |
```

| Directory | +/- | Δ % |
| --- | --- | --- |
| go.sum | +167/-0 | 21.44% |
| run | +142/-0 | 18.23% |
| algorithm | +107/-8 | 14.76% |
| input | +95/-11 | 13.61% |
| transform | +88/-1 | 11.42% |
| output | +86/-0 | 11.04% |
| cmd | +62/-0 | 7.96% |
| go.mod | +12/-0 | 1.54% |

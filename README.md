# gstats
cli tool to read a bunch of numbers from STDIN and dump statistical data about it

**for a more full featured tool, check out `ministat` from the FreeBSD project. Should be available in your package manager of choice.**

## usage

```
> cat examples/input.txt | gstats
total: 13
min: 1
mean: 6137.076923076923
max: 78131
p50: 29
p75: 85.5
p90: 671
p95: 39621
p99: 39621
```

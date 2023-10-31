# valyjson

## What is it

Yet another generator of the JSON-marshalers and unmarshalers code.
Powered by [valyala/fastjson](https://github.com/valyala/fastjson) this thing generates quite efficient parsing code.
And this is the reason of this name "valyjson".

## Why have I built it

The main reason is the experience.
But really I wanted to get a useful generator which I can integrate into my golang project that is used for building structures from specifications.
I've tried [mailru/easyjson](https://github.com/mailru/easyjson) and failed.
At first, I had difficulty setting up my generators that would invoke the mailru generators.
After several nervous attempts the generation was performing without errors.
Then I failed with the embedded structures, something seemed to break inside easyjson.

And I thought: Okay, if you don't want to work with me, I'm going to build my own json-parser-generator with blackjack...

## Main objectives

The goals I wanted to achieve were:

 - No dependencies: you do not need to import my package into your final product;
 - Efficiency: I used [valyala/fastjson](https://github.com/valyala/fastjson) for parsing which is very efficient.
 - Standards: I took care to match the behavior of marshalers/parsers with that of the standard golang parser.
 - Declarative: as few cmd-flags as possible, all generation settings are located in a special comment, that means the code describes itself.

The last point can make your code more understandable.
You can understand the json behavior of each model even if the generation is not yet complete.
For the same reason all fields must have json tags.

## Take a look at the benchmarks

The results of the tests presented here, I borrowed from easyjson.
So if you trust the objectiveness of their tests, then you can trust my tests too.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
```

### Unmarshaling

| lib      | json size | ns/op  |  MB/s |  B/op | allocs/op |
|:---------|:----------|--------|------:|------:|----------:|
| standard | regular   | 106184 | 122.7 | 10016 |       208 |
| standard | small     | 1859   |  44.1 |   688 |        16 |
|          |           |        |       |       |           |
| ffjson   | regular   | 37354  | 348.7 |  9885 |       141 |
| ffjson   | small     | 921    |  89.1 |   496 |        10 |
|          |           |        |       |       |           |
| codec    | regular   | 35267  | 369.1 | 12840 |       125 |
| codec    | small     | 455    | 180.2 |   144 |         1 |
|          |           |        |       |       |           |
| valyjson | regular   | 32251  | 403.9 |  6785 |       122 |
| valyjson | small     | 461    | 177.7 |    80 |         3 |
|          |           |        |       |       |           |
| easyjson | regular   | 26220  | 499.7 |  9512 |       126 |
| easyjson | small     | 414.6  | 197.9 |   128 |         3 |

Tests of unpacking JSON objects show a nice gain in memory allocation.
This is not surprising, I used [valyala/fastjson](https://github.com/valyala/fastjson) parser,
and Aliaksandr Valialkin knows how to save memory.

### Marshaling

For convenience, I cut out all the results except easyjson and left only those tests which work with concurrency.
In real life, we are unlikely to encounter a situation where we have no parallelism,
but still need the generation of marshaller code.

| lib      | json size | ns/op  |   MB/s |   B/op | allocs/op |
|:---------|:----------|--------|-------:|-------:|----------:|
| valyjson | large     | 117508 | 3802.4 | 459605 |        27 |
| valyjson | regular   | 3096   | 4206.4 |  10238 |         9 |
| valyjson | small     | 61.53  | 1316.4 |    128 |         1 |
|          |           |        |        |        |           |
| easyjson | large     | 101827 | 4393.8 | 466120 |        30 |
| easyjson | regular   | 2462   | 5290.9 |  10293 |         9 |
| easyjson | small     | 42.12  | 1923.2 |    128 |         1 |

Here I lost a bit in processing speed, but won in the number of requests to the memory allocator.
Not a great achievement, but I achieved my goals (see above).

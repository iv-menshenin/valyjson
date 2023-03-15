# valyjson

## What is it

Yet another generator of the JSON-marshalers and unmarshalers code.
Powered by [valyala/fastjson](https://github.com/valyala/fastjson) this thing generates quite efficient parsing code.
And this is the reason of this name "valyjson".

## Why I`m built it

The main reason is the experience.
But really I wanted any useful generator that I can integrate into my project that used golang for making structures from specifications.
I've tried [mailru/easyjson](https://github.com/mailru/easyjson) and had failed.
At first, I had difficulty setting up my generators that they would invoke the mailru generators.
After several nervous attempts the generation is performed without errors.
Then I failed with the embedded structures, something seemed to break there in the easyjson.

And I thought, Okay, if you don't want to work with me, I'm going to build my own json-parser-generator with blackjack...

## Main objectives

The main goals I wanted to achieve were:

 - No dependencies: you do not need to import my package into your final product;
 - Efficiency: I used [valyala/fastjson](https://github.com/valyala/fastjson) for parsing, so it's very efficient.
 - Standards: I took care to match the behavior of marshalers/parsers with that of the standard golang parser.
 - Declarative: as few cmd-flags as possible, all generation settings are located in a special comment, which means the code describes itself.

The last point can make your code more understandable.
You can understand the json behavior of each model even if the generation is not yet complete.
For the same reason all fields must have json tags.

```
goos: linux
goarch: amd64
pkg: github.com/mailru/easyjson/benchmark
cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
```

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
| valyjson | regular   | 30826  | 422.5 |  4523 |        56 |
| valyjson | small     | 469    | 174.9 |    72 |         3 |
|          |           |        |       |       |           |
| easyjson | regular   | 26220  | 499.7 |  9512 |       126 |
| easyjson | small     | 414.6  | 197.9 |   128 |         3 |



| lib      | json size | ns/op  |     MB/s |     B/op | allocs/op |
|:---------|:----------|--------|---------:|---------:|----------:|
| valyjson | large     | 320148 |  1382.67 |  2621070 |      1009 |
| valyjson | regular   | 8370   |   1556.1 |    60832 |        23 |
| valyjson | small     | 21.86  |   3705.3 |        0 |         0 |
|          |           |        |          |          |           |
| easyjson | large     | 91492  |   4890.1 |   464261 |        30 |
| easyjson | regular   | 1792   |   7266.8 |    10283 |         9 |
| easyjson | small     | 33.44  |   2421.9 |      128 |         1 |
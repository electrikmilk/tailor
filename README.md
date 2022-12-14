# Tailor

This is a work-in-progress CSS minifier and optimizer. The parser and resulting minifier are the most complete, while the optimization is in progress.

## Usage

```console
./tailor file.css
```

### Current features

 - Parse CSS files
 - Minify CSS files (average 25% file size decrease so far)
 - Compare property values to their initial value
 - Some syntax checking while and after parsing
 - Suggest alternatives to deprecated tags
 - Check if HTML tag selector is a valid HTML tag
 - Warn about redundant use of measurements (eg. 0px)

### Tests

| Test                       | Time  |
|----------------------------|-------|
| 7,271 lines with checks    | 0.70s |
| 7,271 lines without checks | 0.60s |

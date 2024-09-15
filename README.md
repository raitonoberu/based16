# based16

A dead simple [base16](https://github.com/chriskempson/base16) theming tool. It takes the yaml file and fills your template with the colors. The rest is up to you, choom.

```sh
$ echo "my background is #{{.base00}}" | ./based16 -t mocha.yaml
my background is #1e1e2e
```

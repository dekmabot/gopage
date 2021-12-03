
# A simple golang app example for AWS Lambda

Demo is here: https://r7z8q1ohb0.execute-api.eu-west-1.amazonaws.com/default/myFunc

It can:

1. Show simple page with .gohtml template.
2. Show current time in cyrillic locale.
3. Can be deployed to AWS Lambda by one console command.

Supports commands:
```bash
make build # build app into binary, including outer template files

make zip # makes zip-archive prepared to be deployed to AWS Lambda

make deploy # deploy prepared zip file to AWS Lambda

make push # Do all commands together step by step
```


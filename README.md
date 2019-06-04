# Show current time in terminal with ntp client

## How it works

Application is trying to access time.apple.com and show NTP time, if fail local time used (exit code=1)

## Usage example

```shell
export GO111MODULE=on
git clone git@github.com:imorph/ntp-now.git
cd ntp-now
go build
./ntp-now
```

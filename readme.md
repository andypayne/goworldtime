# A Simple Go Web Service


## Building

Build with `go build`

## Running

To listen on port 3000:
```zsh
goworldtime
```

To override the port:
```zsh
goworldtime --port=4040
```

## Using

Edit `test_data.json` and run the tester script to add new times:

```shell
./tester.sh
```

Or upload directly with curl:
```shell
curl -v -X POST \
  -d '{"hours": 0, "minutes": 0, "seconds": 0, "day": 13, "month": 2, "year": 2020, "tz": "America/Los_Angeles"}' \
  -H 'Content-Type: multipart/form-data' \
  http://localhost:4040/times
```

Run a few times with different values and then fetch `/times` to see the times:
[http://localhost:4040/times](http://localhost:4040/times)


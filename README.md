# kinesis-tailf

tail -f command for Amazon Kinesis.

## Install

```
go get -u github.com/fujiwara/kinesis-tailf/cmd/kinesis-tailf
```

## Usage

Required flags are below.

- `-stream`
- `-region` or `AWS_REGION` environment variable

```
Usage of kinesis-tailf:
  -lf
    	append LF(\n) to each record
  -region string
    	region
  -shard-id string
    	shard id (, separated) default: all shards in the stream.
  -stream string
    	stream name
  -timestamp string
    	start timestamp.
    	default LATEST
```

kinesis-tailf supports decoding packed records by Kinesis Producer Library (KPL).

## Licence

MIT


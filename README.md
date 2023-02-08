# fireflake

A unique ID generator in a distributed system over gRPC.
Used at [Ballot India](https://ballotindia.com)

`1 bit sign`
`33 bits time in seconds`
`10 bits machine ID (2^10 machines)`
`20 bits sequence 2^20`

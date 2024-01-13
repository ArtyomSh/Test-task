# Test task
executed by Artem Shmakov
## Installation & Run
```bash
# Download this project
git clone https://github.com/ArtyomSh/Test-task.git
```
```bash
# Build project
go build
```
If you chose `redis` as your RateRepository in the `config.yml`, you need to bring up the redis container before working with the built build. To do this, run the command:
```bash
# Run redis
docker run --name redis-test-instance -p 6379:6379 -d redis
```
## Project structure
```
.
├── cmd                     // Сonsole application commands
│   ├── rate.go
│   ├── root.go
│   └── server.go
├── configs
│   └── config.yml         // Configuration
├── go.mod
├── go.sum
├── internal
│   ├── configs
│   │   └── config.go    
│   ├── handlers          // API core handlers                
│   │   └── handler.go
│   ├── models
│   │   └── model.go
│   └── repositories      // 
│       ├── MemoryRateRepo.go
│       ├── RateRepo.go
│       └── RedisRateRepo.go
├── main.go
└── pkg
    ├── Ticker
    │   └── UpdateRate.go  // Ticker that updates Rate
    ├── loggers
    │   └── ZapLogger.go
    └── utils             // Helper code
        └── helpers.go
```
## Commands
* Server
```bash
# Run server
TestTask server
```
* Rate
```bash
# Get `ETH-USDT` rate
TestTask rate --pairs=ETH-USDT
```
```bash
# Get [BTC-USDT,ETH-USDT] rates
TestTask rate --pairs=BTC-USDT,ETH-USDT
```


# DC-Simulator-

![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)

**DC-Simulator** (Distributed Computing Simulator) is a command-line tool built in Go that illustrates the core concepts of distributed systems. By leveraging Go's powerful concurrency primitives—Goroutines, Channels, and WaitGroups—it simulates a cluster of computational nodes processing a large dataset in parallel.

This project implements a robust **Worker Pool architecture**, making it highly scalable and memory-efficient.

---

## Features

* **Worker Pool Pattern:** Dynamically spawns a fixed number of workers to process a queue of jobs, preventing memory overflow.
* **Concurrency in Action:** Hands-on demonstration of Goroutines, bidirectional Channels, and Mutex/WaitGroups.
* **Dynamic Workloads:** Fully customizable CLI parameters to test different loads and bottleneck scenarios.
* **Simulated Network Latency:** Built-in artificial delays to mimic real-world network packet transfers between the coordinator and nodes.

## Project Structure

Following the standard Go project layout:

```text
dc-simulator/
├── cmd/
│   └── simulator/
│       └── main.go         # Entry point and CLI parsing
└── internal/
    ├── model/
    │   └── task.go         # Data structures (Job, Result)
    ├── worker/
    │   ├── worker.go       # Core computation logic
    │   └── worker_test.go  # Unit tests for the worker node
    └── pool/
        └── pool.go         # Worker pool orchestration and map-reduce logic

```

## Getting Started

### Prerequisites

* Go 1.26 or higher installed on your machine.

### Installation

1. Clone the repository:
```bash
git clone [https://github.com/HairanRins/DC-Simulator-.git](https://github.com/HairanRins/DC-Simulator-.git)
cd DC-Simulator

```


2. Download dependencies (if any are added later) and verify the module:
```bash
go mod tidy

```



## Usage & Benchmarking

You can run the simulator directly using the Go CLI. The tool accepts three main flags to customize the simulation:

* `-workers`: Number of concurrent nodes (default: 4)
* `-size`: Total dataset size to compute (default: 100)
* `-chunk`: Size of the data packet sent to each worker at a time (default: 10)

### Use Cases to Test

**1. Default Scenario**
A balanced, lightweight run to ensure everything is working.

```bash
go run cmd/simulator/main.go

```

**2. Heavy Load (Optimized for Multi-core)**
Ideal scenario for distributed computing: large dataset distributed in significant chunks across multiple workers.

```bash
go run cmd/simulator/main.go -size=5000000 -workers=8 -chunk=10000

```

**3. Massive Concurrency (Stress Test)**
Stress-testing the Go scheduler by creating an enormous amount of tiny jobs and a massive pool of workers.

```bash
go run cmd/simulator/main.go -size=100000 -workers=100 -chunk=50

```

**4. The Bottleneck (Monolithic approach)**
Forces a single worker to do all the heavy lifting. Compare the execution time against the "Heavy Load" scenario to see the power of concurrency!

```bash
go run cmd/simulator/main.go -size=5000000 -workers=1 -chunk=10000

```

### Example Output

```text
 ____   ____       _____ _                 _       _             
|  _ \ / ___|     / ____(_)               | |     | |            
| | | | |   _____| (___  _ _ __ ___  _   _| | __ _| |_ ___  _ __ 
| | | | |  |______\___ \| | '_ ' _ \| | | | |/ _' | __/ _ \| '__|
| |_| | |___      ____) | | | | | | | |_| | | (_| | || (_) | |   
|____/ \____|    |_____/|_|_| |_| |_|\__,_|_|\__,_|\__\___/|_|   
                                                                 
          Distributed Computing Simulator in Go

▶ Parameters: Data = 100 | Workers = 4 | Chunk size = 10
------------------------------------------------------------------
[Network] Worker 2 finished Job 1 -> Local sum: 385
[Network] Worker 4 finished Job 2 -> Local sum: 2185
...
------------------------------------------------------------------
✔ Total sum of squares : 338350
⏱ Execution time       : 152.3ms

```

## Testing

The project includes Table-Driven unit tests to ensure the mathematical accuracy of the workers, even with edge cases like empty slices or negative numbers.

Run the tests using:

```bash
go test ./... -v

```

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://www.google.com/search?q=https://github.com/HairanRins/DC-Simulator/issues).

## License

This project is [MIT](https://choosealicense.com/licenses/mit/) licensed.

```

<FollowUp label="Do you want to add a Makefile to simplify the commands?" query="How can I create a Makefile to simplify the build, run, and test commands for this project?"/>

```
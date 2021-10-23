package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Task struct {
    work float64
}

func doWork(t Task) {
    time.Sleep(time.Duration(t.work * 1000) * time.Millisecond)
}

func runWorker(id int, w *bufio.Writer, tasks <-chan Task, stopChannel chan int) {
    for {
        select {
        case t := <-tasks:
            w.Write([]byte(
                fmt.Sprintf("worker:%d sleep:%.1f\n", id, t.work),
            ))
            w.Flush()
            doWork(t)
        default:
            stopChannel <- id
            return
        }
    }
}

func getInputFloat(reader bufio.Reader) (float64, error) {
    inp, err := reader.ReadString('\n')
    if err != nil {
        return float64(0), err
    }
    inp = strings.Trim(inp, "\n")
    return strconv.ParseFloat(inp, 64)
}

func Run(poolSize int) {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)

    tasks := make(chan Task, poolSize)
    workersOn := 0
    stopChannel := make(chan int)

    stopIdlingWorker := func() bool {
        select {
        case s := <-stopChannel:
            writer.Write([]byte(
                "worker:" + strconv.Itoa(s) + " stopping\n",
            )) 
            writer.Flush()
            workersOn--
            return false
        default:
            return true
        }
    }

    stopReaderChan := make(chan int, 1)

    go func() {
        for {
            value, err := getInputFloat(*reader)
            if err == nil {
                // spawn new worker if needed and schedule the task
                if workersOn < poolSize {
                    writer.Write([]byte("worker:" + strconv.Itoa(workersOn + 1) + " spawning\n"))
                    writer.Flush()
                    go runWorker(workersOn + 1, writer, tasks, stopChannel)
                    workersOn++
                }

                tasks <- Task{value}
            }

            // stops idling workers
            stopIdlingWorker()

            // stop goroutine if needed
            select {
            case <-stopReaderChan:
                return
            default:
                continue
            }
        }
    }()

    wg := &sync.WaitGroup{}
    stopSig := make(chan int, 1)
    go func() {
        wg.Add(1)
        defer wg.Done()

        sigc := make(chan os.Signal, 1)
        signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)
        // check on sigterm, sigint signals
        select {
        case <-sigc:
            os.Stdin.Close()
            stopReaderChan <- 0
            for {
                if stopIdlingWorker() {
                    break
                }
            }
            return
        case <-stopSig:
            return
        }
    }()

    wg.Wait()
}

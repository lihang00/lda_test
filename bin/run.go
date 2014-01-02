package main

import(
    "lda_test"
    "fmt"
    //"runtime/pprof"
    //"os"
    //"log"
)

func main() {
    params := lda.PrepareParams()

    if params.Input == "" || params.Output == "" {
        fmt.Println("Missing input or output")
        return
    }

    if params.Method == "est" {
        estimator := lda.NewEstimator(params)
        estimator.Train()
    }
}
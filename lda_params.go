package lda

import(
    "flag"
    "fmt"
)

type LDAParams struct{
    Input string
    Output string
    Alpha float64
    Beta float64
    Ntopics int64
    Niters int64
    Method string
}

func PrepareParams() *LDAParams {
    params := LDAParams{}
    input_path := flag.String("input", "", "Input path, seprated by comma")
    output_path := flag.String("output", "", "Output path")
    alpha := flag.Float64("alpha", 0.1, "alpha")
    beta := flag.Float64("beta", 0.1, "beta")
    ntopics := flag.Int64("ntopics", 100, "Number of topics")
    niters := flag.Int64("niters", 20, "the number of iterations")
    method := flag.String("method", "", "mathod")

    flag.Parse()
    fmt.Println("Input path : ", *input_path)
    fmt.Println("Output path : ", *output_path)

    params.Input = *input_path
    params.Output = *output_path
    params.Alpha = *alpha
    params.Beta = *beta
    params.Ntopics = *ntopics
    params.Niters = *niters
    params.Method = *method
    return &params
}
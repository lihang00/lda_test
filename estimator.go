package lda

import(
    "bufio"
    "os"
    "strings"
    "fmt"
    "math/rand"
)

type Estimator struct {
    Model *LDAModel
    Param *LDAParams
    Data []*Document
}

func NewEstimator(param *LDAParams) *Estimator {
    e := Estimator{}
    e.Param = param
    e.Model = InitModel(param)
    e.Data = []*Document{}
    e.LoadData(e.Param.Input)
    return &e
}

func (e *Estimator) LoadData(fileName string) error {
    file, err := os.Open(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        tks := strings.Split(line, "\t");
        if len(tks) < 2  { continue }
        doc := NewDoc(tks[1], e.Model.word_dic)
        doc.Id = tks[0]
        e.Data = append(e.Data, doc)
    }

    return nil
}

func (e *Estimator) Train() {
    // init model
    z := []*IntVector{} // topic assignments for words, size M(number of docs) x doc.size() (word count in a doc)
    nd := []*IntVector{} //nd[i][j]: number of words in document i assigned to topic j, size M x K
    nwsum := NewIntVector() //nwsum[j]: total number of words assigned to topic j, size K
    ndsum := NewIntVector() //ndsum[i]: total number of words in document i, size M

    for m := int64(0); m < int64(len(e.Data)); m++ {
        doc := e.Data[m]
        z = append(z, NewIntVector())
        nd = append(nd, NewIntVector())

        //initilize for z
        for n := int64(0); n < int64(len(doc.Words)); n++ {
            topic := int64(rand.Float64() * float64(e.Param.Ntopics))
            z[m].AddValue(n, topic)

            // number of instances of word assigned to topic j
            e.Model.IncreaseNWZ(doc.Words[n], topic)
                
            // number of words in document i assigned to topic j
            nd[m].AddValue(topic, 1);
            // total number of words assigned to topic j
            nwsum.AddValue(topic, 1);
        }

        // total number of words in document i
        ndsum.AddValue(m, int64(len(doc.Words)))
    }


    // train lda model
    for iter := int64(0); iter < e.Param.Niters; iter++ {
        fmt.Println("Iteration ", iter, " ...")
            
        // for all z_i


    }



}


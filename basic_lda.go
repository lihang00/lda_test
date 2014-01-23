package lda

import(
    "bufio"
    "os"
    "strings"
    "fmt"
    "math/rand"
)

type BasicLDA struct {
    Model *LDAModel
    Param *LDAParams
    Data []*Document
}

func NewBasicLDA(param *LDAParams) *BasicLDA {
    e := BasicLDA{}
    e.Param = param
    if e.Param.Method != "inference" {
        e.Model = InitModel(e.Param)
    }else if e.Param.Method == "inference" {
        e.Model.LoadModel(e.Param.ModelPath)
    }
    e.Data = []*Document{}
    e.LoadData(e.Param.Input)
    return &e
}

func (e *BasicLDA) LoadData(fileName string) error {
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

func (e *BasicLDA) Train() {
    // init model
    z := []*IntVector{} // topic assignments for words, size M(number of docs) x doc.size() (word count in a doc)
    nd := []*IntVector{} //nd[i][j]: number of words in document i assigned to topic j, size M x K
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
            e.Model.IncreaseNZ(topic)
        }

        // total number of words in document i
        ndsum.AddValue(m, int64(len(doc.Words)))
    }


    // train lda model
    for iter := int64(0); iter < e.Param.Niters; iter++ {
        fmt.Println("Iteration ", iter, " ...")
            
        // for all z_i        
        for m := int64(0); m < int64(len(e.Data)); m++ {
            doc := e.Data[m]
            for n := int64(0); n < int64(len(doc.Words)); n++ {
                // z_i = z[m][n]
                topic := z[m].GetValue(n)
                word := doc.Words[n]

                // before sample
                e.Model.DecreaseNWZ(word, topic)
                e.Model.DecreaseNZ(topic)
                nd[m].AddValue(topic, -1)
                ndsum.AddValue(m, -1)

                // sample from p(z_i|z_-i, w)
                pr := e.CalSamplePr(word, nd[m], ndsum.GetValue(m))
                new_topic := pr.Sample()

                // add newly estimated z_i to count variables
                e.Model.IncreaseNWZ(word, new_topic)
                e.Model.IncreaseNZ(topic)
                nd[m].AddValue(topic, 1)
                ndsum.AddValue(m, 1)

                z[m].SetValue(n, topic)
            }// end for each word
        }// end for each document
    }// end iterations  

    fmt.Println("Gibbs sampling completed!");
    fmt.Println("Saving the final model!");

    // calcPerplexity

    // save model
    e.Model.SaveModel(e.Param.Output)
}

func (e *BasicLDA) CalSamplePr(word int64, ndm *IntVector, ndsumm int64) *Prob{
    pr := NewProb()
    Vbeta := e.Model.beta * float64(e.Model.GetWordCount())
    Kalpha := e.Model.alpha * float64(e.Model.ntopics)
    for k := int64(0); k < e.Model.ntopics ; k++ {
        pr_k := (float64(e.Model.GetNWZElement(word, k)) + e.Model.beta) / (float64(e.Model.GetNZElement(k)) + Vbeta) * 
                (float64(ndm.GetValue(k)) + e.Model.alpha) / (float64(ndsumm) + Kalpha)
        pr.AddValue(k, pr_k)
    }

    return pr

}
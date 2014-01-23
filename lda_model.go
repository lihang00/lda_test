package lda

import (
    "strconv"
    "os"
    "bufio"
    "strings"
    "hector/util"
)

type LDAModel struct {
    ntopics int64 // number of topics K 
    alpha float64 // LDA  hyperparameters of the Dirichlet prior on the per-document topic distributions
    beta float64  // LDA  hyperparameters of the Dirichlet prior on the per-topic word distribution

    nwz []*IntVector // nwz[i][j] number of instances of word/term i assigned to topic j, size V x K
    nz []int64 // nz[] total number of words assigned to topic j, size K
    word_dic *WordDic
}

func InitModel(param *LDAParams) *LDAModel {
    m := LDAModel{}
    m.ntopics = param.Ntopics
    m.alpha = param.Alpha
    m.beta = param.Beta
    m.nwz = []*IntVector{}
    m.nz = make([]int64, m.ntopics)
    for i := int64(0); i < m.ntopics; i += 1 { m.nz[i] = 0 }
    m.word_dic = NewWordDic()

    return &m
}

func (m *LDAModel) AddNWZ(word int64, topic int64, freq int64) {
    for int64(len(m.nwz)) <= word {
        m.nwz = append(m.nwz, NewIntVector())
    }

    m.nwz[word].AddValue(topic, freq)
}

func (m *LDAModel) IncreaseNWZ(word int64, topic int64) {
    m.AddNWZ(word, topic, int64(1))
}

func (m *LDAModel) DecreaseNWZ(word int64, topic int64) {
    m.AddNWZ(word, topic, int64(-1))
}

func (m *LDAModel) AddNZ(topic int64, cnt int64) {
    m.nz[topic] += cnt
}

func (m *LDAModel) IncreaseNZ(topic int64) {
    m.AddNZ(topic, int64(1))
}

func (m *LDAModel) DecreaseNZ(topic int64) {
    m.AddNZ(topic, int64(-1))
}

func (m *LDAModel) GetNWZElement(word int64, topic int64) int64{
    return m.nwz[word].GetValue(topic)
}

func (m *LDAModel) GetNZElement(topic int64) int64{
    return m.nz[topic]
}

func (m *LDAModel) GetWordCount() int64 {
    return int64(len(m.nwz))
}

func (m *LDAModel) SaveModel(path string) string{
    sb := util.StringBuilder{}
    for n := int64(0); n < int64(len(m.nwz)) ; n++ {
        for k, cnt := range m.nwz[n].data {
            sb.Write(m.word_dic.GetWord(n))
            sb.Write("\t")
            sb.Int64(k)
            sb.Write("\t")
            sb.Int64(cnt)
            sb.Write("\n")
        }
    }

    if path != "" {
        sb.WriteToFile(path)
    }

    return sb.String()
}

func (m *LDAModel) LoadModel(path string) {
    file, _ := os.Open(path)
    defer file.Close()

    scaner := bufio.NewScanner(file)
    for scaner.Scan() {
        line := scaner.Text()
        m.LoadModelLine(line)
    }
}

func (m *LDAModel) LoadModelLine(line string) {
    tks := strings.Split(line, "\t")
    if len(tks) != 3 {
        return
    }
    word := m.word_dic.AddWord(tks[0])
    topic, _ := strconv.ParseInt(tks[1], 10, 64)
    freq, _ := strconv.ParseInt(tks[2], 10, 64)

    for int64(len(m.nwz)) <= word {
        m.nwz = append(m.nwz, NewIntVector())
    }  

    m.nwz[word].AddValue(topic, freq)   
    m.nz[topic] += freq
}

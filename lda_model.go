package lda

import (
    //"math/rand"
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

func (m *LDAModel) IncreaseNWZ(word int64, topic int64) {
    for int64(len(m.nwz)) <= word {
        m.nwz = append(m.nwz, NewIntVector())
    }

    m.nwz[word].AddValue(topic, int64(1))
}
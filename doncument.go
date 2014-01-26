package lda

import (
    "strings"
)

type Document struct {
    Id string
    Words []int64
}

func NewEmptyDoc() *Document {
    d := Document{}
    d.Words = []int64{}
    return &d
}

func NewDoc(line string, wordDic *WordDic) *Document {
    d := Document{}
    d.Words = []int64{}

    tks := strings.Split(line, " ")
    for _, tk := range tks {
         if len(tk) == 0 {
            continue
         }
        id := wordDic.AddWord(tk)
        d.Words = append(d.Words, id)
    }

    return &d
}

func (d *Document) ReadLine(line string, wordDic *WordDic) {
    tks := strings.Split(line, " ")
    for _, tk := range tks {
         if len(tk) == 0 {
            continue
        }
        id := wordDic.GetId(tk)
        if id >= 0 {
            d.Words = append(d.Words, id)
        }
    }    
}
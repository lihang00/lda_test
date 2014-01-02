package lda

type WordDic struct {
    Word2Id map[string]int64
    Id2Word []string
}

func NewWordDic() *WordDic {
    dic := WordDic{}
    dic.Word2Id = make(map[string]int64)
    dic.Id2Word = []string{}
    return &dic
}

func (dic *WordDic) GetWord(id int64) string {
    if id < int64(len(dic.Id2Word)) {
        return dic.Id2Word[id]
    }

    return ""
}

func (dic *WordDic) GetId(word string) int64 {
    id, ok := dic.Word2Id[word]
    if !ok {
        return -1
    } else {
        return id
    }
}

func (dic *WordDic) AddWord(word string) int64 {
    var id = dic.GetId(word)
    if id >= 0 {
        return id
    }

    id = int64(len(dic.Id2Word))
    dic.Word2Id[word] = id
    dic.Id2Word = append(dic.Id2Word, word)

    return id
}
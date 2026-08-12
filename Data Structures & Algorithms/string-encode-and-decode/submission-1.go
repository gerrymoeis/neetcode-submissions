const DIFF uint16 = 256

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, str := range strs {
        res += str
        res += string(DIFF)
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    str := ""
    for _, c := range encoded {
        if uint16(c) == DIFF {
            res = append(res, str)
            str = ""
            continue
        }
        str += string(c)
    }
    return res
}
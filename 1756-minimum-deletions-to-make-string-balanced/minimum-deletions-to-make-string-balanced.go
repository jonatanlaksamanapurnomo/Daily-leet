//aababbab
//
func minimumDeletions(s string) int {
    bLeft := 0 
    aRight := 0
    resp := math.MaxInt32
    //Pass 1 : cari total A
    for i := 0 ; i<len(s) ; i++{
        if s[i] == 'a'{
            aRight++
        }
    }

    //Pass 2: cari total b sesuai dgn index
    for i := 0 ; i<len(s) ; i++{
        resp = min(resp, bLeft+aRight)
        if s[i] == 'b'{
            bLeft++
        }else{
            aRight--
        }
    }

    resp = min(resp,bLeft+aRight)

    return resp
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
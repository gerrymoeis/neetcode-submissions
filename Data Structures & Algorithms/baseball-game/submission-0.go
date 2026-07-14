func calPoints(operations []string) int {
    total := 0
    records := make([]int, len(operations))
    i := 0
    for _, op := range operations {
        if op == "+" {
            records[i] = records[i-1] + records[i-2]
        } else if op == "D" {
            records[i] = records[i-1]*2
        } else if op == "C" {
            i--
            total -= records[i]
            records[i] = 0
            continue
        } else {
            records[i], _ = strconv.Atoi(op)
        }
        total += records[i]
        i++
    }
    return total
}

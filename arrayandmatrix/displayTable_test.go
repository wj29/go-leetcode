package arrayandmatrix

import (
    "sort"
    "strconv"
    "testing"
)

func Test_DisplayTable(t *testing.T) {
    orders := [][]string{
        {"David", "3", "Ceviche"},
        {"Corina", "10", "Beef Burrito"},
        {"David", "3", "Fried Chicken"},
    }
    t.Log(displayTable(orders))
}

func displayTable(orders [][]string) [][]string {
    foodMap := make(map[string]struct{})
    tableMap := make(map[int]struct{})
    valMap := make(map[int]map[string]int)

    for _, order := range orders {
        tableTmp, _ := strconv.Atoi(order[1])
        if _, ok := tableMap[tableTmp]; !ok {
            tableMap[tableTmp] = struct{}{}
        }
        if _, ok := foodMap[order[2]]; !ok {
            foodMap[order[2]] = struct{}{}
        }
        if _, ok := valMap[tableTmp]; !ok {
            valMap[tableTmp] = make(map[string]int)
        }
        valMap[tableTmp][order[2]]++

    }
    ret := make([][]string, len(tableMap)+1)
    ret[0] = append(ret[0], "Table")
    for food := range foodMap {
        ret[0] = append(ret[0], food)
    }
    sort.Strings(ret[0][1:])
    tables := make([]int, 0)
    for table := range tableMap {
        tables = append(tables, table)
    }
    sort.Ints(tables)
    for i := 1; i < len(tables)+1; i++ {
        ret[i] = append(ret[i], strconv.Itoa(tables[i-1]))
        for j := 1; j < len(foodMap)+1; j++ {
            ret[i] = append(ret[i], strconv.Itoa(valMap[tables[i-1]][ret[0][j]]))
        }
    }
    return ret
}

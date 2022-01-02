package main

//func main() {
//	Map := make(map[int]int)
//
//	for i := 0; i < 30; i++ {
//		go writeMap(Map, i, i) // concurrent map writes
//		//go readMap(Map, i)
//	}
//
//}

func readMap(Map map[int]int, key int) int {
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	Map[key] = value
}

package aboutSet

import (
	"log"
	"testing"
)

func TestUnion(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	un := union(slice1, slice2)
	log.Printf("slice1与slice2的并集为:%v", un)
}
func TestIntersect(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	in := intersect(slice1, slice2)
	log.Printf("slice1与slice2的交集为:%v", in)
}
func TestDifference(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	di := difference(slice1, slice2)
	log.Printf("slice1与slice2的差集为:%v", di)
}

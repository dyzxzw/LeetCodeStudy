package Set

import "testing"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-07-21 20:44
 **/
func TestSet(t *testing.T) {
	set:=NewSet()
	for i:=0;i<=100;i+=10{
		set.Add(i)
	}

	t.Log(set.SetPrint())

}
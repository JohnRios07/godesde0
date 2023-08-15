package variables

import "fmt"

func MuestroEnteros() {
	var intComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intComun = ", intComun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}

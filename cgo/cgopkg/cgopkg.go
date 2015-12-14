package cgopkg

/*
void cGoCall();
void cCall();
*/
import "C"

var count int

//export GoCall
func GoCall() {
	count += 1
}

func GoBenchCall() {
	C.cGoCall()
}

func CBenchCall() {
	C.cCall()
}

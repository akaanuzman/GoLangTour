package main

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct{}
type YEncoder struct{}

func (x *XEncoder) Encode(value string) {
	println("XEncoder is encoding")
}

func (x *XEncoder) Decode(value string) {
	println("XEncoder is decoding")
}

func (x *YEncoder) Encode(value string) {
	println("YEncoder is encoding")
}

func (x *YEncoder) Decode(value string) {
	println("YEncoder is decoding")
}

func main() {
	var encoder IEncoder = &YEncoder{}
	encoder.Encode("Hello")
	encoder.Decode("Hello")
}

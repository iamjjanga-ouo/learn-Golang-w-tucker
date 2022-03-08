package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {}

func (f *File)Read()  {
	
}

func ReadFile(reader Reader)  {
	// Close에 대한 메서드 구현이 안되었기 때문에 인터페이스 변환은 runtime error를 발생시킨다.
	// 타입변환 성공 여부 반환을 체크
	if c, ok := reader.(Closer); ok {
		c.Close()	
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}
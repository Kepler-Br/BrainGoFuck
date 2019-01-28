package main
import (
	"fmt"
	"os"
)
import "io/ioutil"
import "errors"


type BrainGoFuck struct {
	source string
	memory []uint8
	memoryCarriage int
	sourceCarriage int
	loopStack []int
}

func (this BrainGoFuck) PrintSource() {
	println(this.source)
}

func (this *BrainGoFuck) Step() {

}

func (this *BrainGoFuck) RunString(bfCode string) {
	this.source = bfCode
	for ; this.sourceCarriage < len(this.source); {
		this.step(string(this.source[this.sourceCarriage]))
	}
}

func (this *BrainGoFuck) nextCell() {
	if this.memoryCarriage == len(this.memory)-1 {
		this.memoryCarriage = 0
	} else {
		this.memoryCarriage++
	}
}

func (this *BrainGoFuck) prevCell() {
	if this.memoryCarriage == 0 {
		this.memoryCarriage = len(this.memory)
	} else {
		this.memoryCarriage--
	}
}

func (this *BrainGoFuck) incrementCell() {
	this.memory[this.memoryCarriage]++
}

func (this *BrainGoFuck) deincrementCell() {
	this.memory[this.memoryCarriage]--
}

func (this *BrainGoFuck) printCell(){
	print(string(this.memory[this.memoryCarriage]))
}

func (this *BrainGoFuck) readInput() {
	var ololo string = ""
	fmt.Scan(&ololo)
	this.memory[this.memoryCarriage] = ololo[0]
}

func (this *BrainGoFuck) beginLoop() {
	if this.memory[this.memoryCarriage] != 0 {
        this.loopStack = append(this.loopStack, this.sourceCarriage)
        
        this.sourceCarriage++
        return
    }
    loops:= -1
	for i := this.sourceCarriage; ; i++ {
		if i >=len(this.source) {
			panic("OW FUCK! Not closed loop")
		}
		if string(this.source[i]) == "]"{
			loops--
		}
		if string(this.source[i]) == "["{
			loops++
		}
		if string(this.source[i]) == "]"  && loops == 0{
            if this.loopStack[len(this.loopStack)-1] == this.sourceCarriage {
                this.loopStack = this.loopStack[:len(this.loopStack)-1]
            }
            this.sourceCarriage = i+1
			break
		}

	}
	
}

func (this *BrainGoFuck) endLoop() {
	if this.memory[this.memoryCarriage] == 0 {
		if len(this.loopStack) == 0 {
			panic("OW FUCK. STACK UNDERFLOW.")
		}
		this.sourceCarriage++
		this.loopStack = this.loopStack[:len(this.loopStack)-1]
	} else {
		this.sourceCarriage = this.loopStack[len(this.loopStack)-1]+1
	}
}

func (this *BrainGoFuck) step(symbol string) {
	switch(symbol) {
		case ">":
			this.nextCell()
			break
		case "<":
			this.prevCell()
			break
		case ".":
			this.printCell()
			break
		case ",":
			this.readInput()
			break
		case "+":
			this.incrementCell()
			break
		case "-":
			this.deincrementCell()
			break
		case "[":
			this.beginLoop()
			return
		case "]":
			this.endLoop()
			return
	}
	this.sourceCarriage++
	
}

func NewBrainGoFuck () BrainGoFuck {
	var source string = ""
	var memoryCells uint = 30000
	var memory []uint8 = make([]uint8, memoryCells)
	var memoryCarriage int = 0
	var sourceCarriage int = 0
	var loopStack []int = make([]int, 0)
	return BrainGoFuck{source, memory, memoryCarriage, sourceCarriage, loopStack}
}

func ReadFile(fileName string) (string, error) {
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", errors.New("cannot open file")
	}
	return string(code), nil
}

func main() {
	if len(os.Args) == 1 {
		println("No file provided.")
		os.Exit(-1)
	}

	var filename = os.Args[1]
	bfCode, err := ReadFile(filename)
	if err != nil {
		println("Cannot open file.")
		os.Exit(-1)
	}
	var bfInterpreter = NewBrainGoFuck()
	bfInterpreter.RunString(bfCode)
}

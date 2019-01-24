package main
import "fmt"
import "io/ioutil"
import "errors"


type BrainGoFuck struct {
	source string
	memory []uint8
	memoryCarriage uint
	sourceCarriage uint
	loopStack []uint
}

func (this *BrainGoFuck) ReadFile(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
    if err != nil {
        return errors.New("Cannot open file")
    }
    this.source = string(b)
    return nil
}

func (this BrainGoFuck) PrintSource() {
	println(this.source)
}

func (this *BrainGoFuck) Step() {

}

func (this *BrainGoFuck) Run() {
	for ; this.sourceCarriage < uint(len(this.source)); {
		this.interpretate(string(this.source[this.sourceCarriage]))
//         print(this.source[this.sourceCarriage-1])
//         print(" ")
//         println(this.sourceCarriage)
	}
// 	this.PrintSource()
}

func (this *BrainGoFuck) nextCell() {
	if this.memoryCarriage == uint(len(this.memory)-1) {
		this.memoryCarriage = 0
	} else {
		this.memoryCarriage++
	}
}

func (this *BrainGoFuck) prevCell() {
	if this.memoryCarriage == 0 {
		this.memoryCarriage = uint(len(this.memory)-1)
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
		if i >= uint(len(this.source)) {
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
		this.sourceCarriage = this.loopStack[len(this.loopStack)-1]
		this.loopStack = this.loopStack[:len(this.loopStack)-1]
	} else {
		this.sourceCarriage = this.loopStack[len(this.loopStack)-1]
	}
}

func (this *BrainGoFuck) interpretate(symbol string) {
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
		case "-":
			this.incrementCell()
			break
		case "+":
			this.deincrementCell()
			break
// 		case "[":
// 			this.beginLoop()
// 			break
// 		case "]":
// 			this.endLoop()
// 			break
	}
	this.sourceCarriage++
	
}

func NewBrainGoFuck () BrainGoFuck {
	var source string = ""
	var memoryCells uint = 30000
	var memory []uint8 = make([]uint8, memoryCells)
	var memoryCarriage uint = 0
	var sourceCarriage uint = 0
	var loopStack []uint = make([]uint, 0)
	return BrainGoFuck{source, memory, memoryCarriage, sourceCarriage, loopStack}
}

func loopFinder(src string) {
	var beginLoopPointer uint = 0
	var endLoopPointer uint = 0
	endLoopPointer = endLoopPointer
	loops := -1
	for i := beginLoopPointer; ; i++ {
		if i >= uint(len(src)) {
			panic("OW FUCK! Not closed loop")
		}
		if string(src[i]) == "]"  && loops == 0{
			endLoopPointer = i
			break
		}
		if string(src[i]) == "]"{
			loops--
		}
		if string(src[i]) == "["{
			loops++
		}
	}
	println(beginLoopPointer, endLoopPointer)
}


func main() {
	var interpretator BrainGoFuck = NewBrainGoFuck() 
	interpretator.ReadFile("hello.bf")
	interpretator.PrintSource()
	interpretator.Run()
	// loopFinder("[[++]++]");
}

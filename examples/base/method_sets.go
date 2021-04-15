type Num int

func (num Num) A1()        { num++ }
func (num Num) B1(val Num) { num += val }

func (num *Num) A2()        { *num++ }
func (num *Num) B2(val Num) { *num += val }

type Num int

func (num Num) A1()        { num++ }
func (num Num) B1(val int) { num += val }

func (num *Num) A2()        { *num++ }
func (num *Num) B2(val int) { *num += val }
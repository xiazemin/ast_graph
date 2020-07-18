package dot

func getDotHead() string {
	return `digraph g {
    nodesep = .5;
    rankdir = LR;    //指定绘图的方向 (LR从左到右绘制)
   `
}

func getDotTail() string {
	return "\n}"
}

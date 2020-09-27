package tree

/*
	B+树也是一种用于查找的多路查找树，常用于数据库和操作系统的文件系统中

	相同阶数的B+树与B树的区别有(假设阶数为m):
	1. 在B+树中，具有n个关键字的结点只含有n棵子树，即每个关键字对应一棵子树，而不像B树中，含有n个关键字的结点有n+1棵子树

	2. 在B+树中，每个结点(非根非叶子)关键字个数n的范围为 math.Ceil(m/2) <= n <= m(根结点1<=n<=m)，而在B树中，每个结点(非根非叶子)关键字
	   个数n的范围为 math.Ceil(m/2)-1 <= n <= m-1(根结点1<=n<=m-1)

	3. 在B+树中，信息(即数据)存储在叶子结点中，而所有非叶子结点仅起到索引作用，非叶子结点中的每个索引项只含有对应子树的最大关键字和指向该子树的
	   指针，不含有关键字对应记录的存储地址

	4. 在B+树中，叶子结点包含了全部关键字，即在非叶子结点中出现的关键字也会出现在叶子结点中，而在B树中，每个结点都包含关键字且每个结点的关键字
	   不重复

	5. 在B+树中，有一个指针指向关键字最小的叶子结点，并且所有叶子结点链接成一个单链表
*/

//TODO code

//叶子结点中的记录
type Record struct {
	Key   int         //关键字
	Value interface{} //记录
}

//非终端结点中的关键字信息
type Key struct {
	Key   int        //关键字
	Child *BPlusNode //每个关键字都对应一个子树(如果存在子树的话则子树和关键字为一一对应的关系)
}

//B+树
type BPlusTree struct {
	M int //阶数
}

//结点
type BPlusNode struct {

}

type Leaf struct {
	Next *Leaf      //指向下一个相邻的叶子结点
	Data []*KeyData //叶子结点中包含记录(数据)
}
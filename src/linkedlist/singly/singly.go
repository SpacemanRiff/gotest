package singly

type Node struct{
	val int
	next *Node
}

func New_List(init_val int) Node {
	return Node{val: init_val, next: nil}
}

func Get_Val_At(node Node, des_index int) int {
	return val_at(node, des_index, 0)
}

func val_at(node Node, des_index int, cur_index int) int {
	if cur_index == des_index {
		return node.val
	} else if node.next == nil {
		return -1 //to-do: throw error
	} else {
		next := node.next
		return val_at(*next, des_index, cur_index+1)
	}
}

func Add_Node(node Node, new_val int) Node {
	last := go_to_last(node)
	next := Node{val: new_val, next: nil}
	last.next = &next
	return node
}

func go_to_last(node Node) Node {
	if(node.next != nil){
		next := node.next
		return go_to_last(*next)
	} else {
		return node
	}
}

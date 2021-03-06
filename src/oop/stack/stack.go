package stack

type Element struct{
	val int
	next *Element
}

func New_Stack(top_val int) (Element){
	return Element{val: top_val, next: nil}
}

func Push(head Element, new_val int) (Element){
	return Element{val: new_val, next: &head}
}

func Peak(head Element) int{
	return head.val
}

func Pop(head Element) (int, Element){
	return head.val, Element{val: head.next.val, next: head.next.next}
}

func Has_Next(head Element) bool{
	return head.next != nil
}

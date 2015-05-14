//the point is to see how Go handles generic types. They use interface{} which is loosely baesd on Object in Java.
//differs from functional programming languages where Go can run into subtyping problems.
// need to come up with an example

type Hashable interface {
  Hash() []byte
}

func printHash(item Hashable) {
   fmt.Println(item.Hash())
}


type LinkedList struct {
   value interface{} //interface{} is what we call the "top type", meaning that all other types are subtypes of interface{}. This is roughly equivalent to Object in Java. Subtyping of sorts, but argued to not exist in Go.
   next *LinkedList
}

func (oldNode *LinkedList) prepend(value interface{}) *LinkedList {
   return &LinkedList{value, oldNode}
}

func tail(value interface{}) *LinkedList {
   return &LinkedList{value, nil}
}

func traverse(ll *LinkedList) {
   if ll == nil {
     return
   }
   fmt.Println(ll.value)
   traverse(ll.next)
}

func main() {
   node := tail(5).prepend(6).prepend(7)
   node := tail(5).prepend("Hello").prepend([]byte{1,2,3,4}) //which will just crash the program
   //Because generic data structures in Go don't know the type of their values, the Go compiler won't correct the programmer, and your program will simply crash when you try to down-cast from interface{}.
   traverse(node)
}


//Go does not support operator overloading or keyword extensibility.

//:= looks at the return type of the function and sets that variable to be that type

//Go uses heap a lot

//http://www.golang-book.com/4/index.htm

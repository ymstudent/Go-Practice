设链表长度为k，快指针先走n步，然后快慢指针一起走，
当快指针走到终点的时候，慢指针走了k-n，这个节点就是
链表的倒数第n个节点  
eg:  
[1,2,3,4,5], n = 2  
快指针走了2步后:  
slow = 1, fast = 3  
然后快慢一起走:  
slow = 2, fast = 4  
slow = 3, fast = 5  
slow = 4, fast = nil  
因为我们要删除的就是倒数第n个节点，所以需要慢指针在
前一个节点停下，这就需要快指针只走到最后一个节点是就
停止，不能走到nil。因此for循环的终止条件为fast.Next ！= nil

ps：注意，如果初始化时 slow, fast = head, head
当测试用例为[1]时，slow.Next = slow.Next.Next 
会报空指针错误，所以需要用一个虚拟头节点来避免这种情况
修改：slow, fast := dummy, dummy


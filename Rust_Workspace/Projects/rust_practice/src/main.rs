#[derive(Debug)]
struct Node {
    value: i32,
    next_node: Option<Box<Node>>,
}

struct LinkedList {
    head: Option<Box<Node>>,
}

impl LinkedList {
    fn new() -> Self {
        LinkedList { head: None }
    }

    fn add_node(&mut self, value: i32) {
        let new_node = Box::new( Node {
            value,
            next_node: self.head.take(),
        });
        self.head = Some(new_node);
    }

    fn print(&self) {
        let mut current = &self.head;
        while let Some(node) = current {
            print!("{} -> ", node.value);
            current = &node.next_node;
        }
        print!("None");
    }


}

fn main() {

    let mut list = LinkedList::new();
    list.add_node(1);
    list.add_node(2);
    list.add_node(3);
    list.print();

}
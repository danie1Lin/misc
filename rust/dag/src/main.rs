use std::borrow::Borrow;
use std::cell::RefCell;
use std::rc::Rc;

type SomeMutRc<T> = Option<Rc<RefCell<T>>>;

#[derive(Debug)]
struct Node {
    value: usize,
    next: SomeMutRc<Node>,
}

impl Node {
    fn new(v: usize) -> Self {
        Self {
            value: v,
            next: None,
        }
    }

    fn set_next(&mut self, n: SomeMutRc<Node>) {
        self.next = n
    }

    fn get_next(&self) -> SomeMutRc<Node> {
        self.next.as_ref().map(|v| v.clone())
    }
}
fn main() {
    let mut node = Node::new(1);
    let mut node2 = Node::new(2);
    let node3 = Node::new(3);
    node2.set_next(Some(Rc::new(RefCell::new(node3))));
    node.set_next(node.get_next());
    node.set_next(None);
    println!("node 1 {:?}, node 2 {:?}", node, node2);

    let node5 = node2.get_next().unwrap(); // if only use Rc<Node> is a readonly reference
    let mut n5 = node5.borrow_mut();
    n5.set_next(node.get_next());
    let bnode5: &Rc<RefCell<Node>> = node5.borrow();
    println!("{:?}", bnode5); // why this will not fail?
    println!("{:?}", node5.borrow() as &Rc<RefCell<Node>>);

    let a = RefCell::new(1);
    *(a.borrow_mut()) += 1; // scope in the equation
                            // comment out to try
                            // let mutA = a.borrow_mut();
                            // *(mutA) += 1;
    println!("a: {:?}", a.borrow());
}

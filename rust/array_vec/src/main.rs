#![feature(const_maybe_uninit_assume_init)]
#![feature(const_fn_trait_bound)]
use std::{
    fmt::{self, write, Display},
    mem::MaybeUninit,
};

#[derive(Debug)]
pub struct ArrayVec<T, const N: usize> {
    items: [MaybeUninit<T>; N],
    length: usize,
}

impl<T, const N: usize> ArrayVec<T, { N }>
where
    T: Copy,
{
    pub const fn new() -> ArrayVec<T, { N }> {
        let m: [MaybeUninit<T>; N] = unsafe { MaybeUninit::uninit().assume_init() };
        ArrayVec {
            items: m,
            length: 0,
        }
    }

    pub fn push(&mut self, a: T) {
        self.items[self.length] = MaybeUninit::new(a);
        self.length += 1;
    }

    pub fn get(&self, i: usize) -> T {
        unsafe { self.items[i].assume_init() }
    }

    pub fn get_ptr(&self, i: usize) -> &MaybeUninit<T> {
        &self.items[i]
    }
}

impl<T, const N: usize> Display for ArrayVec<T, N>
where
    T: Copy + Display,
{
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut fmt_str = "[".to_string();
        for i in 0..self.length {
            fmt_str.push_str(format!("{}", self.get(i)).as_str());
            if i != self.items.len() - 1 {
                fmt_str.push_str(", ");
            }
        }
        fmt_str.push_str("Uninitialized: ");
        for i in self.length..self.items.len() {
            fmt_str.push_str(format!("{:?}", self.get_ptr(i).as_ptr()).as_str());
            if i != self.items.len() - 1 {
                fmt_str.push_str(", ");
            }
        }
        fmt_str.push_str("]");
        write(f, format_args!("{}", fmt_str))
    }
}

fn main() {
    let mut arrVec = ArrayVec::<i32, 5>::new();
    arrVec.push(1);
    println!("{}", arrVec);
    println!("{:?}", arrVec.get(0));
    println!("{:?}", arrVec.get(4));
    for i in 0..5 {
        print!("{}:{:?},", i, arrVec.get_ptr(i).as_ptr());
    }
}

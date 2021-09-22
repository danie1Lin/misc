#![feature(const_fmt_arguments_new)]
#![feature(const_fn_fn_ptr_basics)]
// #![feature(min_const_generics)] stable since 1.51
const fn test() -> i32 {
    1
}

const fn test1(x: i32) -> i32 {
    //for let i in 1..2 { print!(1)}
    // println!("{:?}", 1); //alls in constant functions are limited to
    //constant functions, tuple structs and tuple variants
    match x {
        1 => x + 1,
        _ => x,
    }
}

static mut a1: Vec<i32> = vec![];
fn main() {
    "test".to_string();
    test();
    test1(1);
    println!(
        "{}",
        const_sha1::sha1(&const_sha1::ConstBuffer::from_slice(
            stringify!(MyType).as_bytes()
        ))
    );
    unsafe {
        for i in 1..100 {
            a1.push(i);
        }
        println!("{:?}", a1);
    }
}

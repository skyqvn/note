use libloader::{get_libfn, libloading};

fn main() {
    get_libfn!("add.dll", "add", add, i32, a:i32,b:i32);
    // get_libfn!($链接库路径,$函数名,$变量,$返回值类型,$参数1类型,···);
    println!("{}", add(1, 2));
}

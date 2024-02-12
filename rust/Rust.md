# Rust



## &和ref

```rust
fn main() {
    let ref a = 100_i32;
    //等于
    let a = &100_i32;
}

```



## 字符串

```rust
let mut s: &str = "Hello World";
let mut s2: String = String::from(s);
s.as_bytes();
s.bytes();
s.chars();
s.chars().count();
s.chars();
s.contains("Hello");
s2.clear();
s.parse::<i32>().unwrap();
s.trim();

```



## 向量

```rust
let mut vect: Vec<i32> = vec![100, 200, 300, 400, 500];
vect.get(0);
vect[0];
vect[0] = 120;
vect.push(600);
vect.insert(2, 250);
vect.remove(0);
vect.append(&mut vec![700, 800]);
//vect一般是引用（&vect或&mut vect）
for i in &vect {
    println!("{i}");
}
vect.pop();
vect.clear();

```



## 哈希表

```rust
let mut hm: HashMap<String, i32> = HashMap::new();
hm.insert(String::from("k"), 100);
hm.get("k").copied().unwrap();
hm.entry(String::from("k")).or_insert(100);
hm.remove("k");
hm.remove_entry("k").unwrap();

```



## 迭代器

```rust
let iter: Copied<Iter<'_, i32>> = vect.iter().copied();
iter.clone().filter(|x| *x == 300);
iter.clone().map(|x| x * x);
iter.clone().collect::<Vec<i32>>();
iter.clone().chain(vec![600, 700, 800].iter().copied());
iter.clone().enumerate();
iter.clone().count();

```



## Result/Option处理

```rust
let err: Result<String, Error> = read_to_string("a.txt");
err.unwrap();
err.expect("···");
err.unwrap_or(String::from("Error!"));
err.unwrap_or_default();
err.unwrap_or_else(|e| format!("Error:{}", e.to_string()));
err?;
//Result特有
err.is_ok();
err.is_err();
//Option特有
opt.is_some();
opt.is_none();

```



## Derive

```rust
/// Debug：通过实现 Debug trait，可以使用 println!("{:?}", my_struct) 来打印结构体的调试信息。
/// Copy：可以按位复制（不含引用等），可以在赋值时自动复制，所有实现了 Copy 的类型都必须实现 Clone。
/// Clone：通过实现 Clone trait，可以使用 my_struct.clone() 创建结构体的克隆副本。
/// PartialEq 和 Eq：通过实现 PartialEq trait，可以进行结构体的部分相等性比较，而 Eq trait 则实现了完全相等性比较。
/// PartialOrd 和 Ord：通过实现 PartialOrd trait，可以对结构体进行部分有序性比较，而 Ord trait 实现了完全有序性比较。
#[derive(Debug,Clone)]
struct MyStruct {
    // 结构体字段
}

```



## 类型转换

```rust
let decimal = 65.4321_f32;
let integer = decimal as u8;
let character = integer as char;
```



## 类型别名

```rust
type NanoSecond = u64;
```



## 运算符重载

```rust
use std::ops;

struct Foo;
struct Bar;

#[derive(Debug)]
struct FooBar;

#[derive(Debug)]
struct BarFoo;

// `std::ops::Add` trait 用来指明 `+` 的功能，这里我们实现 `Add<Bar>`，它是用于
// 把对象和 `Bar` 类型的右操作数（RHS）加起来的 `trait`。
// 下面的代码块实现了 `Foo + Bar = FooBar` 这样的运算。
impl ops::Add<Bar> for Foo {
    type Output = FooBar;

    fn add(self, _rhs: Bar) -> FooBar {
        println!("> Foo.add(Bar) was called");

        FooBar
    }
}

// 通过颠倒类型，我们实现了不服从交换律的加法。
// 这里我们实现 `Add<Foo>`，它是用于把对象和 `Foo` 类型的右操作数加起来的 trait。
// 下面的代码块实现了 `Bar + Foo = BarFoo` 这样的运算。
impl ops::Add<Foo> for Bar {
    type Output = BarFoo;

    fn add(self, _rhs: Foo) -> BarFoo {
        println!("> Bar.add(Foo) was called");

        BarFoo
    }
}

fn main() {
    println!("Foo + Bar = {:?}", Foo + Bar);
    println!("Bar + Foo = {:?}", Bar + Foo);
}

```



## 常量

```rust
///const：不可改变的值（通常使用这种）。
///static：具有 'static 生命周期的，可以是可变的变量（static mut）。
///有个特例就是 "string" 字面量。它可以不经改动就被赋给一个 static 变量，因为它的类型标记：&'static str 就包含了所要求的生命周期 'static。其他的引用类型都必须特地声明，使之拥有 'static 生命周期。这两种引用类型的差异似乎也无关紧要，因为无论如何，static 变量都得显式地声明。
static A: &'static str = "Rust";
static mut B: &'static str = "Rust";
const C: &str = "Rust";
```



## From与Into

> [`From`](https://rustwiki.org/zh-CN/std/convert/trait.From.html) trait 允许一种类型定义 “怎么根据另一种类型生成自己”，因此它提供了一种类型转换的简单机制。
> [`Into`](https://rustwiki.org/zh-CN/std/convert/trait.Into.html) trait 就是把 `From` trait 倒过来而已。也就是说，如果你为你的类型实现了 `From`，那么同时你也就免费获得了 `Into`。

```rust
use std::convert::From;

#[derive(Debug)]
struct Number {
    value: i32,
}

impl From<i32> for Number {
    fn from(item: i32) -> Self {
        Number { value: item }
    }
}

fn main() {
    let num = Number::from(30);
    let num: Number = 30.into();
    println!("My number is {:?}", num);
}

```

> [`TryFrom`](https://rustwiki.org/zh-CN/rust-by-example/conversion/try_from_try_into.html) 和 [`TryInto`](https://rustwiki.org/zh-CN/rust-by-example/conversion/try_from_try_into.html) trait 用于易出错的转换，也正因如此，其返回值是 [`Result`](https://rustwiki.org/zh-CN/std/result/enum.Result.html) 型。

> 要把任何类型转换成 `String`，您应该实现[`fmt::Display`](https://rustwiki.org/zh-CN/std/fmt/trait.Display.html) trait，它会自动提供 [`ToString`](https://rustwiki.org/zh-CN/std/string/trait.ToString.html)，并且还可以用来打印类型，就像 [`print!`](https://rustwiki.org/zh-CN/rust-by-example/hello/print.html) 一节中讨论的那样。See:[ToString、fmt::Display](https://rustwiki.org/zh-CN/rust-by-example/conversion/string.html)，[FromStr](https://rustwiki.org/zh-CN/std/str/trait.FromStr.html)
>



## 文件分层

<https://rustwiki.org/zh-CN/rust-by-example/mod/split.html>



## Drop

<https://rustwiki.org/zh-CN/rust-by-example/trait/drop.html>

```rust
struct Droppable {
    name: &'static str,
}

// 这个简单的 `drop` 实现添加了打印到控制台的功能。
impl Drop for Droppable {
    fn drop(&mut self) {
        println!("> Dropping {}", self.name);
    }
}

fn main() {
    let _a = Droppable { name: "a" };

    // 代码块 A
    {
        let _b = Droppable { name: "b" };

        // 代码块 B
        {
            let _c = Droppable { name: "c" };
            let _d = Droppable { name: "d" };

            println!("Exiting block B");
        }
        println!("Just exited block B");

        println!("Exiting block A");
    }
    println!("Just exited block A");

    // 变量可以手动使用 `drop` 函数来销毁。
    drop(_a);
    //或std::mem::drop(_a);

    println!("end of the main function");

    // `_a` 不会在这里再次销毁，因为它已经被（手动）销毁。
}

```



## Crate
<https://rustwiki.org/zh-CN/rust-by-example/crates.html>

```bash
# 默认情况下，库会使用 crate 文件的名字，前面加上 “lib” 前缀，但这个默认名称可以使用 crate_name 属性 覆盖。
rustc --crate-type=lib rary.rs
rustc main.rs --extern rary=library.rlib --edition=2018
```

```toml
[package]
# See: https://doc.rust-lang.org/cargo/reference/manifest.html
name = "hello_world"
version = "0.1.0"
edition = "2018"
sql = "0.4.3"# <-- Add dependency here
```



## cfg

`#[cfg()]` 是 Rust 中的一个属性，用于根据配置条件来选择性地包含或排除代码。

`#[cfg()]` 属性中使用的逻辑运算符有以下几种：

1. `all(expr1, expr2, ...)`：逻辑与运算符，要求所有条件表达式都为真才返回真。例如：`#[cfg(all(feature = "foo", target_os = "linux"))]` 表示只有在启用 "foo" 功能并且目标操作系统是 Linux 时条件成立。
2. `any(expr1, expr2, ...)`：逻辑或运算符，只要有任一条件表达式为真就返回真。例如：`#[cfg(any(feature = "foo", feature = "bar"))]` 表示只要启用 "foo" 或 "bar" 中的任意一个功能时条件成立。
3. `not(expr)`：逻辑非运算符，对条件表达式取反。例如：`#[cfg(not(debug_assertions))]` 表示只有在非调试断言模式下条件成立。

```rust
#[cfg(target_os = "linux")]
fn only_on_linux() {
    // 仅在 Linux 系统上编译和执行的代码
}

#[cfg(all(unix, not(target_os = "linux")))]
fn on_unix_but_not_linux() {
    // 仅在 Unix 系统但不是 Linux 上编译和执行的代码
}

#[cfg(any(windows, target_os = "macos"))]
fn on_windows_or_macos() {
    // 仅在 Windows 或 macOS 上编译和执行的代码
}

#[cfg(not(debug_assertions))]
fn when_not_debug_assertions() {
    // 仅在非调试断言模式下编译和执行的代码
}

#[cfg(feature = "my_feature")]
fn with_my_feature_enabled() {
    // 仅在启用 "my_feature" 功能时编译和执行的代码
}

```

```rust
if cfg!(target_os = "linux") {
    println!("Yes. It's definitely linux!");
} else {
    println!("Yes. It's definitely *not* linux!");
}

```

## 宏

```text
macro_rules! $name {
    $rule0 ;
    $rule1 ;
    // …
    $ruleN ;
}
```

### \$rule：
`($pattern:$identifier) => {$expansion}`

### \$identifier:
- `item: 条目，比如函数、结构体、模组等。`
- `block: 区块(即由花括号包起的一些语句加上/或是一项表达式)。`
- `stmt: 语句`
- `pat: 模式`
- `expr: 表达式`
- `ty: 类型`
- `ident: 标识符`
- `path: 路径 (例如 foo, ::std::mem::replace, transmute::<_, int>, …)`
- `meta: 元条目，即被包含在 #[...]及#![...]属性内的东西。`
- `tt: 标记树`

### 重复
- `$` 是字面标记。
- `( ... )` 代表了将要被重复匹配的模式，由小括号包围。
- `sep`是一个可选的分隔标记。常用例子包括`,`和`;`。
- `rep`是重复控制标记。当前有两种选择，分别是`*` (代表接受0或多次重复)以及`+` (代表1或多次重复)。目前没有办法指定“0或1”或者任何其它更加具体的重复计数或区间。

```rust
macro_rules! vec_strs {
    (
        // 重复开始：
        $(
            // 每次重复必须有一个表达式...
            $element:expr
        )
        // ...重复之间由“,”分隔...
        ,
        // ...总共重复0或多次.
        *
    ) => {
        // 为了能包含多条语句，
        // 我们将扩展部分包裹在花括号中...
        {
            let mut v = Vec::new();
            // 重复开始：
            $(
                // 每次重复将包含如下元素，其中
                // “$element”将被替换成其相应的展开...
                v.push(format!("{}", $element));
            )*
            v
        }
    };
}

```



## Mutex和Arc

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();

            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
}

```



## 智能指针

- `Box<T>`：它提供了最简单的堆资源分配方式，单所有权。Box类型拥有其中的值，并且可用于保存结构体中的值，或者从函数返回它们。

- `Rc<T>`：它用于引用计数。每当获取新引用时，计数器会执行递增操作，并在用户释放引用时对计数器执行递减操作。当计数器的值为零时，该值将被移除。

- `Arc<T>`：它用于原子引用计数。这与之前的类型类似，但具有原子性以保证多线程的安全性。

- `Cell<T>`：它为我们提供实现了Copy特征的类型的内部可变性。换句话说，我们有可能获得多个可变引用。

- `RefCell<T>`：它为我们提供了类型的内部可变性，并且不需要实现Copy特征。它用于运行时的锁定以确保安全性。



### RefCell内部可变性

```rust
let a = RefCell::new(100_i32);
//编译失败（RefCell不能解引用）：
// println!("{}", *a);
println!("{}", *a.borrow());
//编译失败（a不可变）：
// *a = 120;
*a.borrow_mut() = 120;
//Panic（同时两个可变引用）：
let b = a.borrow_mut();
let c = a.borrow_mut();
```



### RefCell+Rc: 多个所有者 并且 可以修改值

```rust
#[derive(Debug)]
enum List {
    Cons(Rc<RefCell<i32>>, Rc<List>),
    Nil,
}

use crate::List::{Cons, Nil};
use std::cell::RefCell;
use std::rc::Rc;

fn main() {
    let value = Rc::new(RefCell::new(5));
    let a = Rc::new(Cons(Rc::clone(&value), Rc::new(Nil)));
    let b = Cons(Rc::clone(&value), Rc::clone(&a));
    *value.borrow_mut() += 10;
    println!("{:?}\n{:?}", a, b);
}

```



### 强/弱引用

```rust
let strong = Rc::new(100_i32);
//由强引用创建弱引用。
let weak = Rc::downgrade(&strong);
//由弱引用创建强引用。
//如果弱引用已无效，就返回None。
let strong = Weak::upgrade(&weak).unwrap();

```


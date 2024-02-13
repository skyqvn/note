cd add
cargo build --release
cd ..\
cp .\add\target\release\add.dll .\test_add
cd test_add
cargo build --release
.\target\release\test_add.exe


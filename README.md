cfdGO
cfdGO is a Golang libarary for solving fluid mechanics problems. It provide efficient impliment for working with 
sparse matrix with forma csr, dok and coo. This library also provide efficient linear solvin by using amgcl library from c++/c.
INSTALL
    WINDOWS
    1)install MSYS MINGW (and set up PATH in env. variables) => in MSYS command line...
    2)pacman -S --needed base-devel mingw-w64-ucrt-x86_64-toolchain
    3)install Boost library sudo apt install libboost-all-dev
    4)install golang pacman -S mingw-w64-x86_64-go
    echo 'export GOROOT="/mingw64/lib/go"' >> ~/.bashrc
    echo 'export PATH="/mingw64/bin:$PATH"' >> ~/.bashrc
    source ~/.bashrc
    5)to user references you can set up amgcl in /solvers/amgcl/ and build .a labrary file by using the command 
    g++ -c amgcl_wrapper.cpp -Iinclude -O2 -g -o amgcl_wrapper.o
    ar rcs libamgcl_wrapper.a amgcl_wrapper.o
    go build
    6)to run project in source directory 
    go build
    ./test.com.exe
    LINUX(UBUNTU, DEBIAN, ...)
    1)install boost library 
    sudo apt update
    sudo apt install libboost-all-dev
    2)install golang
    3)to user references you can set up amgcl in /solvers/amgcl/ and build .a labrary file by using the command
    g++ -shared -fPIC -o libamgcl.so amgcl_wrapper.cpp -I./include -std=c++11
    3)to run project in source directory
    go run .
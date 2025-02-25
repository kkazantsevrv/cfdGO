# cfdGO

**cfdGO** is a Golang library for solving fluid mechanics problems. It provides efficient implementations for working with sparse matrices in formats such as CSR, DOK, and COO. Additionally, the library offers efficient linear solving capabilities by leveraging the **AMGCL** library from C++/C.

---

## Features

- **Sparse Matrix Formats**:
  - CSR (Compressed Sparse Row)
  - DOK (Dictionary of Keys)
  - COO (Coordinate Format)

- **Linear Solvers**:
  - Integration with the **AMGCL** library for efficient linear algebra operations.

---

## Installation

### Windows

1. **Install MSYS2 and MINGW**:
   - Download and install [MSYS2](https://www.msys2.org/).
   - Add MSYS2 and MINGW to your system's `PATH` environment variable.

2. **Install Required Tools**:
   - Open the MSYS2 terminal and run:
     ```bash
     pacman -S --needed base-devel mingw-w64-ucrt-x86_64-toolchain
     ```

3. **Install Boost Library**:
   - In the MSYS2 terminal, run:
     ```bash
     pacman -S mingw-w64-x86_64-boost
     ```

4. **Install Golang**:
   - In the MSYS2 terminal, run:
     ```bash
     pacman -S mingw-w64-x86_64-go
     ```
   - Add Go to your environment variables:
     ```bash
     echo 'export GOROOT="/mingw64/lib/go"' >> ~/.bashrc
     echo 'export PATH="/mingw64/bin:$PATH"' >> ~/.bashrc
     source ~/.bashrc
     ```

5. **Build AMGCL Library**:
   - Place the AMGCL library in `/solvers/amgcl/`.
   - Build the `.a` library file:
     ```bash
     g++ -c amgcl_wrapper.cpp -Iinclude -O2 -g -o amgcl_wrapper.o
     ar rcs libamgcl_wrapper.a amgcl_wrapper.o
     ```

6. **Build and Run the Project**:
   - In the project's root directory, run:
     ```bash
     go build
     ./test.com.exe
     ```

---

### Linux (Ubuntu, Debian, etc.)

1. **Install Boost Library**:
   - Run the following commands:
     ```bash
     sudo apt update
     sudo apt install libboost-all-dev
     ```

2. **Install Golang**:
   - Follow the official [Go installation guide](https://golang.org/doc/install).

3. **Build AMGCL Library**:
   - Place the AMGCL library in `/solvers/amgcl/`.
   - Build the `.so` library file:
     ```bash
     g++ -shared -fPIC -o libamgcl.so amgcl_wrapper.cpp -I./include -std=c++11
     ```

4. **Build and Run the Project**:
   - In the project's root directory, run:
     ```bash
     go run .
     ```

---

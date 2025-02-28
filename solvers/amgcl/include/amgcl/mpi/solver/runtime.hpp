#ifndef AMGCL_MPI_SOLVER_RUNTIME_HPP
#define AMGCL_MPI_SOLVER_RUNTIME_HPP

/*
The MIT License

Copyright (c) 2012-2022 Denis Demidov <dennis.demidov@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

/**
 * \file   amgcl/mpi/solver/runtime.hpp
 * \author Denis Demidov <dennis.demidov@gmail.com>
 * \brief  Runtime-configurable MPI wrapper around amgcl iterative solvers.
 */

#include <amgcl/solver/runtime.hpp>
#include <amgcl/mpi/inner_product.hpp>

namespace amgcl {
namespace runtime { 
namespace mpi {
namespace solver {

template <class Backend, class InnerProduct = amgcl::mpi::inner_product>
struct wrapper : public amgcl::runtime::solver::wrapper<Backend, InnerProduct> {
    typedef amgcl::runtime::solver::wrapper<Backend, InnerProduct> Base;
    using Base::Base;
};

} // namespace solver
} // namespace mpi
} // namespace runtime
} // namespace amgcl

#endif

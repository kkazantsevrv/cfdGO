#include <vector>
#include <amgcl/amg.hpp>
#include <amgcl/make_solver.hpp>
#include <amgcl/solver/cg.hpp>
#include <amgcl/adapter/crs_tuple.hpp>
#include <amgcl/coarsening/smoothed_aggregation.hpp>
#include <amgcl/relaxation/damped_jacobi.hpp>
#include "amgcl_wrapper.h"

#ifdef _WIN32
#define EXPORT __declspec(dllexport)
#else
#define EXPORT
#endif

typedef amgcl::make_solver<
    amgcl::amg<amgcl::backend::builtin<double>, 
                amgcl::coarsening::smoothed_aggregation, 
                amgcl::relaxation::damped_jacobi>,
    amgcl::solver::cg<amgcl::backend::builtin<double>>
> Solver;

struct AMGCLSolverImpl {
    std::shared_ptr<Solver> solver;
};

AMGCLSolver create_solver(int n, int* rows, int* cols, double* values) {
    std::vector<int> ptr(rows, rows + n + 1);
    std::vector<int> col(cols, cols + ptr[n]);
    std::vector<double> val(values, values + ptr[n]);
    AMGCLSolverImpl* impl = new AMGCLSolverImpl();
    impl->solver = std::make_shared<Solver>(std::make_tuple(n, ptr, col, val));
    return (AMGCLSolver)impl;
}

void solve_system(AMGCLSolver solver, double* rhs, double* x) {
    AMGCLSolverImpl* impl = (AMGCLSolverImpl*)solver;
    std::vector<double> RHS(rhs, rhs + impl->solver->size());
    std::vector<double> X(x, x + impl->solver->size());
    size_t iters;
    double error;
    std::tie(iters, error) = (*impl->solver)(RHS, X);
    std::copy(X.begin(), X.end(), x);
}

void destroy_solver(AMGCLSolver solver) {
    AMGCLSolverImpl* impl = (AMGCLSolverImpl*)solver;
    delete impl;
}
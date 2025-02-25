#ifndef AMGCL_WRAPPER_H
#define AMGCL_WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct amgcl_solver_t amgcl_solver_t;
typedef void* AMGCLSolver;
AMGCLSolver create_solver(int n, int* rows, int* cols, double* values);
void solve_system(AMGCLSolver solver, double* rhs, double* x);
void destroy_solver(AMGCLSolver solver);

#ifdef __cplusplus
}
#endif

#endif
# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.22

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = "/home/user/Рабочий стол/projects/golang/ff/vendor/tests"

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = "/home/user/Рабочий стол/projects/golang/ff/build"

# Include any dependencies generated for this target.
include CMakeFiles/test_solver_block_crs.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include CMakeFiles/test_solver_block_crs.dir/compiler_depend.make

# Include the progress variables for this target.
include CMakeFiles/test_solver_block_crs.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/test_solver_block_crs.dir/flags.make

CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o: CMakeFiles/test_solver_block_crs.dir/flags.make
CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o: /home/user/Рабочий\ стол/projects/golang/ff/vendor/tests/test_solver_block_crs.cpp
CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o: CMakeFiles/test_solver_block_crs.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir="/home/user/Рабочий стол/projects/golang/ff/build/CMakeFiles" --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o -MF CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o.d -o CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o -c "/home/user/Рабочий стол/projects/golang/ff/vendor/tests/test_solver_block_crs.cpp"

CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.i"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E "/home/user/Рабочий стол/projects/golang/ff/vendor/tests/test_solver_block_crs.cpp" > CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.i

CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.s"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S "/home/user/Рабочий стол/projects/golang/ff/vendor/tests/test_solver_block_crs.cpp" -o CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.s

# Object files for target test_solver_block_crs
test_solver_block_crs_OBJECTS = \
"CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o"

# External object files for target test_solver_block_crs
test_solver_block_crs_EXTERNAL_OBJECTS =

test_solver_block_crs: CMakeFiles/test_solver_block_crs.dir/test_solver_block_crs.o
test_solver_block_crs: CMakeFiles/test_solver_block_crs.dir/build.make
test_solver_block_crs: CMakeFiles/test_solver_block_crs.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir="/home/user/Рабочий стол/projects/golang/ff/build/CMakeFiles" --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable test_solver_block_crs"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/test_solver_block_crs.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/test_solver_block_crs.dir/build: test_solver_block_crs
.PHONY : CMakeFiles/test_solver_block_crs.dir/build

CMakeFiles/test_solver_block_crs.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/test_solver_block_crs.dir/cmake_clean.cmake
.PHONY : CMakeFiles/test_solver_block_crs.dir/clean

CMakeFiles/test_solver_block_crs.dir/depend:
	cd "/home/user/Рабочий стол/projects/golang/ff/build" && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" "/home/user/Рабочий стол/projects/golang/ff/vendor/tests" "/home/user/Рабочий стол/projects/golang/ff/vendor/tests" "/home/user/Рабочий стол/projects/golang/ff/build" "/home/user/Рабочий стол/projects/golang/ff/build" "/home/user/Рабочий стол/projects/golang/ff/build/CMakeFiles/test_solver_block_crs.dir/DependInfo.cmake" --color=$(COLOR)
.PHONY : CMakeFiles/test_solver_block_crs.dir/depend


CXX = g++
CXXFLAGS = -std=c++11 -fPIC -I../../vendor/amgcl -O3
LDFLAGS = -shared

SOURCES = amgcl_wrapper.cpp
OBJECTS = $(SOURCES:.cpp=.o)
TARGET = libamgcl.so

.PHONY: all clean

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CXX) $(LDFLAGS) -o $@ $^

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<

clean:
	rm -f $(OBJECTS) $(TARGET)

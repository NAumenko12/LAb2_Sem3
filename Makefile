CXX = clang++
CXXFLAGS = -std=c++11 -O2 -Wall -Wextra

STRUCT_CPP = $(wildcard struct_cplas/*/*.cpp)
TASK_CPP = $(wildcard zadachi/*/*.cpp)
HEADERS = $(wildcard struct_cplas/*/*.h) $(wildcard zadachi/*/*.h)
TARGET = build/lab2
GO_TASKS = zad_stek zad_atd_mnozh zad_mnozh zad_array zad_tree zad_hesh_tab zad_llru_lfu

all: $(TARGET)

$(TARGET): main.cpp $(STRUCT_CPP) $(TASK_CPP) $(HEADERS)
	@mkdir -p build
	$(CXX) $(CXXFLAGS) main.cpp $(TASK_CPP) $(STRUCT_CPP) -o $(TARGET)

go:
	@mkdir -p build/go
	cd go && go build -o ../build/go/lab2 .
	cd go && for task in $(GO_TASKS); do go build -o ../build/go/$$task ./$$task || exit 1; done

clean:
	rm -rf build

.PHONY: all go clean

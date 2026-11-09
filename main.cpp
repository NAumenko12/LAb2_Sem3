#include <iostream>
#include <sstream>
#include <string>
#include "zadachi/zadanie1/zad_stek.h"
#include "zadachi/zadanie2/zad_atd_mnozh.h"
#include "zadachi/zadanie3/zad_mnozh.h"
#include "zadachi/zadanie4/zad_array.h"
#include "zadachi/zadanie5/zad_tree.h"
#include "zadachi/zadanie6/zad_hesh_tab.h"
#include "zadachi/zadanie7/zad_llru_lfu.h"
using namespace std;

void PrintHelp(){
    cout << "Запуск: lab2 --query 'КОМАНДА данные'\n"
         << "DFS TREE 7 3 9 1 5\n"
         << "DFS GRAPH 3 2 1 2 2 3 1\n"
         << "INTER {{1,2,3},{2,3,4}}\n"
         << "PYRAMID 3 3 1 2 2 3 3\n"
         << "HUFFMAN abracadabra\n"
         << "SUBSTRING abcabcbb\n"
         << "CUCKOO a l w\n"
         << "LFU 2 4 SET 1 10 SET 2 20 GET 1 GET 3\n"
         << "lab2 --file set.json --query 'SETADD 5'\n"
         << "lab2 --query 'HUFFMAN aaaa' --bits 0000\n";
}

int RunQuery(const string& query, const string& filename, const string& bits){
    istringstream input(query);
    string command;
    if (!(input >> command)){
        cerr << "Ошибка: пустой запрос\n";
        return 1;
    }

    if (command == "SETADD" || command == "SETDEL" || command == "SET_AT" || command == "SETPRINT"){
        if (filename.empty()){
            cerr << "Ошибка: укажи --file для множества\n";
            return 1;
        }
        return RunSetQuery(filename, query);
    }
    if (command == "DFS"){
        string type;
        input >> type;
        if (type == "TREE") return RunTreeDfs(input);
        if (type == "GRAPH") return RunGraphDfs(input);
        cerr << "Ошибка: укажи DFS TREE или DFS GRAPH\n";
        return 1;
    }
    if (command == "PYRAMID") return RunPyramid(input);
    if (command == "LFU") return RunLFU(input);

    string data;
    if (input.get() != char_traits<char>::eof()){
        getline(input, data, '\0');
    }
    if (command == "INTER") return RunIntersections(data);
    if (command == "HUFFMAN") return RunHuffman(data, bits);
    if (command == "SUBSTRING"){
        RunLongestSubstring(data);
        return 0;
    }
    if (command == "CUCKOO"){
        DemoHashTables(data.empty() ? "apple banana cherry" : data);
        return 0;
    }

    cerr << "Ошибка: неизвестная команда " << command << '\n';
    return 1;
}

int main(int argc, char* argv[]){
    if (argc == 1 || (argc == 2 && string(argv[1]) == "--help")){
        PrintHelp();
        return 0;
    }

    string query, filename, bits;
    for (int i = 1; i < argc; ++i){
        string arg = argv[i];
        if (i + 1 >= argc || (arg != "--query" && arg != "--file" && arg != "--bits")){
            cerr << "Ошибка: неизвестный или неполный параметр " << arg << '\n';
            return 1;
        }
        if (arg == "--query") query = argv[++i];
        else if (arg == "--file") filename = argv[++i];
        else bits = argv[++i];
    }
    return RunQuery(query, filename, bits);
}

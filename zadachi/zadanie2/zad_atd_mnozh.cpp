#include "zad_atd_mnozh.h"
#include "../../struct_cplas/Mnozh/mnozh.h"
#include <fstream>
#include <sstream>
using namespace std;

void LoadSet(Set& set, const string& filename){
    ifstream in(filename);
    if (!in.is_open()){
        return;
    }
    string text, line;
    while (getline(in, line)){
        text += line;
    }
    size_t i = 0;
    while (i < text.size()){
        if (text[i] != '"'){
            i++;
            continue;
        }
        string word;
        size_t j = i + 1;
        while (j < text.size()){
            if (text[j] == '\\' && j + 1 < text.size()){
                word += text[j + 1];
                j += 2;
            } else if (text[j] == '"'){
                break;
            } else {
                word += text[j];
                j++;
            }
        }
        if (j == text.size()){
            break;
        }

        SetAdd(set, word);
        i = j + 1;
    }
}

void SaveSet(const Set& set, ofstream& out){
    Array elements;
    MInitArray(elements, SetSize(set));
    SetElements(set, elements);
    out << '[';
    for (int i = 0; i < MLength(elements); ++i){
        if (i > 0) out << ", ";
        out << '"';
        for (char c : MGet(elements, i)){
            if (c == '"' || c == '\\'){
                out << '\\';
            }
            out << c;
        }
        out << '"';
    }
    out << "]\n";
    MDestroy(elements);
}

int RunSetQuery(const string& filename, const string& query){
    istringstream input(query);
    string command, value, extra;
    input >> command;
    if (command != "SETADD" && command != "SETDEL" && command != "SET_AT" && command != "SETPRINT"){
        cerr << "Ошибка: неизвестная команда " << command << '\n';
        return 1;
    }
    if (command != "SETPRINT" && !(input >> value)){
        cerr << "Ошибка: нужно значение элемента\n";
        return 1;
    }
    if (input >> extra){
        cerr << "Ошибка: слишком много аргументов команды\n";
        return 1;
    }
    for (unsigned char c : value){
        if (c < ' '){
            cerr << "Ошибка: управляющий символ в элементе\n";
            return 1;
        }
    }

    Set set;
    SetInit(set);
    LoadSet(set, filename);

    int size_before = SetSize(set);
    if (command == "SETADD"){
        SetAdd(set, value);
        if (SetSize(set) == size_before){
            cout << value << " уже есть в множестве\n";
        }
    } else if (command == "SETDEL"){
        SetDel(set, value);
        if (SetSize(set) == size_before){
            cout << value << " нет в множестве\n";
        }
    } else if (command == "SET_AT"){
        cout << "-> " << (SetAt(set, value) ? "TRUE" : "FALSE") << '\n';
    } else {
        Array elements;
        MInitArray(elements, SetSize(set));
        SetElements(set, elements);
        cout << '{';
        for (int i = 0; i < MLength(elements); ++i){
            if (i > 0) cout << ", ";
            cout << MGet(elements, i);
        }
        cout << "}\n";
        MDestroy(elements);
    }
    if (SetSize(set) != size_before){
        ofstream out(filename);
        if (out.is_open()){
            SaveSet(set, out);
            out.close();
        }
        if (!out){
            cerr << "Ошибка: не удалось записать файл " << filename << '\n';
            SetDestroy(set);
            return 1;
        }
        cout << "-> " << value;
        if (command == "SETDEL") cout << " удалён";
        cout << '\n';
    }

    SetDestroy(set);
    return 0;
}

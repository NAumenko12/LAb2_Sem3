#include "zad_tree.h"
#include "../../struct_cplas/tree/huffman_tree.h"
#include <iostream>
using namespace std;

void PrintHuffmanTable(const Array& codes, const int frequency[]){
    cout << "Символ — частота — код:\n";
    for (int i = 0; i < MLength(codes); ++i){
        string entry = MGet(codes, i);
        unsigned char symbol = entry[0];
        if (symbol == ' '){
            cout << "пробел";
        } else if (symbol > 32 && symbol < 127){
            cout << symbol;
        } else {
            cout << "байт " << (int)symbol;
        }
        cout << " — " << frequency[symbol] << " — " << entry.substr(2) << '\n';
    }
}

int RunHuffman(const string& text, const string& decode_bits){
    if (text.empty()){
        cout << "Строка пустая\n";
        return 1;
    }
    int frequency[256]{};
    HufCountFrequencies(text, frequency);
    HufNode* root = HufBuildTree(frequency);

    Array codes;
    MInitArray(codes, 256);
    HufBuildCodes(root, "", codes);
    PrintHuffmanTable(codes, frequency);

    string bits = HufEncode(text, codes);
    string decoded;
    HufDecode(root, bits, decoded);
    cout << "Код: " << bits << '\n';
    cout << "Расшифровка: " << decoded << '\n';
    bool correct = true;
    if (!decode_bits.empty()){
        correct = HufDecode(root, decode_bits, decoded);
        if (correct){
            cout << "Результат: " << decoded << '\n';
        } else {
            cout << "Ошибка: код содержит не 0/1 или обрывается внутри символа\n";
        }
    }
    MDestroy(codes);
    HufDeleteTree(root);
    return correct ? 0 : 1;
}

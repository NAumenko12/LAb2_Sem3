#include "zad_stek.h"
#include <iostream>
#include <sstream>
#include <string>
#include "../../struct_cplas/Stack/stack.h"
#include "../../struct_cplas/Array/massiv.h"
#include "../../struct_cplas/tree/binary_tree.h"
using namespace std;

void TreeDFS(BNode* root){
    if (root == nullptr){
        cout << "дерево пустое\n";
        return;
    }
    Stack stack;
    SInitStack(&stack);
    SPush(stack, to_string(root->data));
    while (!SIsEmptyStack(&stack)){
        int chislo = stoi(STop(stack));
        SPop(stack);
        BNode* node = BFind(root, chislo);
        cout << node->data << ' ';
        if (node->right != nullptr){
            SPush(stack, to_string(node->right->data));
        }
        if (node->left != nullptr){
            SPush(stack, to_string(node->left->data));
        }
    }
}


int RunTreeDfs(istream& input){
    BTree tree;
    BInitTree(tree);
    string token;
    while (input >> token){
        istringstream number(token);
        int chislo;
        if (!(number >> chislo) || number.peek() != char_traits<char>::eof()){
            cout << "Ошибка: нужны целые числа для дерева\n";
            BClear(tree.root);
            return 1;
        }
        if (chislo == 0) break;
        BAdd(tree.root, chislo);
    }
    cout << "Обход в глубину  ";
    TreeDFS(tree.root);
    cout << '\n';
    BClear(tree.root);
    return 0;
}

void GraphDFS(Array* sosed, int n, int start){
    bool* visit = new bool[n + 1]{};
    Stack stack;
    SInitStack(&stack);
    SPush(stack, to_string(start));
    while (!SIsEmptyStack(&stack)){
        int v = stoi(STop(stack));
        SPop(stack);
        if (visit[v]){
            continue;
        }
        visit[v] = true;
        cout << v << ' ';
        for (int i = 0; i < MLength(sosed[v]); ++i){
            SPush(stack, MGet(sosed[v], i));
        }
    }
    cout << '\n';
    delete[] visit;
}
int RunGraphDfs(istream& input){
    int n, m;
    if (!(input >> n >> m) || n < 1 || m < 0){
        cout << "Ошибка: неверное количество вершин или рёбер\n";
        return 1;
    }
    Array* sosed = new Array[n + 1];
    for (int i = 0; i <= n; ++i){
        MInitArray(sosed[i], 2);
    }
    bool correct = true;
    for (int i = 0; i < m; ++i){
        int u, v;
        if (!(input >> u >> v) || u < 1 || u > n || v < 1 || v > n){
            cout << "Ошибка: неверное ребро\n";
            correct = false;
            break;
        }
        MAddEnd(sosed[u], to_string(v));
        MAddEnd(sosed[v], to_string(u));
    }
    if (correct){
        int start;
        string extra;
        if (input >> start && start >= 1 && start <= n && !(input >> extra)){
            cout << "Обход в глубину: ";
            GraphDFS(sosed, n, start);
        } else {
            cout << "Ошибка: неверная начальная вершина или лишние данные\n";
            correct = false;
        }
    }
    for (int i = 0; i <= n; ++i){
        MDestroy(sosed[i]);
    }
    delete[] sosed;
    return correct ? 0 : 1;
}


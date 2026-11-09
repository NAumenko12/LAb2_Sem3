#include <iostream>
#include <string>
#include "stack.h"
#include <stdexcept>
using namespace std;

void SInitStack(Stack* stack){
    stack->head = nullptr;
}
bool SIsEmptyStack(Stack* stack){
    return stack->head==nullptr;

}
void SPush(Stack& stack, string chislo){
    SNode* new_node = new SNode{chislo, stack.head};
    stack.head = new_node;
}

void SPop(Stack&  stack){
    if (SIsEmptyStack(&stack)){
        cout<<"no";
        return;
    }
    SNode* delete_node = stack.head;
    stack.head = stack.head->next;
    delete delete_node;
}
string STop(const Stack& stack){
    if (stack.head == nullptr){
        throw runtime_error("Стек пуст");
    }
    return stack.head->chislo;
}

void SClear(Stack& stack){
    while (!SIsEmptyStack(&stack)){
        SPop(stack);
    } 
}
void SPrint(SNode* flist){
    while(flist != nullptr){
        cout << flist->chislo<<"\n";
        flist = flist->next;
    }
}
void SPrintAddresses(const Stack& stack){
    SNode* cur_node = stack.head;
    while (cur_node != nullptr){
        cout << cur_node->chislo << " | адрес узла: " << cur_node << " | next: " << cur_node->next << '\n';
        cur_node = cur_node->next;
    }
    cout << '\n';
}


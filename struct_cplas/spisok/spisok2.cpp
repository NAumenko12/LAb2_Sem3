#include <iostream>
#include <string>
#include "spisok2.h"
using namespace std;

void LInitList(LList& list){
    list.head = nullptr;
    list.tail = nullptr;
}

void LPrint(LNode* flist){
    while(flist != nullptr){
        cout << flist->chislo<<"\n";
        flist = flist->next;
    }
}
void LPrintBac(LNode* last){
    while (last != nullptr){
        cout << last->chislo << '\n';
        last = last->prev;
    }
}

void LAddAfter(LList& list, LNode* flist, string chislo){
    LNode* new_node = new LNode{chislo, flist->next, flist};
    if (flist->next != nullptr){
        flist->next->prev = new_node;
    } else {
        list.tail = new_node;
    }
    flist->next = new_node;
}
void LAddFirst(LList& list, string chislo){
    LNode* new_node = new LNode{chislo, list.head, nullptr};
    if (list.head != nullptr){
        list.head->prev = new_node;
    } else {
        list.tail = new_node;
    }
    list.head = new_node;
}
void LAddEnd(LList& list, string chislo){
    LNode* new_node = new LNode{chislo, nullptr, list.tail};
    if (list.tail != nullptr){
        list.tail->next = new_node;
    } else {
        list.head = new_node;
    }
    list.tail = new_node;
}
void LAddBefore(LList& list, LNode* flist, string chislo){
    LNode* new_node = new LNode{chislo, flist, flist->prev};
    if (flist->prev != nullptr){
        flist->prev->next = new_node;
    } else {
        list.head = new_node;
    }
    flist->prev = new_node;
}
void LAddIndex(LList& list, string chislo, int index){
    if (index == 0){
        LAddFirst(list, chislo);
        return;
    }
    LNode* cur_node = list.head;
    for (int i = 0; i < index - 1; ++i){
        cur_node = cur_node->next;
    }
    LAddAfter(list, cur_node, chislo);
}

void LDeleteFirst(LList& list){
    LNode* delete_node = list.head;
    list.head = list.head->next;
    if (list.head != nullptr){
        list.head->prev = nullptr;
    } else {
        list.tail = nullptr;
    }
    delete delete_node;
}
void LDeleteEnd(LList& list){
    LNode* delete_node = list.tail;
    list.tail = list.tail->prev;
    if (list.tail != nullptr){
        list.tail->next = nullptr;
    } else {
        list.head = nullptr;
    }
    delete delete_node;
}
void LDeleteAfter(LList& list, LNode* flist){
    LNode* delete_node = flist->next;
    flist->next = delete_node->next;
    if (delete_node->next != nullptr){
        delete_node->next->prev = flist;
    } else {
        list.tail = flist;
    }
    delete delete_node;
}
void LDeleteBefore(LList& list, LNode* flist){
    LNode* delete_node = flist->prev;
    flist->prev = delete_node->prev;
    if (delete_node->prev != nullptr){
        delete_node->prev->next = flist;
    } else {
        list.head = flist;
    }
    delete delete_node;
}
void LDeleteIndex(LList& list, int index){
    LNode* cur_node = list.head;
    for (int i = 0; i < index; ++i){
        cur_node = cur_node->next;
    }
    LDeleteNode(list, cur_node);
}
void LDeleteNode(LList& list, LNode* flist){
    if (flist->prev != nullptr){
        flist->prev->next = flist->next;
    } else {
        list.head = flist->next;
    }
    if (flist->next != nullptr){
        flist->next->prev = flist->prev;
    } else {
        list.tail = flist->prev;
    }
    delete flist;
}
void LDeleteValue(LList& list, string chislo){
    LNode* cur_node = list.head;
    while (cur_node != nullptr && cur_node->chislo != chislo){
        cur_node = cur_node->next;
    }
    if (cur_node == nullptr){
        return;
    }
    LDeleteNode(list, cur_node);
}

void LFind(LNode* flist, string chislo){
    while (flist != nullptr){
        if (flist->chislo == chislo){
            cout << chislo << " найден, адрес: " << flist << "\n";
            return;
        }
        flist = flist->next;
    }
    cout << chislo << " не найден\n";
}
void LPrintAddresses(LNode* first){
    LNode* cur_node = first;
    while (cur_node != nullptr){
        cout << cur_node->chislo << " | адрес: " << cur_node << " | next: " << cur_node->next << " | prev: " << cur_node->prev << '\n';
        cur_node = cur_node->next;
    }
    cout << '\n';
}

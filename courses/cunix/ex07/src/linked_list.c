#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

node_t *list_create(void *data) {
    node_t *p = malloc(sizeof(node_t));
    p->data = data;
    p->next = NULL;
    return p;
}

void list_destroy(node_t **head, void (*fp)(void *data)) {
    if (*head == NULL) {
        return;
    }
    node_t *curr = *head;
    node_t *next = curr->next;
    while (curr != NULL) {
        fp(curr->data);
        free(curr);
        curr = next;
        if (curr != NULL) {
            next = curr->next;
        }
    }
}

// Add element by the end
void list_push(node_t *head, void *data) {
   node_t *p = head;
   while (p->next != NULL) {
        p = p->next;
   }
   p->next = list_create(data);
}

// Add element by the beginning
void list_unshift(node_t **head, void *data) {
    node_t *p = list_create(data);
    p->next = (*head);
    *head = p;
}

// Remove element by the end
void *list_pop(node_t **head) {
    node_t *to_pop = *head;
    node_t *prev = NULL;
    while (to_pop->next != NULL) {
        prev = to_pop;
        to_pop = to_pop->next;
    }
    if (prev != NULL) {
        prev->next = NULL;
    }
    void* data = to_pop->data;
    free(to_pop);
    return data;
}

// Remove element by the beginning
void *list_shift(node_t **head) {
    void *data = (*head)->data;
    node_t *to_pop = *head;
    *head = to_pop->next;
    free(to_pop);
    return data;
}

// Remove element by a pointer
void *list_remove(node_t **head, int pos) {
    node_t *prev = *head;
    while (pos > 0) {
        prev = prev->next;
        pos--;
    }
    node_t *to_del = prev->next;
    prev->next = prev->next->next;
    void *data = to_del->data;
    free(to_del);
    return data;
}

void list_print(node_t *head) {
    while (head != NULL) {
        printf("%s ", (char*)head->data);
        head = head->next;
    }
    printf("\n");
}

// Visits entire list and applies func fp
void list_visitor(node_t *head, void (*fp)(void *data)) {
    node_t *p = head;
    while (p != NULL) {
        fp(p->data);
        p = p->next;
    }
}

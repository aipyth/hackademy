#include <stdio.h>
#include "binary_tree.h"
#include <stdlib.h>
#include <string.h>

node_t *allocnode()
{
    node_t *n = (node_t *)malloc(sizeof(node_t));
    n->left = NULL;
    n->right = NULL;
    n->key = NULL;
    n->data = NULL;
    return n;
}

node_t *insert(node_t *root, char *key, void *data)
{
    node_t *new_node = allocnode();
    new_node->key = key;
    new_node->data = data;
    if (root != NULL)
    {
        node_t *curr = root;
        while (1)
        {
            int cmp = strcmp(key, curr->key);
            if (cmp == 0)
            {
                // this key exists
                free(new_node);
                return NULL;
            }
            else if (cmp == -1)
            {
                if (curr->left != NULL)
                {
                    curr = curr->left;
                }
                else
                {
                    curr->left = new_node;
                    break;
                }
            }
            else
            {
                if (curr->right != NULL)
                {
                    curr = curr->right;
                }
                else
                {
                    curr->right = new_node;
                    break;
                }
            }
            
        }
    }
    return new_node;
}

void print_node(node_t *node)
{
    if (node != NULL)
    {
        printf("%s: %s\n", node->key, (char *)node->data);
    }
}

void visit_tree(node_t *node, void (*fp)(node_t *root))
{
    if (node != NULL)
    {
        visit_tree(node->left, fp);
        fp(node);
        visit_tree(node->right, fp);
    }
}

void destroy_tree(node_t *node, void (*fdestroy)(node_t *root))
{
    if (node != NULL)
    {
        destroy_tree(node->left, fdestroy);
        destroy_tree(node->right, fdestroy);
        fdestroy(node);
        free(node);
    }
}

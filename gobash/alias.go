package gobash
/* alias.c -- Not a full alias, but just the kind that we use in the shell.  Csh style alias is somewhere else (`over there, in a 
   box'). */

/* Copyright (C) 1987-2009 Free Software Foundation, Inc.

   This file is part of GNU Bash, the Bourne Again SHell.

   Bash is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published 
   by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

   Bash is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along with Bash.  If not, see
   <http://www.gnu.org/licenses/>. */

type alias_t struct {
	name string
	value string
	flags int
}

/* Values for `flags' member of struct alias. */
const AL_EXPANDNEXT = 0x1
const AL_BEINGEXPANDED = 0x2

const ALIAS_HASH_BUCKETS = 16 /* must be power of two */

/* The list of known aliases. */

// TODO(krasin): add aliases state
//extern HASH_TABLE *aliases;

/* Non-zero means expand all words on the line.  Otherwise, expand after first expansion if the expansion ends in a space. */
//int alias_expand_all = 0;

/* The list of aliases that we have. */
//HASH_TABLE *aliases = (HASH_TABLE *) NULL;
//
//void initialize_aliases() {
//	if (aliases == 0)
//		aliases = hash_create(ALIAS_HASH_BUCKETS);
//}
//
///* Scan the list of aliases looking for one with NAME.  Return NULL if the alias doesn't exist, else a pointer to the alias_t. */
//alias_t *find_alias(name)
//	 char *name;
//{
//	BUCKET_CONTENTS *al;
//
//	if (aliases == 0)
//		return ((alias_t *) NULL);
//
//	al = hash_search(name, aliases, 0);
//	return (al ? (alias_t *) al->data : (alias_t *) NULL);
//}
//
///* Return the value of the alias for NAME, or NULL if there is none. */
//char *get_alias_value(name)
//	 char *name;
//{
//	alias_t *alias;
//
//	if (aliases == 0)
//		return ((char *)NULL);
//
//	alias = find_alias(name);
//	return (alias ? alias->value : (char *)NULL);
//}
//
///* Make a new alias from NAME and VALUE.  If NAME can be found, then replace its value. */
//void add_alias(name, value)
//	 char *name, *value;
//{
//	BUCKET_CONTENTS *elt;
//	alias_t *temp;
//	int n;
//
//	if (!aliases) {
//		initialize_aliases();
//		temp = (alias_t *) NULL;
//	} else
//		temp = find_alias(name);
//
//	if (temp) {
//		free(temp->value);
//		temp->value = savestring(value);
//		temp->flags &= ~AL_EXPANDNEXT;
//		n = value[strlen(value) - 1];
//		if (n == ' ' || n == '\t')
//			temp->flags |= AL_EXPANDNEXT;
//	} else {
//		temp = (alias_t *) xmalloc(sizeof(alias_t));
//		temp->name = savestring(name);
//		temp->value = savestring(value);
//		temp->flags = 0;
//
//		n = value[strlen(value) - 1];
//		if (n == ' ' || n == '\t')
//			temp->flags |= AL_EXPANDNEXT;
//
//		elt = hash_insert(savestring(name), aliases, HASH_NOSRCH);
//		elt->data = temp;
//	}
//}
//
///* Delete a single alias structure. */
//static void free_alias_data(data)
//	 PTR_T data;
//{
//	register alias_t *a;
//
//	a = (alias_t *) data;
//	free(a->value);
//	free(a->name);
//	free(data);
//}
//
///* Remove the alias with name NAME from the alias table.  Returns the number of aliases left in the table, or -1 if the alias
//   didn't exist. */
//int remove_alias(name)
//	 char *name;
//{
//	BUCKET_CONTENTS *elt;
//
//	if (aliases == 0)
//		return (-1);
//
//	elt = hash_remove(name, aliases, 0);
//	if (elt) {
//		free_alias_data(elt->data);
//		free(elt->key);			/* alias name */
//		free(elt);				/* XXX */
//		return (aliases->nentries);
//	}
//	return (-1);
//}
//
/* Delete all aliases. */
//void delete_all_aliases() {
//	if (aliases == 0)
//		return;
//
//	hash_flush(aliases, free_alias_data);
//	hash_dispose(aliases);
//	aliases = (HASH_TABLE *) NULL;
//}
//
///* Return an array of aliases that satisfy the conditions tested by FUNCTION. If FUNCTION is NULL, return all aliases. */
//static alias_t **map_over_aliases(function)
//	 sh_alias_map_func_t *function;
//{
//	register int i;
//	register BUCKET_CONTENTS *tlist;
//	alias_t *alias, **list;
//	int list_index;
//
//	i = HASH_ENTRIES(aliases);
//	if (i == 0)
//		return ((alias_t **) NULL);
//
//	list = (alias_t **) xmalloc((i + 1) * sizeof(alias_t *));
//	for (i = list_index = 0; i < aliases->nbuckets; i++) {
//		for (tlist = hash_items(i, aliases); tlist; tlist = tlist->next) {
//			alias = (alias_t *) tlist->data;
//
//			if (!function || (*function) (alias)) {
//				list[list_index++] = alias;
//				list[list_index] = (alias_t *) NULL;
//			}
//		}
//	}
//	return (list);
//}
//
//static void sort_aliases(array)
//	 alias_t **array;
//{
//	qsort(array, strvec_len((char **)array), sizeof(alias_t *), (QSFUNC *) qsort_alias_compare);
//}
//
//static int qsort_alias_compare(as1, as2)
//	 alias_t **as1, **as2;
//{
//	int result;
//
//	if ((result = (*as1)->name[0] - (*as2)->name[0]) == 0)
//		result = strcmp((*as1)->name, (*as2)->name);
//
//	return (result);
//}
//
///* Return a sorted list of all defined aliases */
//alias_t **all_aliases() {
//	alias_t **list;
//
//	if (aliases == 0 || HASH_ENTRIES(aliases) == 0)
//		return ((alias_t **) NULL);
//
//	list = map_over_aliases((sh_alias_map_func_t *) NULL);
//	if (list)
//		sort_aliases(list);
//	return (list);
//}
//
//char *alias_expand_word(s)
//	 char *s;
//{
//	alias_t *r;
//
//	r = find_alias(s);
//	return (r ? savestring(r->value) : (char *)NULL);
//}
//

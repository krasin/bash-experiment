/* syntax.h -- Syntax definitions for the shell */

/* Copyright (C) 2000, 2001, 2005, 2008,2009 Free Software Foundation, Inc.

   This file is part of GNU Bash, the Bourne Again SHell.

   Bash is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published 
   by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

   Bash is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along with Bash.  If not, see
   <http://www.gnu.org/licenses/>. */

#ifndef _SYNTAX_H_
#define _SYNTAX_H_

/* Defines for use by mksyntax.c */

#define slashify_in_quotes "\\`$\"\n"
#define slashify_in_here_document "\\`$"

#define shell_meta_chars   "()<>;&|"
#define shell_break_chars  "()<>;&| \t\n"

#define shell_quote_chars	"\"`'"

#if defined (PROCESS_SUBSTITUTION)
#  define shell_exp_chars		"$<>"
#else
#  define shell_exp_chars		"$"
#endif

#if defined (EXTENDED_GLOB)
#  define ext_glob_chars	"@*+?!"
#else
#  define ext_glob_chars	""
#endif
#define shell_glob_chars	"*?[]^"

/* Defines shared by mksyntax.c and the rest of the shell code. */

/* Values for character flags in syntax tables */

#define CWORD		0x0000		/* nothing special; an ordinary character */
#define CSHMETA		0x0001		/* shell meta character */
#define CSHBRK		0x0002		/* shell break character */
#define CBACKQ		0x0004		/* back quote */
#define CQUOTE		0x0008		/* shell quote character */
#define CSPECL		0x0010		/* special character that needs quoting */
#define CEXP		0x0020		/* shell expansion character */
#define CBSDQUOTE	0x0040		/* characters escaped by backslash in double quotes */
#define CBSHDOC		0x0080		/* characters escaped by backslash in here doc */
#define CGLOB		0x0100		/* globbing characters */
#define CXGLOB		0x0200		/* extended globbing characters */
#define CXQUOTE		0x0400		/* cquote + backslash */
#define CSPECVAR	0x0800		/* single-character shell variable name */
#define CSUBSTOP	0x1000		/* values of OP for ${word[:]OPstuff} */
#define CBLANK		0x2000		/* whitespace (blank) character */

/* Defines for use by the rest of the shell. */
extern int sh_syntaxtab[];
extern int sh_syntabsiz;

#define shellmeta(c)	(sh_syntaxtab[(unsigned char)(c)] & CSHMETA)
#define shellbreak(c)	(sh_syntaxtab[(unsigned char)(c)] & CSHBRK)
#define shellquote(c)	(sh_syntaxtab[(unsigned char)(c)] & CQUOTE)
#define shellxquote(c)	(sh_syntaxtab[(unsigned char)(c)] & CXQUOTE)

#define shellblank(c)	(sh_syntaxtab[(unsigned char)(c)] & CBLANK)

#define issyntype(c, t)	((sh_syntaxtab[(unsigned char)(c)] & (t)) != 0)
#define notsyntype(c,t) ((sh_syntaxtab[(unsigned char)(c)] & (t)) == 0)

#if defined (PROCESS_SUBSTITUTION)
#  define shellexp(c)	((c) == '$' || (c) == '<' || (c) == '>')
#else
#  define shellexp(c)	((c) == '$')
#endif

#if defined (EXTENDED_GLOB)
#  define PATTERN_CHAR(c) \
	((c) == '@' || (c) == '*' || (c) == '+' || (c) == '?' || (c) == '!')
#else
#  define PATTERN_CHAR(c) 0
#endif

#define GLOB_CHAR(c) \
	((c) == '*' || (c) == '?' || (c) == '[' || (c) == ']' || (c) == '^')

#define CTLESC '\001'
#define CTLNUL '\177'

#if !defined (HAVE_ISBLANK) && !defined (isblank)
#  define isblank(x)	((x) == ' ' || (x) == '\t')
#endif

int sh_syntabsiz = 256;
int sh_syntaxtab[256] = {
	CWORD,						/* 0 */
	CSPECL,						/* CTLESC */
	CWORD,						/* 2 */
	CWORD,						/* 3 */
	CWORD,						/* 4 */
	CWORD,						/* 5 */
	CWORD,						/* 6 */
	CWORD,						/* \a */
	CWORD,						/* \b */
	CSHBRK | CBLANK,			/* \t */
	CSHBRK | CBSDQUOTE,			/* \n */
	CWORD,						/* \v */
	CWORD,						/* \f */
	CWORD,						/* \r */
	CWORD,						/* 14 */
	CWORD,						/* 15 */
	CWORD,						/* 16 */
	CWORD,						/* 17 */
	CWORD,						/* 18 */
	CWORD,						/* 19 */
	CWORD,						/* 20 */
	CWORD,						/* 21 */
	CWORD,						/* 22 */
	CWORD,						/* 23 */
	CWORD,						/* 24 */
	CWORD,						/* 25 */
	CWORD,						/* 26 */
	CWORD,						/* ESC */
	CWORD,						/* 28 */
	CWORD,						/* 29 */
	CWORD,						/* 30 */
	CWORD,						/* 31 */
	CSHBRK | CBLANK,			/* SPC */
	CXGLOB | CSPECVAR,			/* ! */
	CQUOTE | CBSDQUOTE | CXQUOTE,	/* " */
	CSPECVAR,					/* # */
	CEXP | CBSDQUOTE | CBSHDOC | CSPECVAR,	/* $ */
	CWORD,						/* % */
	CSHMETA | CSHBRK,			/* & */
	CQUOTE | CXQUOTE,			/* ' */
	CSHMETA | CSHBRK,			/* ( */
	CSHMETA | CSHBRK,			/* ) */
	CGLOB | CXGLOB | CSPECVAR,	/* * */
	CXGLOB | CSUBSTOP,			/* + */
	CWORD,						/* , */
	CSPECVAR | CSUBSTOP,		/* - */
	CWORD,						/* . */
	CWORD,						/* / */
	CWORD,						/* 0 */
	CWORD,						/* 1 */
	CWORD,						/* 2 */
	CWORD,						/* 3 */
	CWORD,						/* 4 */
	CWORD,						/* 5 */
	CWORD,						/* 6 */
	CWORD,						/* 7 */
	CWORD,						/* 8 */
	CWORD,						/* 9 */
	CWORD,						/* : */
	CSHMETA | CSHBRK,			/* ; */
	CSHMETA | CSHBRK,			/* < */
	CSUBSTOP,					/* = */
	CSHMETA | CSHBRK,			/* > */
	CGLOB | CXGLOB | CSPECVAR | CSUBSTOP,	/* ? */
	CXGLOB | CSPECVAR,			/* @ */
	CWORD,						/* A */
	CWORD,						/* B */
	CWORD,						/* C */
	CWORD,						/* D */
	CWORD,						/* E */
	CWORD,						/* F */
	CWORD,						/* G */
	CWORD,						/* H */
	CWORD,						/* I */
	CWORD,						/* J */
	CWORD,						/* K */
	CWORD,						/* L */
	CWORD,						/* M */
	CWORD,						/* N */
	CWORD,						/* O */
	CWORD,						/* P */
	CWORD,						/* Q */
	CWORD,						/* R */
	CWORD,						/* S */
	CWORD,						/* T */
	CWORD,						/* U */
	CWORD,						/* V */
	CWORD,						/* W */
	CWORD,						/* X */
	CWORD,						/* Y */
	CWORD,						/* Z */
	CGLOB,						/* [ */
	CBSDQUOTE | CBSHDOC | CXQUOTE,	/* \ */
	CGLOB,						/* ] */
	CGLOB,						/* ^ */
	CWORD,						/* _ */
	CBACKQ | CQUOTE | CBSDQUOTE | CBSHDOC | CXQUOTE,	/* ` */
	CWORD,						/* a */
	CWORD,						/* b */
	CWORD,						/* c */
	CWORD,						/* d */
	CWORD,						/* e */
	CWORD,						/* f */
	CWORD,						/* g */
	CWORD,						/* h */
	CWORD,						/* i */
	CWORD,						/* j */
	CWORD,						/* k */
	CWORD,						/* l */
	CWORD,						/* m */
	CWORD,						/* n */
	CWORD,						/* o */
	CWORD,						/* p */
	CWORD,						/* q */
	CWORD,						/* r */
	CWORD,						/* s */
	CWORD,						/* t */
	CWORD,						/* u */
	CWORD,						/* v */
	CWORD,						/* w */
	CWORD,						/* x */
	CWORD,						/* y */
	CWORD,						/* z */
	CWORD,						/* { */
	CSHMETA | CSHBRK,			/* | */
	CWORD,						/* } */
	CWORD,						/* ~ */
	CSPECL,						/* CTLNUL */
	CWORD,						/* 128 */
	CWORD,						/* 129 */
	CWORD,						/* 130 */
	CWORD,						/* 131 */
	CWORD,						/* 132 */
	CWORD,						/* 133 */
	CWORD,						/* 134 */
	CWORD,						/* 135 */
	CWORD,						/* 136 */
	CWORD,						/* 137 */
	CWORD,						/* 138 */
	CWORD,						/* 139 */
	CWORD,						/* 140 */
	CWORD,						/* 141 */
	CWORD,						/* 142 */
	CWORD,						/* 143 */
	CWORD,						/* 144 */
	CWORD,						/* 145 */
	CWORD,						/* 146 */
	CWORD,						/* 147 */
	CWORD,						/* 148 */
	CWORD,						/* 149 */
	CWORD,						/* 150 */
	CWORD,						/* 151 */
	CWORD,						/* 152 */
	CWORD,						/* 153 */
	CWORD,						/* 154 */
	CWORD,						/* 155 */
	CWORD,						/* 156 */
	CWORD,						/* 157 */
	CWORD,						/* 158 */
	CWORD,						/* 159 */
	CWORD,						/* 160 */
	CWORD,						/* 161 */
	CWORD,						/* 162 */
	CWORD,						/* 163 */
	CWORD,						/* 164 */
	CWORD,						/* 165 */
	CWORD,						/* 166 */
	CWORD,						/* 167 */
	CWORD,						/* 168 */
	CWORD,						/* 169 */
	CWORD,						/* 170 */
	CWORD,						/* 171 */
	CWORD,						/* 172 */
	CWORD,						/* 173 */
	CWORD,						/* 174 */
	CWORD,						/* 175 */
	CWORD,						/* 176 */
	CWORD,						/* 177 */
	CWORD,						/* 178 */
	CWORD,						/* 179 */
	CWORD,						/* 180 */
	CWORD,						/* 181 */
	CWORD,						/* 182 */
	CWORD,						/* 183 */
	CWORD,						/* 184 */
	CWORD,						/* 185 */
	CWORD,						/* 186 */
	CWORD,						/* 187 */
	CWORD,						/* 188 */
	CWORD,						/* 189 */
	CWORD,						/* 190 */
	CWORD,						/* 191 */
	CWORD,						/* 192 */
	CWORD,						/* 193 */
	CWORD,						/* 194 */
	CWORD,						/* 195 */
	CWORD,						/* 196 */
	CWORD,						/* 197 */
	CWORD,						/* 198 */
	CWORD,						/* 199 */
	CWORD,						/* 200 */
	CWORD,						/* 201 */
	CWORD,						/* 202 */
	CWORD,						/* 203 */
	CWORD,						/* 204 */
	CWORD,						/* 205 */
	CWORD,						/* 206 */
	CWORD,						/* 207 */
	CWORD,						/* 208 */
	CWORD,						/* 209 */
	CWORD,						/* 210 */
	CWORD,						/* 211 */
	CWORD,						/* 212 */
	CWORD,						/* 213 */
	CWORD,						/* 214 */
	CWORD,						/* 215 */
	CWORD,						/* 216 */
	CWORD,						/* 217 */
	CWORD,						/* 218 */
	CWORD,						/* 219 */
	CWORD,						/* 220 */
	CWORD,						/* 221 */
	CWORD,						/* 222 */
	CWORD,						/* 223 */
	CWORD,						/* 224 */
	CWORD,						/* 225 */
	CWORD,						/* 226 */
	CWORD,						/* 227 */
	CWORD,						/* 228 */
	CWORD,						/* 229 */
	CWORD,						/* 230 */
	CWORD,						/* 231 */
	CWORD,						/* 232 */
	CWORD,						/* 233 */
	CWORD,						/* 234 */
	CWORD,						/* 235 */
	CWORD,						/* 236 */
	CWORD,						/* 237 */
	CWORD,						/* 238 */
	CWORD,						/* 239 */
	CWORD,						/* 240 */
	CWORD,						/* 241 */
	CWORD,						/* 242 */
	CWORD,						/* 243 */
	CWORD,						/* 244 */
	CWORD,						/* 245 */
	CWORD,						/* 246 */
	CWORD,						/* 247 */
	CWORD,						/* 248 */
	CWORD,						/* 249 */
	CWORD,						/* 250 */
	CWORD,						/* 251 */
	CWORD,						/* 252 */
	CWORD,						/* 253 */
	CWORD,						/* 254 */
	CWORD,						/* 255 */
};

#endif /* _SYNTAX_H_ */

/** \file
 * Copyright (c) 2008 Free Software Foundation, Inc. All rights reserved.<br>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License version 3 for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
 *
 * HISTORY
 * - 2008-03-08 DI Creation
 */

#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>

#include "utils.h"

/**
 *
 */

extern void DieError(
   int          errval,
   const char * fmt,
   ...)
{
    va_list ap;
    va_start(ap, fmt);

    vfprintf(stderr, fmt, ap);

    va_end(ap);

    exit(errval);
}

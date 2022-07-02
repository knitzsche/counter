/*
 * Copyright (C) 2022 Kyle Nitzsche 
 *
 */

package counter

import (
    "strings"
)


func Words(in string) (int, error) {
    words := strings.Fields(in)
    return len(words), nil

}


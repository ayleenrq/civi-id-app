package utils

import "strings"

func InputToML(g string) string {
    g = strings.ToLower(strings.TrimSpace(g))

    if g == "laki-laki" || g == "laki" || g == "pria" {
        return "MALE"
    }
    if g == "perempuan" || g == "wanita" {
        return "FEMALE"
    }
    return ""
}

func MLToIndo(g string) string {
    g = strings.ToUpper(strings.TrimSpace(g))

    if g == "MALE" {
        return "Laki-laki"
    }
    if g == "FEMALE" {
        return "Perempuan"
    }
    return ""
}

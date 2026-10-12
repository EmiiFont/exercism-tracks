fun reverse(input: String): String {
    var j = input.length - 1
    var reversed = ""

    while (j >= 0) {
        reversed += input[j]
        j--
    }

    return reversed
}

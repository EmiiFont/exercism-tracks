object Bob {
    fun hey(input: String): String {
        var cleaned = input.trim()
        if (cleaned.isBlank()) {
            return "Fine. Be that way!"
        }
        if (cleaned[cleaned.length - 1] == '?' && strIsUpperCase(cleaned)) {
            return "Calm down, I know what I'm doing!"
        }
        if (cleaned[cleaned.length - 1] == '?') {
            return "Sure."
        }
        if (strIsUpperCase(cleaned)) {
            return "Whoa, chill out!"
        }

        return "Whatever."
    }

    fun strIsUpperCase(input: String): Boolean {
        var hasNolowercase = input.none { it.isLowerCase() }
        var hasUppercase = input.any { it.isUpperCase() }
        return hasNolowercase && hasUppercase
    }
}

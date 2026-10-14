object CollatzCalculator {
    fun computeStepCount(start: Int): Int {
        if (start <= 0) {
            throw IllegalArgumentException("invalid value")
        }

        var res = start
        var steps = 0

        while (res > 1) {
            if (res % 2 == 0) {
                res /= 2
            } else if (res % 2 >= 1) {
                res = res * 3 + 1
            }
            steps++
        }

        return steps
    }
}

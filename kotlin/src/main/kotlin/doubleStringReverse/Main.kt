package info.ls.doubleStringReverse

fun main(args: Array<String>) {
    Util.doubleStringReverse("yo")
}

object Util{
    fun doubleStringReverse(input: String):String {
        return input.reversed().swapCase()
    }
}

fun String.swapCase() = map {
    when {
        it.isUpperCase() -> it.toLowerCase()
        it.isLowerCase() -> it.toUpperCase()
        else -> it
    }
}.joinToString("")
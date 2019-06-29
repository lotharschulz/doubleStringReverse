package info.ls.doubleStringReverse

fun main() {
    println(Util.doubleStringReverse("ESREVEr GNIRTs ELBUOd NILTOk"))
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
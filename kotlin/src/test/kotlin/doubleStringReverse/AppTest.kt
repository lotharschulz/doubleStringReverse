package info.ls.doubleStringReverse

import org.junit.Test
import org.junit.runner.RunWith
import org.junit.runners.Parameterized
import kotlin.test.assertEquals

@RunWith(Parameterized::class)
class AppTest (val input: String, val expectedOutput: String){
    companion object {
        @JvmStatic
        @Parameterized.Parameters(name = "{index}: doubleStringReverse({0})={1}")
        fun data() = listOf(
                arrayOf("be thEre oR Be sqUare", "ERAuQS Eb rO EReHT EB"),
                arrayOf("Hello World $", "$ DLROw OLLEh"),
                arrayOf("abcde", "EDCBA"),
                arrayOf("?=)(/&%$§!ß", "ß!§$%&/()=?"),
                arrayOf("123", "321")
        )
    }

    @Test
    fun test() {
        assertEquals(expectedOutput, Util.doubleStringReverse(input))
    }
}
import org.junit.Test

class GreetingTest {
    @Test fun greetingTest(){
        val t = "Hello World"
        val g = Greeting(t)
        val s = g.getGreeting()
        assert(s == t)
    }
}
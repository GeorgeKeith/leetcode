import org.junit.Test;

public class MainTest {
    @Test
    public void greetingTest() {
        String t = "Hello, world";
        Greeting g = new Greeting(t);
        assert(g.getT().equalsIgnoreCase(t));
    }
}

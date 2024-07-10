import org.junit.jupiter.api.Test;

import java.lang.reflect.Method;
import java.util.Arrays;

public class TestRunner {
    public static void main(String[] args) throws Exception {
        // Get the test class
        Class<?> testClass = BinarySearchTest.class;

        // Create an instance of the test class
        Object testInstance = testClass.getDeclaredConstructor().newInstance();

        // Get all methods in the test class
        Method[] methods = testClass.getDeclaredMethods();

        // Filter out only test methods
        Method[] testMethods = Arrays.stream(methods)
                .filter(method -> method.isAnnotationPresent(Test.class))
                .toArray(Method[]::new);

        // Run each test method
        for (Method method : testMethods) {
            try {
                method.invoke(testInstance);
                System.out.println(method.getName() + " PASSED");
            } catch (Exception e) {
                System.out.println(method.getName() + " FAILED");
                e.printStackTrace();
            }
        }
    }
}
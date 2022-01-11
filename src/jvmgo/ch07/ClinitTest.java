public class ClinitTest {
	public static int staticVar = 100;
	public int instanceVar = 200;

	public static void main(String[] args) {
		System.out.println(ClinitTest.staticVar);
		ClinitTest obj = new ClinitTest();
		System.out.println(obj.instanceVar);
	}
}

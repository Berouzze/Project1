Console.WriteLine("Введите числа");
Console.WriteLine("Число a: ");
int a = Convert.ToInt32(Console.ReadLine());
Console.WriteLine("Число b: ");
int b = Convert.ToInt32(Console.ReadLine());
if (a == b * b) {
    Console.WriteLine("true");
} else {
    Console.WriteLine("false");
}
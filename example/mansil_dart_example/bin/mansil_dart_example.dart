import 'package:mansil/mansil.dart';

void main(List<String> arguments) async {
  print("One");
  await Future.delayed(Duration(seconds: 2));
  print(Mansil.cursorUp1 + Mansil.clearLine  + "two");
}

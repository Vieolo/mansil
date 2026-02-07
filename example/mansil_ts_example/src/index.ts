import Mansil from "@vieolo/mansil";

async function main() {
    console.log("One");
    await new Promise((resolve) => setTimeout(resolve, 2000));
    console.log(Mansil.cursorUp1 + Mansil.clearLine + "two");
}

main();
